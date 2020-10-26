package mws

import "context"

type ProductRequest struct {
	//商城编码。指定返回商品信息的商城。
	MarketplaceId string

	//一个 SellerSKU 值的结构化列表。用于标识指定商城中的商品。SellerSKU 由您的 SellerId 限定，您提交的每个亚马逊商城网络服务（亚马逊 MWS）操作都需要包含您的 SellerId。
	// 最大值：20 个 SellerSKU
	SellerSKUList []string `mws:"SellerSKUList.SellerSKU"`

	//一个 ASIN 值的结构化列表。用于标识指定商城中的商品。 最大值：20 个 ASIN。
	ASINList []string `mws:"ASINList.ASIN"`

	//根据商品状况筛纳入考虑范围的商品。有效值：New、Used、Collectible、Refurbished、Club。
	ItemCondition string
}

// GetMyPriceForSKU 根据 SellerSKU，返回您自己的商品的价格信息。
// GetMyPriceForSKU 操作会根据您指定的 SellerSKU 和 MarketplaceId，返回您自己的商品的价格信息。请注意，如果您提交了并未销售的商品的 SellerSKU，则此操作会返回空的 Offers 元素。此操作最多可返回 20 款商品的价格信息。
func (s *ProductClient) GetMyPrice(ctx context.Context, request ProductRequest) (prices []GetMyPriceResult, err error) {
	data := Param{}.Load(request)

	if len(request.ASINList) > 0 {
		data.SetAction("GetMyPriceForASIN")
		var result struct {
			ResponseMetadata
			GetMyPriceForASINResult []GetMyPriceResult
		}
		if err = s.getResult(ctx, data, &result); err != nil {
			return
		}
		prices = result.GetMyPriceForASINResult
	} else if len(request.SellerSKUList) > 0 {
		data.SetAction("GetMyPriceForSKU")
		var result struct {
			ResponseMetadata
			GetMyPriceForSKUResult []GetMyPriceResult
		}
		if err = s.getResult(ctx, data, &result); err != nil {
			return
		}
		prices = result.GetMyPriceForSKUResult
	}
	return
}

//GetMyPriceResult GetMyPriceResult
type GetMyPriceResult struct {
	MarketplaceASIN MarketplaceASIN `xml:"Identifiers>MarketplaceASIN"` //商品ASIN标识
	SKUIdentifier   SKUIdentifier   `xml:"Identifiers>SKUIdentifier"`   //商品SKU标识, SKU请求时返回
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
