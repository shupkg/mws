package mws

import "os"

//Credential 令牌
type Credential struct {
	SellerID       string //商户: SellerId
	MWSAuthToken   string //商户: 令牌
	AWSAccessKeyID string //开发者凭证: AWSAccessKeyId
	SecretKey      string //开发者凭证: SecretKey
	MarketplaceID  string //站点: MarketplaceId
}

//GetCredentialFromEnv 从环境变量获取授权令牌
func GetCredentialFromEnv(sellerID, mwsAuthToken string) *Credential {
	return &Credential{
		SellerID:       sellerID,
		MWSAuthToken:   mwsAuthToken,
		AWSAccessKeyID: os.Getenv("AWSAccessKeyId"),
		SecretKey:      os.Getenv("SecretKey"),
		MarketplaceID:  os.Getenv("MarketplaceId"),
	}
}

//GetCredentialForTest 从环境变量获取授权令牌（用于测试）
func GetCredentialForTest() *Credential {
	return GetCredentialFromEnv(os.Getenv("TestSellerId"), os.Getenv("TestMWSAuthToken"))
}

func CredentialOption(credential Credential) ClientOption {
	return ClientOptionFunc(func(c *Client) {
		c.credential = credential
	})
}
