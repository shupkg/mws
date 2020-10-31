package mws

import (
	"context"
	"time"
)

type GetReportListRequest struct {
	MaxCount            int64     //非负整数，表明待返回报告请求的最大数量。如果指定一个大于 100 的数字，请求将被拒绝。
	ReportTypeList      []string  //ReportType 枚举值的结构化列表
	Acknowledged        bool      `mws:",required"` //用来指明在之前调用 UpdateReportAcknowledgements 时是否已确认订单报告。所发布订单报告如已经确认，则设置为 true；所发布订单报告如未经确认，则设置为 false。此过滤器仅对订单报告有效；不支持商品报告。
	AvailableFromDate   time.Time //您可进行查找的最早日期，格式为 ISO8601。
	AvailableToDate     time.Time //您可进行查找的最近日期，格式为 ISO8601。
	ReportRequestIdList []string  //ReportRequestId 值的结构化列表。如果您传入 ReportRequestId 的值，则会忽略其他查询条件。
}

//GetReportList 返回在过去 90 天内所创建的报告列表。
//  **参考**
//    http://docs.developer.amazonservices.com/en_US/reports/Reports_GetReportList.html
//    http://docs.developer.amazonservices.com/en_US/reports/Reports_GetReportListByNextToken.html
//  **描述**
//    该操作可返回与查询参数相匹配的、过去 90 天内所创建的报告列表。每个请求最多可返回 100 个结果。
//   如果还可以返回其他结果，则响应中所返回的 HasNext 值为 true。
//   要检索所有结果，您可以将 NextToken 参数的值反复传递给 GetReportListByNextToken 操作，直至 HasNext 的返回值为 false。
//
//    对于首次请求，最多可返回 100 个报告请求。如果要返回更多报告请求，则将响应中所返回的 HasNext 值设置为 true。
//    要检索所有结果，您可以将 NextToken 参数的值重复传递给 GetReportRequestListByNextToken 操作，直至 HasNext 的返回值为 false。
//  **限制**
//    最大请求限额为 10 个，恢复速率为每分钟 1 个请求。
func (c *Client) GetReportList(ctx context.Context, param GetReportListRequest, nextToken string) (result ReportListResult, err error) {
	if nextToken == "" {
		var resp struct {
			ResponseMetadata
			Result ReportListResult `xml:"GetReportListResult"`
		}
		err = c.getResult(ctx, "GetReportList", ParamStruct(param), &resp)
		result = resp.Result
	} else {
		var resp struct {
			ResponseMetadata
			Result ReportListResult `xml:"GetReportListByNextTokenResult"`
		}
		err = c.getResult(ctx, "GetReportListByNextToken", ParamNexToken(nextToken), &resp)
		result = resp.Result
	}
	return
}

//ReportListResult ReportClient.GetReportList 获取报告列表的结果模型
type ReportListResult struct {
	HasNext   bool
	NextToken string
	Reports   []*ReportInfo `xml:"ReportInfo"`
}

//ReportInfo 报告信息在结果列表中的模型
type ReportInfo struct {
	ReportID        string `xml:"ReportId"`
	ReportRequestID string `xml:"ReportRequestId"`
	ReportType      string
	Acknowledged    bool
	AvailableDate   *time.Time
}
