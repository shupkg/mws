package mws

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"strings"
	"time"
)

//Client Client
type Client struct {
	HttpClient *http.Client

	Credential Credential
	Api        string
	Version    string

	userAgent string
	debug     bool

	lastAction string
	lastRID    string
	lastRaw    []byte
}

func createClient(options ...Option) *Client {
	c := &Client{
		HttpClient: &http.Client{},
		userAgent:  fmt.Sprintf("go-mws-sdk/v%s (Language=%s; Platform=%s-%s; sdk=github.com/shupkg/mws)", Version, strings.Replace(runtime.Version(), "go", "go/", -1), runtime.GOOS, runtime.GOARCH),
	}
	c.SetOptions(defaultOptions...)
	c.SetOptions(options...)
	return c
}

func (c *Client) SetOptions(options ...Option) {
	for _, option := range options {
		option.Apply(c)
	}
}

func (c *Client) GetRequestID() string {
	return c.lastRID
}

func (c *Client) GetRaw() []byte {
	return c.lastRaw
}

func (c *Client) getBytes(ctx context.Context, action string, data Param, result interface{}) ([]byte, error) {
	// reset last data
	c.lastAction = action
	c.lastRID = ""
	if len(c.lastRaw) > 0 {
		c.lastRaw = c.lastRaw[:0]
	}
	if data == nil {
		data = Param{}
	}
	data.SetAction(action)

	c.doSignature(c.Credential, data)
	u := "https://" + Region.GetServiceHost(c.Credential.MarketplaceID) + c.Api

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, c.errorFromSDK(u, data.Encode(), "Param", err)
	}
	req.Header.Set("User-Agent", c.userAgent)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, c.errorFromRequest(u, data.Encode(), err)
	}
	defer resp.Body.Close()
	v, err := ioutil.ReadAll(resp.Body)
	c.lastRaw = v
	if err != nil {
		return c.lastRaw, c.errorFromSDK(u, data.Encode(), "Response", err)
	}

	if c.debug {
		log.Debugf("-------------------------")
		log.Debugf("请求地址: %s", u)
		log.Debugf("请求参数: %s", data.Encode())
		if len(v) > 500 {
			log.Debugf("响应内容: %d", len(v))
		} else {
			log.Debugf("响应内容: %s", string(v))
		}
		log.Debugf("-------------------------")
	}

	if resp.StatusCode == http.StatusOK {
		if result != nil {
			if err = xml.Unmarshal(v, result); err != nil {
				return v, c.errorFromXmlParser(u, data.Encode(), err)
			}

			if meta, ok := result.(RequestID); ok {
				c.lastRID = meta.GetRequestID()
			}
		}

		if c.lastRID == "" {
			var meta ResponseMetadata
			_ = xml.Unmarshal(v, &meta)
			c.lastRID = meta.GetRequestID()
		}

		return v, nil
	}

	//请求不成功
	var ex Error
	if err := xml.Unmarshal(v, &ex); err != nil {
		return v, c.errorFromXmlParser(u, data.Encode(), err)
	}
	return v, c.errorFill(u, data.Encode(), ex)
}

func (c *Client) getResult(ctx context.Context, action string, data Param, result interface{}) error {
	_, err := c.getBytes(ctx, action, data, result)
	return err
}

//对参数签名
func (c *Client) doSignature(credential Credential, data Param) {
	data.Set(keyVersion, c.Version)
	data.Set(keyAWSAccessKeyID, credential.AWSAccessKeyID)
	data.Set(keyMWSAuthToken, credential.MWSAuthToken)
	data.Set(keySellerID, credential.SellerID)
	data.Set(keySignatureMethod, "HmacSHA256")
	data.Set(keySignatureVersion, "2")
	data.Set(keyTimestamp, time.Now().UTC().Format(time.RFC3339))
	data.Del(keySignature)

	s := "POST\n" + Region.GetServiceHost(credential.MarketplaceID) + "\n" + c.Api + "\n" + data.Encode()

	mac := hmac.New(sha256.New, []byte(credential.SecretKey))
	mac.Write([]byte(s))
	signature := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	data.Set(keySignature, signature)
}
