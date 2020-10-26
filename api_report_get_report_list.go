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

//GetReportList 操作可返回与查询参数相匹配的、过去 90 天内所创建的报告列表。 每个请求最多可返回 100 个结果。
//
// 参考链接: http://docs.developer.amazonservices.com/en_US/reports/Reports_GetReportList.html
//
// 如果还可以返回其他结果，则响应中所返回的 HasNext 值为 true。
// 要检索所有结果，您可以将 NextToken 参数的值反复传递给 GetReportListByNextToken 操作，直至 HasNext 的返回值为 false。
//
//@acknowledged 用来指明在之前调用 UpdateReportAcknowledgements 时是否已确认订单报告。所发布订单报告如已经确认，则设置为 true；
// 所发布订单报告如未经确认，则设置为 false。(1:true, 0:false -1:未设置)此过滤器仅对订单报告有效；不支持商品报告。
//
// 操作的最大请求限额为 10 个，恢复速率为每分钟 1 个请求。
func (s *ReportClient) GetReportList(ctx context.Context, request GetReportListRequest) (*ReportListResult, error) {
	data := Param{}.SetAction("GetReportList").Load(request)

	var response struct {
		ResponseMetadata
		ReportList *ReportListResult `xml:"GetReportListResult"`
	}

	if err := s.getResult(ctx, data, &response); err != nil {
		return nil, err
	}
	return response.ReportList, nil
}

//GetReportListByNextToken 操作可通过之前调用提供给 GetReportListByNextToken 或 GetReportList 的 NextToken 值
//
// 参考链接: http://docs.developer.amazonservices.com/en_US/reports/Reports_GetReportListByNextToken.html
//
// 返回与查询参数相匹配的报告列表，其中前一调用中的 HasNext 值为 true。
// 操作的最大请求限额为 30 个，恢复速率为每 2 秒 1 个请求。
func (s *ReportClient) GetReportListByNextToken(ctx context.Context, nextToken string) (*ReportListResult, error) {
	data := Param{}.SetAction("GetReportListByNextToken")
	data.Set("NextToken", nextToken)
	var response struct {
		ResponseMetadata
		ReportList *ReportListResult `xml:"GetReportListByNextTokenResult"`
	}
	if err := s.getResult(ctx, data, &response); err != nil {
		return nil, err
	}
	return response.ReportList, nil
}

//ReportListResult 获取报告列表的结果模型
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
