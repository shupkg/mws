package mws

import (
	"context"
	"time"
)

//GetReportRequestListRequest Client.GetReportRequestList 请求所使用的的参数结构
type GetReportRequestListRequest struct {
	MaxCount                   int       // 表明待返回报告请求的最大数量。如果指定一个大于 100 的数字，请求将被拒绝。 默认值：10
	ReportTypeList             []string  `mws:"ReportTypeList.Type"`               //ReportType 枚举值的结构化列表。 默认值：全部
	ReportRequestIdList        []string  `mws:"ReportRequestIdList.Id"`            //ReportRequestId 值的结构化列表。如果您传入 ReportRequestId 的值，则会忽略其他查询条件。 默认值：全部
	ReportProcessingStatusList []string  `mws:"ReportProcessingStatusList.Status"` //报告处理状态的结构化列表，依照其来过滤报告请求。枚举值: ReportProcessingStatus 默认值：全部
	RequestedFromDate          time.Time //用于选择待报告数据日期范围的起始日期，数据格式为 ISO8601。 默认值：90 天前
	RequestedToDate            time.Time //用于选择待报告数据日期范围的结束日期，数据格式为 ISO8601。 默认值：现在
}

//GetReportRequestCountRequest Client.GetReportRequestList 请求所使用的的参数结构
type GetReportRequestCountRequest struct {
	ReportTypeList             []string  `mws:"ReportTypeList.Type"`               //ReportType 枚举值的结构化列表。默认值：全部
	ReportProcessingStatusList []string  `mws:"ReportProcessingStatusList.Status"` //报告处理状态的结构化列表，依照其来过滤报告请求。取值: ReportProcessingStatus, 默认值：全部
	RequestedFromDate          time.Time //用于选择待报告数据日期范围的起始日期，数据格式为 ISO8601。默认值：90 天前
	RequestedToDate            time.Time //用于选择待报告数据日期范围的结束日期，数据格式为 ISO8601。默认值：现在
}

// GetReportRequestList 返回可用于获取报告的 ReportRequestId 的报告请求列表。
//  **参考**
//    http://docs.developer.amazonservices.com/en_US/reports/Reports_GetReportRequestList.html
//  **描述**
//    该方法是Api GetReportRequestList 和 GetReportRequestListByNextToken 合集，通过传递 nextToken 判断
//    该操作返回与查询参数相匹配的报告请求列表。您可为报告状态、日期范围和报告类型指定查询参数。
//    列表中包含每个报告请求的 ReportRequestId。
//    您可以通过在 GetReportList 操作中指定 GetReportRequestListRequest.ReportRequestIdList 值，来获取 ReportId 值。
//
//    对于首次请求，最多可返回 100 个报告请求。如果要返回更多报告请求，则将响应中所返回的 HasNext 值设置为 true。
//    要检索所有结果，您可以将 NextToken 参数的值重复传递给 GetReportList 操作，直至 HasNext 的返回值为 false。
//  **限制**
//    最大请求限额为 10 个，恢复速率为每 45 秒 1 个请求。
func (c *Client) GetReportRequestList(ctx context.Context, param GetReportRequestListRequest, nextToken string) (result ReportRequestListResult, err error) {
	if nextToken == "" {
		var resp struct {
			ResponseMetadata
			Result ReportRequestListResult `xml:"GetReportRequestListResult"`
		}
		err = c.getResult(ctx, "GetReportRequestList", ParamStruct(param), &resp)
		result = resp.Result
	} else {
		var resp struct {
			ResponseMetadata
			Result ReportRequestListResult `xml:"GetReportRequestLisByNextTokenResult"`
		}
		err = c.getResult(ctx, "GetReportRequestListByNextToken", ParamNexToken(nextToken), &resp)
		result = resp.Result
	}
	return
}

//GetReportRequestCount 返回已提交至亚马逊MWS 进行处理的报告请求计数。
//  **参考**
//    http://docs.developer.amazonservices.com/en_US/reports/Reports_GetReportRequestCount.html
//  **参数**
//    ReportTypeList              ReportType 枚举值的结构化列表。默认值：全部
//    ReportProcessingStatusList  报告处理状态的结构化列表，依照其来过滤报告请求。取值: _SUBMITTED_,_IN_PROGRESS_,_CANCELLED_,_DONE_,_DONE_NO_DATA_, 默认值：全部
//    RequestedFromDate           用于选择待报告数据日期范围的起始日期，数据格式为 ISO8601。默认值：90 天前
//    RequestedToDate             用于选择待报告数据日期范围的结束日期，数据格式为 ISO8601。默认值：现在
//  **限制**
//    最大请求限额为 10 个，恢复速率为每 45 秒 1 个请求。
func (c *Client) GetReportRequestCount(ctx context.Context, param GetReportRequestCountRequest) (result int, err error) {
	var resp struct {
		ResponseMetadata
		Result int `xml:"GetReportRequestCountResult>Count"`
	}
	err = c.getResult(ctx, "GetReportRequestCount", ParamStruct(param), &resp)
	result = resp.Result
	return
}

//ReportRequestListResult 获取报告请求列表结果
type ReportRequestListResult struct {
	NextToken         string
	HasNext           bool
	ReportRequestInfo []*ReportRequestInfo
}

//ReportRequestInfo 获取报告请求列表结果
type ReportRequestInfo struct {
	ReportRequestID        string `xml:"ReportRequestId"`
	ReportType             string
	StartDate              time.Time
	EndDate                time.Time
	Scheduled              bool
	SubmittedDate          time.Time
	ReportProcessingStatus string
	GeneratedReportID      string `xml:"GeneratedReportId"`
	StartedProcessingDate  time.Time
	CompletedDate          time.Time
}
