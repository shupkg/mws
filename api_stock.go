package mws

import "context"

//FulfillmentInventoryService 库存管理
type FulfillmentInventoryService struct {
	*Client
}

//FulfillmentInventory FulfillmentInventory
func FulfillmentInventory() *FulfillmentInventoryService {
	return &FulfillmentInventoryService{
		Client: newClient("/FulfillmentInventory/2010-10-01", "2010-10-01"),
	}
}

//InventorySupplyResult InventorySupplyResult
type InventorySupplyResult struct {
	NextToken           string
	InventorySupplyList []*InventorySupplyList `xml:"InventorySupplyList>member"`
}

//InventorySupplyList InventorySupplyList
type InventorySupplyList struct {
	Condition             string
	TotalSupplyQuantity   int
	InStockSupplyQuantity int
	FNSKU                 string
	ASIN                  string
	SellerSKU             string
}

// ListInventorySupply 返回卖家库存状况信息。
//
// **描述**
//
// 该 ListInventorySupply 操作可以返回卖家位于 亚马逊物流 和在当前入库货件中的库存的供应情况相关信息。您可以查看您的亚马逊物流库存当前的供应状态，还可以找到库存供应状态发生变化的时间。
//
// 此操作不会返回库存供应情况的信息，此库存位于： 无法销售 绑定买家订单
//
// **限制**
//
// 该 ListInventorySupply 操作的最大请求限额为 30 个，恢复速率为每秒钟 2 个请求。 有关限制术语的定义以及限制的完整解释，请参阅亚马逊MWS开发者指南中的限制：针对提交请求频率的限制。
//
// **请求参数**
//
// `SellerSkus` 为您想知道库存供应情况的商品指定的卖家 SKU 列表。如果未指定 QueryStartDateTime 的值必填。同时指定 QueryStartDateTime 和 SellerSkus 的值时，将返回一个错误。
// 有效值为您已经发运至亚马逊配送中心的商品指定的卖家 SKU。最大值：50, 类型：xs:string
//
// `QueryStartDateTime` 此日期用于选择您在某个指定日期后（或当时）已更改库存供应情况的商品，日期格式为 ISO 8601。	如果未指定 SellerSkus 的值必填。同时指定 QueryStartDateTime 和 SellerSkus 的值时，将返回一个错误。类型：xs:dateTime
//
// `ResponseGroup` 指明您是否想执行 ListInventorySupply 操作以返回 SupplyDetail 元素。 ResponseGroup 值：`Basic` - 不包括响应中的 SupplyDetail 元素, `Detailed` - 在响应中包含 SupplyDetail 元素, 默认值：Basic, 类型：xs:string
func (s *FulfillmentInventoryService) ListInventorySupply(ctx context.Context, c *Credential, params ...Values) (string, *InventorySupplyResult, error) {
	data := ActionValues("ListInventorySupply")
	data.SetAll(params...)

	var response struct {
		BaseResponse
		InventorySupplyResult *InventorySupplyResult `xml:"ListInventorySupplyResult"`
	}
	if _, err := s.FetchStruct(ctx, c, data, &response); err != nil {
		return "", nil, err
	}
	return response.RequestID, response.InventorySupplyResult, nil
}

//ListInventorySupplyByNextToken 同 ListInventorySupply
func (s *FulfillmentInventoryService) ListInventorySupplyByNextToken(ctx context.Context, c *Credential, nextToken string) (string, *InventorySupplyResult, error) {
	data := ActionValues("ListInventorySupply")
	data.Set("NextToken", nextToken)
	var response struct {
		BaseResponse
		InventorySupplyResult *InventorySupplyResult `xml:"ListInventorySupplyByNextTokenResult"`
	}
	if _, err := s.FetchStruct(ctx, c, data, &response); err != nil {
		return "", nil, err
	}
	return response.RequestID, response.InventorySupplyResult, nil
}
