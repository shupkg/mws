package mws

import (
	"context"
	"time"
)

type RequestReportRequest struct {
	ReportType        string    //*ReportType的值，指明待请求报告的类型。
	EndDate           time.Time //用于选择待报告数据日期范围的起始日期。默认值：现在
	StartDate         time.Time //用于选择待报告数据日期范围的结束日期。默认值：现在
	ReportOptions     string    //传递给报告的其他信息。如果报告接受 ReportOptions，ReportType 枚举 章节的报告描述中会介绍这些参数值。
	MarketplaceIdList []string  `mws:"MarketplaceIdList.Id"` //一个包含您所注册的一个或多个商城编号的列表。生成的报告将包含您指定的所有商城的信息。请注意，MarketplaceIdList 请求参数不在日本和中国使用。默认值:您注册的第一个商城。
}

type CancelReportRequestsRequest struct {
	ReportRequestIdList        []string  `mws:"ReportRequestIdList.Id"`            //ReportRequestId 值的结构化列表。如果您传入 ReportRequestId 的值，则会忽略其他查询条件。
	ReportTypeList             []string  `mws:"ReportTypeList.Type"`               //ReportType 枚举值的结构化列表。默认值：全部
	ReportProcessingStatusList []string  `mws:"ReportProcessingStatusList.Status"` //报告处理状态的结构化列表，依照其来过滤报告请求。取值: _SUBMITTED_,_IN_PROGRESS_,_CANCELLED_,_DONE_,_DONE_NO_DATA_, 默认值：全部
	RequestedFromDate          time.Time //用于选择待报告数据日期范围的起始日期，数据格式为 ISO8601。默认值：90 天前
	RequestedToDate            time.Time //用于选择待报告数据日期范围的结束日期，数据格式为 ISO8601。默认值：现在
}

//RequestReport 操作用于创建报告请求。亚马逊 MWS 处理报告请求，并在报告完成时，将报告请求的状态设置为 _DONE_。报告保留 90 天。
//  **参考**
//    http://docs.developer.amazonservices.com/en_US/reports/Reports_RequestReport.html
//  **描述**
//    当调用 RequestReport 操作时，如果向可选请求参数 MarketplaceIdList 提供商城编号列表，则可以指定报告要涵盖的商城。
//    如果不指定商城编号，将使用本地商城编号。
//    请注意，MarketplaceIdList 请求参数不在日本和中国使用。
//  **参数**
//    ReportType        *ReportType的值，指明待请求报告的类型。
//    EndDate           用于选择待报告数据日期范围的起始日期。默认值：现在
//    StartDate         用于选择待报告数据日期范围的结束日期。默认值：现在
//    ReportOptions     传递给报告的其他信息。如果报告接受 ReportOptions，ReportType 枚举 章节的报告描述中会介绍这些参数值。
//    MarketplaceIdList 一个包含您所注册的一个或多个商城编号的列表。生成的报告将包含您指定的所有商城的信息。请注意，MarketplaceIdList 请求参数不在日本和中国使用。默认值:您注册的第一个商城。
//  **限制**
//	  最大请求限额为 15 个，恢复速率为每分钟 1 个请求。
func (c *Client) RequestReport(ctx context.Context, param RequestReportRequest) (result ReportRequestInfo, err error) {
	var resp struct {
		ResponseMetadata
		Result ReportRequestInfo `xml:"RequestReportResult>ReportRequestInfo"`
	}
	err = c.getResult(ctx, "RequestReport", ParamStruct(param), &resp)
	result = resp.Result
	return
}

//CancelReportRequests 取消一个或多个报告请求。
//  **参考**
//    http://docs.developer.amazonservices.com/en_US/reports/Reports_CancelReportRequests.html
//  **描述**
//    该操作取消一个或多个报告请求，并返回被取消报告请求的数量及报告请求信息。您可以取消超过 100 个报告请求，但只能返回前 100 个被取消报告请求的相关信息。要返回更多被取消报告请求的相关信息，您可以使用 GetReportRequestList 操作。
//    如果报告请求已经开始处理，则无法取消它们。
//  **参数**
//    ReportRequestIdList         ReportRequestId 值的结构化列表。如果您传入 ReportRequestId 的值，则会忽略其他查询条件。 默认值：全部
//    ReportTypeList              ReportType 枚举值的结构化列表。 默认值：全部
//    ReportProcessingStatusList  报告处理状态的结构化列表，依照其来过滤报告请求。 取值:_SUBMITTED_,_IN_PROGRESS_,_CANCELLED_,_DONE_,_DONE_NO_DATA_, 默认值：全部
//    RequestedFromDate           用于选择待报告数据日期范围的起始日期，数据格式为 ISO8601。 默认值：90 天前
//    RequestedToDate             用于选择待报告数据日期范围的结束日期，数据格式为 ISO8601。 默认值：现在
//  **限制**
//    最大请求限额为 10 个，恢复速率为每 45 秒 1 个请求。
func (c *Client) CancelReportRequests(ctx context.Context, param CancelReportRequestsRequest) (result CancelReportRequestsResult, err error) {
	var resp struct {
		ResponseMetadata
		Result CancelReportRequestsResult `xml:"CancelReportRequestsResult"`
	}
	err = c.getResult(ctx, "CancelReportRequests", ParamStruct(param), &resp)
	result = resp.Result
	return
}

type CancelReportRequestsResult struct {
	Count             int
	ReportRequestInfo []ReportRequestInfo
}
