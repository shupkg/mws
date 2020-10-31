package mws

import (
	"context"
	"time"
)

//ManageReportScheduleRequest Client.ManageReportSchedule 请求所使用的的参数结构
type ManageReportScheduleRequest struct {
	ReportType   string    //*ReportType 枚举，指明待请求报告的类型。
	Schedule     string    //*Schedule 枚举，表明报告请求的创建时间间隔。
	ScheduleDate time.Time //	下一个报告请求计划提交的日期。该值只能在 366 天之内。
}

//ManageReportSchedule 创建、更新或删除特定报告类型的报告请求计划。
//  **参考**
//    http://docs.developer.amazonservices.com/en_US/reports/Reports_ManageReportSchedule.html
//  **描述**
//    该操作创建、更新或删除特定报告类型的报告请求计划。目前，只能对订单和亚马逊商品广告报告进行计划。
//    通过组合使用 ReportType 和 Schedule 的值，亚马逊MWS 可确定您希望执行何种操作。如果不存在 ReportType 和 Schedule 的组合值，则将创建新的报告请求计划。
//    如果已计划 ReportType，但 Schedule 值不同，则将对报告请求计划进行更新，以便使用新的 Schedule 值。
//    如果您传入 ReportType，并在请求中将 Schedule 的值设置为 _NEVER_，则该 ReportType 的报告请求计划将被删除。
//  **参数**
//    ReportType    *ReportType 枚举，指明待请求报告的类型。
//    Schedule      *Schedule 枚举，表明报告请求的创建时间间隔。
//    ScheduleDate  下一个报告请求计划提交的日期。该值只能在 366 天之内。
//    下表显示了您可与 ManageReportSchedule 操作配合使用的 ReportType 枚举值。要获取每个报告的完整说明，请参阅 ReportType 枚举。
//      _GET_FLAT_FILE_ACTIONABLE_ORDER_DATA_                     待处理订单报告
//      _GET_ORDERS_DATA_                                         计划的 XML 订单报告
//      _GET_FLAT_FILE_ORDERS_DATA_                               请求或计划的文本文件订单报告
//      _GET_CONVERGED_FLAT_FILE_ORDER_REPORT_DATA_               文件文件订单报告
//      _GET_PADS_PRODUCT_PERFORMANCE_OVER_TIME_DAILY_DATA_TSV_   按 SKU 报告排列的商品广告每日绩效，文本文件
//      _GET_PADS_PRODUCT_PERFORMANCE_OVER_TIME_DAILY_DATA_XML_   按 SKU 报告排列的商品广告每日绩效，XML 文件
//      _GET_PADS_PRODUCT_PERFORMANCE_OVER_TIME_WEEKLY_DATA_TSV_  按 SKU 报告排列的商品广告每周绩效，文本文件
//      _GET_PADS_PRODUCT_PERFORMANCE_OVER_TIME_WEEKLY_DATA_XML_  按 SKU 报告排列的商品广告每周绩效，XML 文件
//      _GET_PADS_PRODUCT_PERFORMANCE_OVER_TIME_MONTHLY_DATA_TSV_ 按 SKU 报告排列的商品广告每月绩效，文本文件
//      _GET_PADS_PRODUCT_PERFORMANCE_OVER_TIME_MONTHLY_DATA_XML_ 按 SKU 报告排列的商品广告每月绩效，XML 文件
//  **限制**
//	  最大请求限额为 10 个，恢复速率为每 45 秒 1 个请求。
func (c *Client) ManageReportSchedule(ctx context.Context, param ManageReportScheduleRequest) (result []ReportSchedule, err error) {
	var resp struct {
		ResponseMetadata
		Result []ReportSchedule `xml:"ManageReportScheduleResult>ReportSchedule"`
	}
	err = c.getResult(ctx, "ManageReportSchedule", ParamStruct(param), &resp)
	result = resp.Result
	return
}

//GetReportScheduleList 返回计划提交至亚马逊MWS 进行处理的订单报告请求列表。
//  **参考**
//    http://docs.developer.amazonservices.com/en_US/reports/Reports_GetReportScheduleList.html
//  **描述**
//    该操作返回与查询参数相匹配的已计划订单报告请求列表。当前，只能对订单和亚马逊商品广告报告进行计划。每个请求最多可返回 100 个结果。
//  **限制**
//    最大请求限额为 10 个，恢复速率为每 45 秒 1 个请求。
func (c *Client) GetReportScheduleList(ctx context.Context, ReportTypeList []string) (result []ReportSchedule, err error) {
	var resp struct {
		ResponseMetadata
		Result []ReportSchedule `xml:"GetReportScheduleListResult>ReportSchedule"`
	}
	err = c.getResult(ctx, "GetReportScheduleList", ParamSet("ReportTypeList.Type", ReportTypeList), &resp)
	result = resp.Result
	return
}

//GetReportScheduleList 返回计划提交至亚马逊MWS 的订单报告请求计数。
//  **参考**
//    http://docs.developer.amazonservices.com/en_US/reports/Reports_GetReportScheduleCount.html
//  **描述**
//    该操作返回计划提交至亚马逊MWS 的报告请求计数。只能对订单和亚马逊商品广告报告进行计划。
//  **限制**
//    最大请求限额为 10 个，恢复速率为每 45 秒 1 个请求。
func (c *Client) GetReportScheduleCount(ctx context.Context, ReportTypeList []string) (result int64, err error) {
	var resp struct {
		ResponseMetadata
		Result int64 `xml:"GetReportScheduleCountResult>Count"`
	}
	err = c.getResult(ctx, "GetReportScheduleList", ParamSet("ReportTypeList.Type", ReportTypeList), &resp)
	result = resp.Result
	return
}

type ReportSchedule struct {
	ReportType    string    //所请求的 ReportType 值。
	Schedule      string    //Schedule 的值，指明报告请求的创建时间间隔。
	ScheduledDate time.Time //下一个报告请求计划提交的日期。
}
