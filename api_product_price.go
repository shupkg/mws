package mws

//GetMyPriceForASINResponse 商品价格获取响应体
type GetMyPriceForASINResponse struct {
	BaseResponse
	GetMyPriceForASINResult []GetMyPriceForASINResult
}

//GetMyPriceForASINResult GetMyPriceForASINResult
type GetMyPriceForASINResult struct {
	MarketplaceASIN MarketplaceASIN `xml:"Identifiers>MarketplaceASIN"` //商品ASIN标识
	OfferPrice      []ProductOffer  `xml:"Product>Offers>Offer"`        //供应价格（销售价格）
}

//GetMyPriceForSKUResponse 商品价格获取响应体
type GetMyPriceForSKUResponse struct {
	BaseResponse
	GetMyPriceForSKUResult []GetMyPriceForSKUResult
}

//GetMyPriceForSKUResult GetMyPriceForASINResult
type GetMyPriceForSKUResult struct {
	MarketplaceASIN MarketplaceASIN `xml:"Identifiers>MarketplaceASIN"` //商品ASIN标识
	SKUIdentifier   SKUIdentifier   `xml:"Identifiers>SKUIdentifier"`   //商品SKU标识
	OfferPrice      []ProductOffer  `xml:"Product>Offers>Offer"`        //供应价格（销售价格）
}

//ProductOffer ProductPrice
type ProductOffer struct {
	BuyingPrice        ProductBuyingPrice //包含价格信息（包括进行促销的商品）以及运费。
	RegularPrice       Money              //商品的当前价格（不包括进行促销的商品）。不包括运费。
	FulfillmentChannel string             //商品的配送渠道。Amazon - 亚马逊物流。Merchant - 卖家自行配送。
	ItemCondition      string             //商品的状况。有效值：New、Used、Collectible、Refurbished、Club
	ItemSubCondition   string             //商品的子状况(成色)。有效值：New、Mint、Very Good、Good、Acceptable、Poor、Club、OEM、Warranty、Refurbished Warranty、Refurbished、Open Box 或 Other。
	SellerSKU          string             //商品的 SellerSKU。
	SellerID           string             `xml:"SellerId"` //在操作中提交的 SellerId。
}

//ProductBuyingPrice BuyingPrice
type ProductBuyingPrice struct {
	ListingPrice Money //商品的当前价格（包括进行促销的商品）。
	LandedPrice  Money //ListingPrice + Shipping - Points.请注意，如果未返回到岸价格，则上市价格代表具有最低到岸价格的产品。
	Shipping     Money //商品的运费。
	Points       Money //购买商品时提供的亚马逊积分数量及其货币价值。请注意，Points元素仅在日本（JP）返回。
}

//SKUIdentifier SKUIdentifier
type SKUIdentifier struct {
	MarketplaceID string `xml:"MarketplaceId"`
	SellerID      string `xml:"SellerId"`
	SellerSKU     string
}

// GetMyPriceForSKU 根据 SellerSKU，返回您自己的商品的价格信息。
//
// GetMyPriceForSKU 操作会根据您指定的 SellerSKU 和 MarketplaceId，返回您自己的商品的价格信息。请注意，如果您提交了并未销售的商品的 SellerSKU，则此操作会返回空的 Offers 元素。此操作最多可返回 20 款商品的价格信息。
func (s *ProductService) GetMyPriceForSKU(c *Credential, marketplace string, sellerSKUList []string, itemCondition string) (requestID string, prices []GetMyPriceForSKUResult, err error) {
	data := ActionValues("GetMyPriceForSKU")
	data.Set("MarketplaceId", marketplace)
	data.Sets("SellerSKUList.SKU", sellerSKUList...)
	data.Set("ItemCondition", itemCondition)
	var result GetMyPriceForSKUResponse
	if err = s.GetModel(c, data, &result); err != nil {
		return
	}
	requestID = result.RequestID
	prices = result.GetMyPriceForSKUResult
	return
}

// GetMyPriceForASIN 根据 ASIN，返回您自己的商品的价格信息。
//
// GetMyPriceForASIN 操作与 GetMyPriceForSKU 大体相同，但前者使用 MarketplaceId 和 ASIN 来唯一标识一件商品，且不会返回 SKUIdentifier 元素。
func (s *ProductService) GetMyPriceForASIN(c *Credential, marketplace string, asinList []string, itemCondition string) (requestID string, prices []GetMyPriceForASINResult, err error) {
	data := ActionValues("GetMyPriceForASIN")
	data.Set("MarketplaceId", marketplace)
	data.Sets("ASINList.ASIN", asinList...)
	data.Set("ItemCondition", itemCondition)
	var result GetMyPriceForASINResponse
	if err = s.GetModel(c, data, &result); err != nil {
		return
	}
	requestID = result.RequestID
	prices = result.GetMyPriceForASINResult
	return
}
