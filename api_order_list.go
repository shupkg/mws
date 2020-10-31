package mws

import (
	"context"
	"time"
)

//ListOrdersRequest Client.ListOrders 请求的结构化参数表
type ListOrdersRequest struct {
	//指定某一格式为 ISO 8601 的日期，用以选择在该日期之后（或当天）创建的订单。
	//同时指定 CreatedAfter 和 LastUpdatedAfter 的值时，将返回一个错误。但必须指定其中一个
	//必须不迟于两分钟，且在请求提交时间之前。
	CreatedAfter time.Time

	//指定某一格式为 ISO 8601 的日期，用以选择在该日期之前（或当天）创建的订单。
	//只有在已指定 CreatedAfter 的值时，才可指定该参数的值。 如果已指定 LastUpdatedAfter 的值，则不可指定该参数的值。
	//必须迟于 CreatedAfter 的值。 必须不迟于两分钟，且在请求提交时间之前。
	//默认值：现在时间减去两分钟
	CreatedBefore time.Time

	//指定某一格式为 ISO 8601 的日期，用以选择最后更新日期为该日期之后（或当天）的订单。更新即为对订单状态进行更改，包括新订单的创建。包括亚马逊和卖家所进行的更新。
	//	同时指定 CreatedAfter 和 LastUpdatedAfter 的值时，将返回一个错误。
	//	如果指定了 LastUpdatedAfter 的值，则无法指定 BuyerEmail 和 SellerOrderId 的值。
	//	必须不迟于两分钟，且在请求提交时间之前。
	LastUpdatedAfter time.Time

	//指定某一格式为 ISO 8601 的日期，用以选择最后更新日期为该日期之前（或当天）的订单。更新即为对订单状态进行更改，包括新订单的创建。包括亚马逊和卖家所进行的更新。
	//	只有在已指定 LastUpdatedAfter 的值时，才可指定该参数的值。如果已指定 CreatedAfter 的值，则不可指定该参数的值。
	//	必须迟于 LastUpdatedAfter 的值。 必须不迟于两分钟，且在请求提交时间之前。
	//	默认值：现在时间减去两分钟
	LastUpdatedBefore time.Time

	//OrderStatus 值的列表。用来选择当前状态与您所指定的某个状态值相符的订单。
	//	PendingAvailability 只有预订订单才有此状态。订单已生成，但是付款未授权且商品的发售日期是将来的某一天。订单尚不能进行发货。请注意：仅在日本 (JP)，Preorder 才可能是一个 OrderType 值。
	//	Pending             订单已生成，但是付款未授权。订单尚不能进行发货。请注意：对于 OrderType = Standard 的订单，初始的订单状态是 Pending。对于 OrderType = Preorder 的订单（仅适用于 JP），初始的订单状态是 PendingAvailability，当进入付款授权流程时，订单状态将变为 Pending。
	//	Unshipped           付款已经过授权，订单已准备好进行发货，但订单中商品尚未发运。
	//	PartiallyShipped    订单中的一个或多个（但并非全部）商品已经发货。
	//	Shipped             订单中的所有商品均已发货。
	//	InvoiceUnconfirmed  订单中的所有商品均已发货。但是卖家还没有向亚马逊确认已经向买家寄出发票。请注意：此参数仅适用于中国地区。
	//	Canceled            订单已取消。
	//	Unfulfillable       订单无法进行配送。该状态仅适用于通过亚马逊零售网站之外的渠道下达但由亚马逊进行配送的订单。
	//	在此版本的 “订单 API”部分 中，必须同时使用 Unshipped 和 PartiallyShipped。仅使用其中一个状态值，则会返回错误。
	OrderStatus []string `mws:"OrderStatus.Status"`

	//MarketplaceId 值的列表。用来选择您所指定商城中的订单。卖家注册销售商品的商城。 如果该值不是卖家注册销售商品的商城，则会返回错误。最大值：50
	MarketplaceId []string `mws:"MarketplaceId.Id"`

	//指明订单配送方式的结构化列表。
	//	AFN: 亚马逊配送
	//	MFN: 卖家自行配送
	//	默认：全部
	FulfillmentChannel []string `mws:"FulfillmentChannel.Channel"`

	//PaymentMethod 值的列表。用来选择您指定的订单付款方式。
	//	COD: 货到现金付款
	//	CVS: 便利店付款
	//	Other
	//	COD 或 CVS 之外的任意付款方式
	//	注： COD 和 CVS 值只在日本有效。
	PaymentMethod []string

	//BuyerEmail 买家的电子邮件地址。用来选择包含指定电子邮件地址的订单。
	//	如果指定了 BuyerEmail 的值，则无法指定 FulfillmentChannel、 OrderStatus、 PaymentMethod、 LastUpdatedAfter、 LastUpdatedBefore 和 SellerOrderId 的值。
	//	您在请求中所提供的电子邮件地址可以匿名（亚马逊）也可以不匿名。
	BuyerEmail string

	//卖家所指定的订单编码。不是亚马逊订单编号。用来选择与卖家所指定订单编码相匹配的订单。
	//	如果指定了 SellerOrderId 的值，则无法指定FulfillmentChannel、 OrderStatus、 PaymentMethod、 LastUpdatedAfter、 LastUpdatedBefore 和 BuyerEmail 的值。
	SellerOrderId string

	//该数字用来指明每页可返回的最多订单数。该值必须介于 1 到 100 之间。默认值：100
	MaxResultsPerPage int64

	//TFMShipmentStatus 值的列表。用于选择使用亚马逊配送服务 (TFM) 且当前配送状态与您指定的某个状态值相符的订单。如果指定 TFMShipmentStatus，则仅返回 TFM 订单。
	//请注意：TFMShipmentStatus 请求参数仅适用于中国地区。
	//	PendingPickUp     亚马逊尚未从卖家处取件。
	//	LabelCanceled     卖家取消了取件。
	//	PickedUp          亚马逊已从卖家处取件。
	//	AtDestinationFC   包裹已经抵达亚马逊运营中心。
	//	Delivered         包裹已经配送给买家。
	//	RejectedByBuyer   包裹被买家拒收。
	//	Undeliverable     包裹无法配送。
	//	ReturnedToSeller  包裹未配送给买家，已经退还给卖家。
	//	Lost              包裹被承运人丢失。
	TFMShipmentStatus []string `mws:"TFMShipmentStatus.Status"`
}

//ListOrders 返回您在指定时间段内所创建或更新的订单。
//  **参考**
//    http://docs.developer.amazonservices.com/en_US/orders-2013-09-01/Orders_ListOrders.html
//    http://docs.developer.amazonservices.com/en_US/orders-2013-09-01/Orders_ListOrdersByNextToken.html
//  **描述**
//    该操作可返回您在指定时间段内创建或更新的订单列表。
//    您可以通过 `CreatedAfter` 参数或 `LastUpdatedAfter` 参数来指定时间段。
//    您必须使用其中一个参数，但不可同时使用两个参数。您还可以通过应用筛选条件来缩小返回的订单列表范围。
//    该 ListOrders 操作包括每个所返回订单的订单详情，
//    其中包括 `AmazonOrderId`、 `OrderStatus`、 `FulfillmentChannel` 和 `LastUpdateDate`。
//  **参数**
//    参考参数结构 ListOrdersRequest
//  **限额**
//    共享最大请求限额为 6 个，恢复速率为每分钟 1 个请求。
func (c *Client) ListOrders(ctx context.Context, request ListOrdersRequest, nextToken string) (result OrdersResult, err error) {
	if nextToken == "" {
		var resp struct {
			ResponseMetadata
			Orders OrdersResult `xml:"ListOrdersResult"`
		}
		err = c.getResult(ctx, "ListOrders", ParamStruct(request), &resp)
		result = resp.Orders
	} else {
		var resp struct {
			ResponseMetadata
			Orders OrdersResult `xml:"ListOrdersByNextTokenResult"`
		}
		err = c.getResult(ctx, "ListOrdersByNextToken", ParamNexToken(nextToken), &resp)
		result = resp.Orders
	}
	return
}

//OrdersResult Client.ListOrders 请求的订单结果
type OrdersResult struct {
	HasNext           bool
	NextToken         string
	LastUpdatedBefore time.Time
	CreatedBefore     time.Time
	Orders            []Order `xml:"Orders>Order"`
}

//Order Client.ListOrders 请求的订单结果
type Order struct {
	AmazonOrderID                string    `xml:"AmazonOrderId"`                //亚马逊所定义的订单编码，格式为 3-7-7。
	SellerOrderID                string    `xml:"SellerOrderId"`                //卖家所定义的订单编码。
	MarketplaceID                string    `xml:"MarketplaceId"`                //订单生成所在商城的匿名编码。
	PurchaseDate                 time.Time `xml:"PurchaseDate"`                 //创建订单的日期。
	LastUpdateDate               time.Time `xml:"LastUpdateDate"`               //订单的最后更新日期
	OrderStatus                  string    `xml:"OrderStatus"`                  //当前的订单状态。
	FulfillmentChannel           string    `xml:"FulfillmentChannel"`           //订单配送方式：亚马逊配送 (AFN) 或卖家自行配送 (MFN)。
	SalesChannel                 string    `xml:"SalesChannel"`                 //订单中第一件商品的销售渠道。
	OrderChannel                 string    `xml:"OrderChannel"`                 //订单中第一件商品的订单渠道。
	ShipServiceLevel             string    `xml:"ShipServiceLevel"`             //货件服务水平。
	ShippingAddress              Address   `xml:"ShippingAddress"`              //订单的配送地址。
	OrderTotal                   Money     `xml:"OrderTotal"`                   //订单的总费用。
	NumberOfItemsShipped         int       `xml:"NumberOfItemsShipped"`         //已配送的商品数量。
	NumberOfItemsUnshipped       int       `xml:"NumberOfItemsUnshipped"`       //未配送的商品数量。
	PaymentMethod                string    `xml:"PaymentMethod"`                //订单的主要付款方式。COD - 货到付款。仅适用于中国 (CN) 和日本 (JP)。CVS - 便利店。仅适用于日本 (JP)。Other - COD 和 CVS 之外的付款方式。注： 可使用多种次级付款方式为 PaymentMethod = COD的订单付款。每种次级付款方式均表示为 PaymentExecutionDetailItem 对象。
	BuyerEmail                   string    `xml:"BuyerEmail"`                   //买家的匿名电子邮件地址。
	BuyerName                    string    `xml:"BuyerName"`                    //买家姓名。
	BuyerCounty                  string    `xml:"BuyerCounty"`                  //this element is used only in the Brazil marketplace.
	ShipmentServiceLevelCategory string    `xml:"ShipmentServiceLevelCategory"` //订单的配送服务级别分类。 Expedited, FreeEconomy, NextDay, SameDay, SecondDay, Scheduled, Standard
	ShippedByAmazonTFM           bool      `xml:"ShippedByAmazonTFM"`           //指明订单配送方是否是亚马逊配送 (Amazon TFM) 服务。
	TFMShipmentStatus            string    `xml:"TFMShipmentStatus"`            //亚马逊 TFM订单的状态。仅当ShippedByAmazonTFM = True时返回。请注意：即使当 ShippedByAmazonTFM = True 时，如果您还没有创建货件，也不会返回 TFMShipmentStatus。值： PendingPickUp, LabelCanceled, PickedUp, AtDestinationFC, Delivered, RejectedByBuyer, Undeliverable, ReturnedToSeller 注： 亚马逊 TFM 仅适用于中国 (CN)。
	CbaDisplayableShippingLabel  string    `xml:"CbaDisplayableShippingLabel"`  //卖家自定义的配送方式，属于Checkout by Amazon (CBA) 所支持的四种标准配送设置中的一种。有关更多信息，请参阅您所在商城 Amazon Payments 帮助中心的“设置灵活配送方式”主题。 注： CBA 仅适用于美国 (US)、英国 (UK) 和德国 (DE) 的卖家。
	OrderType                    string    `xml:"OrderType"`                    //订单类型。StandardOrder - 包含当前有库存商品的订单。Preorder -所含预售商品（发布日期晚于当前日期）的订单。 注： Preorder 仅在日本 (JP) 是可行的OrderType 值。
	EarliestShipDate             time.Time `xml:"EarliestShipDate"`             //您承诺的订单发货时间范围的第一天。日期格式为 ISO 8601。 仅对卖家配送网络 (MFN) 订单返回。
	LatestShipDate               time.Time `xml:"LatestShipDate"`               //您承诺的订单发货时间范围的最后一天。日期格式为 ISO 8601。对卖家配送网络 (MFN)	和亚马逊物流 (AFN) 订单返回。
	EarliestDeliveryDate         time.Time `xml:"EarliestDeliveryDate"`         //您承诺的订单送达时间范围的第一天。日期格式为 ISO 8601。仅对没有 PendingAvailability、Pending 或 Canceled状态的 MFN 订单返回。
	LatestDeliveryDate           time.Time `xml:"LatestDeliveryDate"`           //您承诺的订单送达时间范围的最后一天。日期格式为 ISO 8601。仅对没有 PendingAvailability、Pending 或 Canceled状态的 MFN 订单返回。
	IsBusinessOrder              bool      `xml:"IsBusinessOrder"`              //true if the order is an Amazon Business order. An Amazon Business order is an order where the buyer is a Verified Business Buyer and the seller is an Amazon Business Seller. For more information about the Amazon Business Seller Program
	PurchaseOrderNumber          string    `xml:"PurchaseOrderNumber"`          //the purchase order (PO) number entered by the buyer at checkout.
	IsPrime                      bool      `xml:"IsPrime"`                      //true if the order is a seller-fulfilled Amazon Prime order.
	IsPremiumOrder               bool      `xml:"IsPremiumOrder"`               //true if the order has a Premium Shipping Service Level Agreement. For more information about Premium Shipping orders, see "Premium Shipping Options" in the Seller Central Help for your marketplace.
	ReplaceOrderID               string    `xml:"ReplaceOrderId"`               //the AmazonOrderId value for the order that is being replaced.
	IsReplacementOrder           bool      `xml:"IsReplacementOrder"`           //true if this is a replacement order.
	//PaymentExecutionDetail 货到付款 (COD) 订单的次级付款方式的相关信息。
	//  COD 订单是带有PaymentMethod = COD 的订单。
	//  包含一个或多个PaymentExecutionDetailItem响应元素。
	//  注： 对于使用某一次级付款方式付款的 COD 订单，将返回一个 PaymentExecutionDetailItem响应元素，
	//  其中 PaymentExecutionDetailItem/PaymentMethod = COD。
	//  对于使用多种次级付款方式付款的 COD 订单，将返回两个或多个 PaymentExecutionDetailItem响应元素。
	//仅对 COD 订单返回。仅适用于中国 (CN) 和日本 (JP)。
	PaymentExecutionDetail []*PaymentExecutionDetailItem `xml:"PaymentExecutionDetail"`
}

//Address  地址信息
type Address struct {
	Name          string `xml:"Name"`          //名称。
	AddressLine1  string `xml:"AddressLine1"`  //街道地址。
	AddressLine2  string `xml:"AddressLine2"`  //街道地址。
	AddressLine3  string `xml:"AddressLine3"`  //街道地址。
	City          string `xml:"City"`          //城市。
	County        string `xml:"County"`        //区县。
	District      string `xml:"District"`      //区。
	StateOrRegion string `xml:"StateOrRegion"` //省/自治区/直辖市或地区。
	PostalCode    string `xml:"PostalCode"`    //邮政编码。
	CountryCode   string `xml:"CountryCode"`   //两位数国家/地区代码。格式为 ISO 3166-1-alpha 2 。
	Phone         string `xml:"Phone"`         //电话号码。
	AddressType   string `xml:"AddressType"`   //地址类型
}

//PaymentExecutionDetailItem  付款信息
type PaymentExecutionDetailItem struct {
	Payment       Money  `xml:"Payment"`
	PaymentMethod string `xml:"PaymentMethod"`
}
