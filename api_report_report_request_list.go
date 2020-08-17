package mws

import (
	"context"
	"time"
)

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

// GetReportRequestList 返回可用于获取报告的 ReportRequestId 的报告请求列表。
//
// GetReportRequestList 操作返回与查询参数相匹配的报告请求列表。您可为报告状态、日期范围和报告类型指定查询参数。列表中包含每个报告请求的 ReportRequestId。您可以通过在 GetReportList 操作中指定 ReportRequestId 值，来获取 ReportId 值。
//
// 对于首次请求，最多可返回 100 个报告请求。如果要返回更多报告请求，则将响应中所返回的 HasNext 值设置为 true。要检索所有结果，您可以将 NextToken 参数的值重复传递给 GetReportRequestListByNextToken 操作，直至 HasNext 的返回值为 false。
//
// GetReportRequestList 操作的最大请求限额为 10 个，恢复速率为每 45 秒 1 个请求。有关限制术语的定义，请参阅限制。
func (s *ReportService) GetReportRequestList(ctx context.Context, c *Credential, params ...Values) (string, *ReportRequestListResult, error) {
	data := ActionValues("GetReportRequestList")
	data.SetAll(params...)
	var response struct {
		BaseResponse
		ReportRequestList *ReportRequestListResult `xml:"GetReportRequestListResult"`
	}
	if _, err := s.FetchStruct(ctx, c, data, &response); err != nil {
		return "", nil, err
	}
	return response.RequestID, response.ReportRequestList, nil
}

// GetReportRequestListByNextToken 可通过之前请求提供给 GetReportRequestListByNextToken 或 GetReportRequestList 的 NextToken 值，返回报告请求列表，其中前一请求中的 HasNext 值为 true。
//
// GetReportRequestListByNextToken 操作返回与查询参数相匹配的报告请求列表。该操作使用之前请求提供给 GetReportRequestListByNextToken 或 GetReportRequestList 的 NextToken 值,其中前一请求中的 HasNext 值为 true。
//
// GetReportRequestListByNextToken 操作的最大请求限额为 30 个，恢复速率为每 2 秒 1 个请求。有关限制术语的定义，请参阅限制。
func (s *ReportService) GetReportRequestListByNextToken(ctx context.Context, c *Credential, nextToken string) (string, *ReportRequestListResult, error) {
	data := ActionValues("GetReportRequestListByNextToken")
	data.Set("NextToken", nextToken)
	var response struct {
		BaseResponse
		ReportRequestList *ReportRequestListResult `xml:"GetReportRequestLisByNextTokentResult"`
	}
	if _, err := s.FetchStruct(ctx, c, data, &response); err != nil {
		return "", nil, err
	}
	return response.RequestID, response.ReportRequestList, nil
}
