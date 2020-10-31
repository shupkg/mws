package mws

//Products 创建商品服务
func Products(credential Credential, options ...Option) *Client {
	options = append(options, ApiOption("/Products/2011-10-01", "2011-10-01"), CredentialOption(credential))
	return createClient(options...)
}

//MarketplaceASIN 商城模型
type MarketplaceASIN struct {
	MarketplaceID string `xml:"MarketplaceId"`
	ASIN          string
}
