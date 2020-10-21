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

var dbg = false

func SetDebug(debug bool) {
	dbg = debug
}

//Client Client
type Client struct {
	httpc     *http.Client
	userAgent string
	api       string
	version   string
}

func newClient(api, version string) *Client {
	return &Client{
		httpc:     CreateHTTPClient(),
		userAgent: fmt.Sprintf("go-mws-sdk/v%s (Language=%s; Platform=%s-%s; sdk=github.com/shupkg/mws)", Version, strings.Replace(runtime.Version(), "go", "go/", -1), runtime.GOOS, runtime.GOARCH),
		api:       api,
		version:   version,
	}
}

//FetchBytes 请求
func (c *Client) FetchBytes(ctx context.Context, credential *Credential, data Values) ([]byte, error) {
	c.doSignature(credential, data)
	u := GetServiceBaseUrl(credential.MarketplaceID, c.api)

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
		return v, err
	}

	if dbg {
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
		return v, nil
	}

	//请求不成功
	errResp := ErrorResponse{}
	err = xml.Unmarshal(v, &errResp)
	if err != nil {
		return v, err
	}
	return v, &errResp
}

//FetchStruct 请求并解析成指定模型
func (c *Client) FetchStruct(ctx context.Context, credential *Credential, data Values, result interface{}) ([]byte, error) {
	v, err := c.FetchBytes(ctx, credential, data)
	if err != nil {
		return v, err
	}
	err = xml.Unmarshal(v, result)
	return v, err
}

//对参数签名
func (c *Client) doSignature(credential *Credential, data Values) {
	data.Set(keyVersion, c.version)
	data.Set(keyAWSAccessKeyID, credential.AWSAccessKeyID)
	data.Set(keyMWSAuthToken, credential.MWSAuthToken)
	data.Set(keySellerID, credential.SellerID)
	data.Set(keySignatureMethod, "HmacSHA256")
	data.Set(keySignatureVersion, "2")
	data.Set(keyTimestamp, time.Now().UTC().Format(time.RFC3339))
	data.Del(keySignature)

	s := "POST\n" + GetServiceHost(credential.MarketplaceID) + "\n" + c.api + "\n" + data.Encode()

	mac := hmac.New(sha256.New, []byte(credential.SecretKey))
	mac.Write([]byte(s))
	signature := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	data.Set(keySignature, signature)
}

//GetServiceStatus 获取服务状态
func (c *Client) GetServiceStatus(ctx context.Context, credential *Credential) (string, error) {
	result := struct {
		Status string `xml:"GetServiceStatusResult>Status"`
	}{}
	_, err := c.FetchStruct(ctx, credential, ActionValues("GetServiceStatus"), &result)
	if err != nil {
		return "", err
	}
	return result.Status, nil
}
