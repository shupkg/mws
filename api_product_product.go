package mws

import "context"

//GetMatchingProductForID 根据 ASIN、GCID、SellerSKU、UPC、EAN、ISBN 和 JAN，返回商品及其属性列表。
//根据您指定的商品编码值列表，GetMatchingProductForId 操作会返回一个包含商品及其属性的列表。可能的商品编号包括：ASIN、GCID、SellerSKU、UPC、EAN、ISBN 和 JAN。
func (s *ProductClient) GetMatchingProductForID(ctx context.Context, request ProductRequest) (string, []*Product, error) {
	data := Param{}.SetAction("GetMatchingProductForId").Load(request)

	var response struct {
		ResponseMetadata
		ProductResults []GetMatchingProductForIDResult `xml:"GetMatchingProductForIdResult"`
	}

	if err := s.getResult(ctx,  data, &response); err != nil {
		return "", nil, err
	}
	var products []*Product
	for _, result := range response.ProductResults {
		if result.Status == "Success" {
			products = append(products, result.Products...)
		}
	}
	return response.RequestID, products, nil
}

//GetMatchingProductForIDResult GetMatchingProductForIdResult
type GetMatchingProductForIDResult struct {
	ResponseMetadata
	ID       string     `xml:"Id,attr"`
	IDType   string     `xml:"IdType,attr"`
	Status   string     `xml:"status,attr"`
	Products []*Product `xml:"Products>Product"`
}

//Product 商品信息模型
type Product struct {
	MarketplaceASIN                      MarketplaceASIN  `xml:"Identifiers>MarketplaceASIN"`                                       //商品ASIN 商城
	Binding                              string           `xml:"AttributeSets>ItemAttributes>Binding"`                              //
	Brand                                string           `xml:"AttributeSets>ItemAttributes>Brand"`                                //商品品牌
	Color                                string           `xml:"AttributeSets>ItemAttributes>Color"`                                //型号，颜色
	Department                           string           `xml:"AttributeSets>ItemAttributes>Department"`                           //
	Feature                              string           `xml:"AttributeSets>ItemAttributes>Feature"`                              //功能
	ItemDimensions                       Dimension        `xml:"AttributeSets>ItemAttributes>ItemDimensions"`                       //商品尺寸
	IsAdultProduct                       bool             `xml:"AttributeSets>ItemAttributes>IsAdultProduct"`                       //
	IsAutographed                        bool             `xml:"AttributeSets>ItemAttributes>IsAutographed"`                        //
	IsEligibleForTradeIn                 bool             `xml:"AttributeSets>ItemAttributes>IsEligibleForTradeIn"`                 //
	IsMemorabilia                        bool             `xml:"AttributeSets>ItemAttributes>IsMemorabilia"`                        //
	IssuesPerYear                        bool             `xml:"AttributeSets>ItemAttributes>IssuesPerYear"`                        //
	Label                                string           `xml:"AttributeSets>ItemAttributes>Label"`                                //标签
	ListPrice                            Money            `xml:"AttributeSets>ItemAttributes>ListPrice"`                            //生产厂商
	Manufacturer                         string           `xml:"AttributeSets>ItemAttributes>Manufacturer"`                         //
	ManufacturerPartsWarrantyDescription string           `xml:"AttributeSets>ItemAttributes>ManufacturerPartsWarrantyDescription"` //厂商保修描述
	Warranty                             string           `xml:"AttributeSets>ItemAttributes>Warranty"`                             //保修描述
	MaterialType                         string           `xml:"AttributeSets>ItemAttributes>MaterialType"`                         //
	Model                                string           `xml:"AttributeSets>ItemAttributes>Model"`                                //型号
	PackageDimensions                    Dimension        `xml:"AttributeSets>ItemAttributes>PackageDimensions"`                    //包装信息
	PackageQuantity                      int              `xml:"AttributeSets>ItemAttributes>PackageQuantity"`                      //包装数量
	PartNumber                           string           `xml:"AttributeSets>ItemAttributes>PartNumber"`                           //
	ProductGroup                         string           `xml:"AttributeSets>ItemAttributes>ProductGroup"`                         //
	ProductTypeName                      string           `xml:"AttributeSets>ItemAttributes>ProductTypeName"`                      //
	Publisher                            string           `xml:"AttributeSets>ItemAttributes>Publisher"`                            //发布者
	SmallImage                           Image            `xml:"AttributeSets>ItemAttributes>SmallImage"`                           //商品图片（小图）
	Studio                               string           `xml:"AttributeSets>ItemAttributes>Studio"`                               //商品标题
	Title                                string           `xml:"AttributeSets>ItemAttributes>Title"`                                //商品标题
	VariationParent                      MarketplaceASIN  `xml:"Relationships>VariationParent>Identifiers>MarketplaceASIN"`         //父Asin信息
	VariationChild                       []VariationChild `xml:"Relationships>VariationChild"`                                      //子Asin信息
	SalesRankings                        []SalesRank      `xml:"SalesRankings>SalesRank"`                                           //排行信息
}

//SalesRank 分类排行
type SalesRank struct {
	ProductCategoryID string `xml:"ProductCategoryId"`
	Rank              int    `xml:"Rank"`
}

//VariationChild VariationChild
type VariationChild struct {
	Color           string          `xml:"Color"`
	MarketplaceASIN MarketplaceASIN `xml:"Identifiers>MarketplaceASIN"`
}

//Dimension 包装信息
type Dimension struct {
	Height UnitValue `xml:"Height"`
	Width  UnitValue `xml:"Width"`
	Length UnitValue `xml:"Length"`
	Weight UnitValue `xml:"Weight"`
}

//UnitValue 带单位的数量
type UnitValue struct {
	Unit  string  `xml:"Units,attr"`
	Value float64 `xml:",chardata"`
}

//Image 图片
type Image struct {
	URL    string    `xml:"URL"`
	Width  UnitValue `xml:"Width"`
	Height UnitValue `xml:"Height"`
}
