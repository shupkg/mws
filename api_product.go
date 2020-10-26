package mws

//Products 创建商品服务
func Products(credential Credential) *OrderClient {
	return &OrderClient{createClient(ApiOption("/Products/2011-10-01", "2011-10-01"), CredentialOption(credential))}
}

//ProductService 商品服务
type ProductClient struct {
	*Client
}

//MarketplaceASIN 商城模型
type MarketplaceASIN struct {
	MarketplaceID string `xml:"MarketplaceId"`
	ASIN          string
}
