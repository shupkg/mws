package mws

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/tls"
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
	httpc     *http.Client
	userAgent string
	api       string
	version   string
}

func newClient(api, version string) *Client {
	return &Client{
		httpc:     &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}},
		userAgent: fmt.Sprintf("go-mws-sdk/v%s (Language=%s; Platform=%s-%s; sdk=github.com/shupkg/mws)", Version, strings.Replace(runtime.Version(), "go", "go/", -1), runtime.GOOS, runtime.GOARCH),
		api:       api,
		version:   version,
	}
}

//GetBytes 请求
func (c *Client) GetBytes(credential *Credential, data Values) ([]byte, error) {
	c.doSignature(credential, data)
	u := GetServiceBaseUrl(credential.MarketplaceID, c.api)

	req, err := http.NewRequest(http.MethodPost, u, strings.NewReader(data.Encode()))
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

//GetModel 请求并解析成指定模型
func (c *Client) GetModel(credential *Credential, data Values, result interface{}) error {
	v, err := c.GetBytes(credential, data)
	if err != nil {
		return err
	}
	return xml.Unmarshal(v, result)
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
func (c *Client) GetServiceStatus(credential *Credential) (string, error) {
	result := struct {
		Status string `xml:"GetServiceStatusResult>Status"`
	}{}
	err := c.GetModel(credential, ActionValues("GetServiceStatus"), &result)
	if err != nil {
		return "", err
	}
	return result.Status, nil
}
