package mws

import (
	"context"
	"time"
)

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
func (s *ReportService) GetReportList(ctx context.Context, c *Credential, params ...Values) (string, *ReportListResult, error) {
	data := ActionValues("GetReportList")
	data.SetAll(params...)
	// data.SetInt("MaxCount", maxCount)
	// data.SetTime("AvailableFromDate", availableFromDate)
	// data.SetTime("AvailableToDate", availableToDate)
	// data.SetBool("Acknowledged", acknowledged)
	// data.Sets("ReportTypeList.Type", reportTypeList...)
	// data.Sets("ReportRequestIdList.Id", reportRequestIDList...)

	var response struct {
		BaseResponse
		ReportList *ReportListResult `xml:"GetReportListResult"`
	}
	if _, err := s.FetchStruct(ctx, c, data, &response); err != nil {
		return "", nil, err
	}
	return response.RequestID, response.ReportList, nil
}

//GetReportListByNextToken 操作可通过之前调用提供给 GetReportListByNextToken 或 GetReportList 的 NextToken 值
//
// 参考链接: http://docs.developer.amazonservices.com/en_US/reports/Reports_GetReportListByNextToken.html
//
// 返回与查询参数相匹配的报告列表，其中前一调用中的 HasNext 值为 true。
// 操作的最大请求限额为 30 个，恢复速率为每 2 秒 1 个请求。
func (s *ReportService) GetReportListByNextToken(ctx context.Context, c *Credential, nextToken string) (string, *ReportListResult, error) {
	data := ActionValues("GetReportListByNextToken")
	data.Set("NextToken", nextToken)
	var response struct {
		BaseResponse
		ReportList *ReportListResult `xml:"GetReportListByNextTokenResult"`
	}
	if _, err := s.FetchStruct(ctx, c, data, &response); err != nil {
		return "", nil, err
	}
	return response.RequestID, response.ReportList, nil
}
