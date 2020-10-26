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

var CreateHTTPClient = func() *http.Client {
	return &http.Client{}
}

//Client Client
type Client struct {
	httpc      *http.Client
	userAgent  string
	api        string
	version    string
	credential Credential
	debug      bool

	id  string
	raw []byte
}

type ClientOption interface {
	Apply(*Client)
}

type ClientOptionFunc func(*Client)

func (f ClientOptionFunc) Apply(c *Client) {
	f(c)
}

type RequestID interface {
	GetRequestID() string
}

func createClient(options ...ClientOption) *Client {
	c := &Client{
		httpc:     CreateHTTPClient(),
		userAgent: fmt.Sprintf("go-mws-sdk/v%s (Language=%s; Platform=%s-%s; sdk=github.com/shupkg/mws)", Version, strings.Replace(runtime.Version(), "go", "go/", -1), runtime.GOOS, runtime.GOARCH),
	}
	c.SetOptions(options...)
	return c
}

func (c *Client) SetOptions(options ...ClientOption) {
	for _, option := range options {
		option.Apply(c)
	}
}

func (c *Client) GetRequestID() string {
	return c.id
}

func (c *Client) GetRaw() []byte {
	return c.raw
}

//getResult 请求
func (c *Client) getBytes(ctx context.Context, data Param) ([]byte, error) {
	c.id = ""
	if len(c.raw) > 0 {
		c.raw = c.raw[:0]
	}

	c.doSignature(c.credential, data)
	u := "https://" + Amazon.GetServiceHost(c.credential.MarketplaceID) + c.api

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", c.userAgent)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.httpc.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	v, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.raw, err
	}
	c.raw = v
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
		return c.raw, nil
	}

	//请求不成功
	ex := Error{}
	if err := xml.Unmarshal(v, &ex); err != nil {
		return v, err
	}
	return v, &ex
}

//getResult 请求
func (c *Client) getResult(ctx context.Context, data Param, result interface{}) error {
	v, err := c.getBytes(ctx, data)
	if err != nil {
		return err
	}
	c.raw = v
	if len(v) > 0 {
		if result != nil {
			err = xml.Unmarshal(v, result)
			if err == nil {
				return err
			}
			if meta, ok := result.(RequestID); ok {
				c.id = meta.GetRequestID()
			}
		}

		if c.id == "" {
			var meta ResponseMetadata
			_ = xml.Unmarshal(v, &meta)
			c.id = meta.GetRequestID()
		}
	}
	return nil
}

//对参数签名
func (c *Client) doSignature(credential Credential, data Param) {
	data.Set(keyVersion, c.version)
	data.Set(keyAWSAccessKeyID, credential.AWSAccessKeyID)
	data.Set(keyMWSAuthToken, credential.MWSAuthToken)
	data.Set(keySellerID, credential.SellerID)
	data.Set(keySignatureMethod, "HmacSHA256")
	data.Set(keySignatureVersion, "2")
	data.Set(keyTimestamp, time.Now().UTC().Format(time.RFC3339))
	data.Del(keySignature)

	s := "POST\n" + Amazon.GetServiceHost(credential.MarketplaceID) + "\n" + c.api + "\n" + data.Encode()

	mac := hmac.New(sha256.New, []byte(credential.SecretKey))
	mac.Write([]byte(s))
	signature := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	data.Set(keySignature, signature)
}

//GetServiceStatus 获取服务状态
func (c *Client) GetServiceStatus(ctx context.Context) (string, error) {
	result := struct {
		Status string `xml:"GetServiceStatusResult>Status"`
	}{}

	if err := c.getResult(ctx, Param{}.SetAction("GetServiceStatus"), &result); err != nil {
		return "", err
	}
	return result.Status, nil
}
