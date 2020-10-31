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

func (c Credential) MaskedToken() string {
	return c.Mask(c.MWSAuthToken, 13)
}

func (c Credential) MaskedSecretKey() string {
	return c.Mask(c.SecretKey, 11)
}

func (c Credential) Mask(s string, size int) string {
	prefix := (size - 3) / 2
	suffix := size - 3 - (prefix * 2)

	if l := len(s); l >= size {
		return s[:prefix] + "***" + s[l-suffix:]
	}
	return s
}

//FromEnviron 从环境变量获取授权令牌
func (c Credential) FromEnviron(sellerID, mwsAuthToken string) {
	c.SellerID = sellerID
	c.MWSAuthToken = mwsAuthToken
	c.AWSAccessKeyID = os.Getenv("AWSAccessKeyId")
	c.SecretKey = os.Getenv("SecretKey")
	c.MarketplaceID = os.Getenv("MarketplaceId")
}

//GetCredentialForTest 从环境变量获取授权令牌（用于测试）
func (c Credential) ForTest() {
	c.FromEnviron(os.Getenv("TestSellerId"), os.Getenv("TestMWSAuthToken"))
}
