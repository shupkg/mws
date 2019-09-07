package mws

//ProductService 商品服务
type ProductService struct {
	*Client
}

//Products 创建商品服务
func Products() *ProductService {
	return &ProductService{
		Client: newClient("/Products/2011-10-01", "2011-10-01"),
	}
}

//MarketplaceASIN 商城模型
type MarketplaceASIN struct {
	MarketplaceID string `xml:"MarketplaceId"`
	ASIN          string `xml:"ASIN"`
}
