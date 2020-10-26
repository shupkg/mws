package mws

import (
	"context"
	"time"
)

type GetReportRequestListRequest struct {
	//非负整数，表明待返回报告请求的最大数量。如果指定一个大于 100 的数字，请求将被拒绝。 默认值：10
	MaxCount int
	//	ReportType 枚举值的结构化列表。 默认值：全部
	ReportTypeList []string `mws:"ReportTypeList.Type"`
	//ReportRequestId 值的结构化列表。如果您传入 ReportRequestId 的值，则会忽略其他查询条件。 默认值：全部
	ReportRequestIdList []string `mws:"ReportRequestIdList.Id"`
	//报告处理状态的结构化列表，依照其来过滤报告请求。默认值：全部
	//枚举值: _SUBMITTED_, _IN_PROGRESS_, _CANCELLED_, _DONE_, _DONE_NO_DATA_
	ReportProcessingStatusList []string `mws:"ReportProcessingStatusList.Status"`
	//用于选择待报告数据日期范围的起始日期，数据格式为 ISO8601。 默认值：90 天前
	RequestedFromDate time.Time
	//用于选择待报告数据日期范围的结束日期，数据格式为 ISO8601。 默认值：现在
	RequestedToDate time.Time
}

// GetReportRequestList 返回可用于获取报告的 ReportRequestId 的报告请求列表。
// GetReportRequestList 操作返回与查询参数相匹配的报告请求列表。您可为报告状态、日期范围和报告类型指定查询参数。列表中包含每个报告请求的 ReportRequestId。您可以通过在 GetReportList 操作中指定 ReportRequestId 值，来获取 ReportId 值。
// 对于首次请求，最多可返回 100 个报告请求。如果要返回更多报告请求，则将响应中所返回的 HasNext 值设置为 true。要检索所有结果，您可以将 NextToken 参数的值重复传递给 GetReportRequestListByNextToken 操作，直至 HasNext 的返回值为 false。
// GetReportRequestList 操作的最大请求限额为 10 个，恢复速率为每 45 秒 1 个请求。有关限制术语的定义，请参阅限制。
func (s *ReportClient) GetReportRequestList(ctx context.Context, request GetReportRequestListRequest) (*ReportRequestListResult, error) {
	data := Param{}.SetAction("GetReportRequestList").Load(request)
	var response struct {
		ResponseMetadata
		ReportRequestList *ReportRequestListResult `xml:"GetReportRequestListResult"`
	}
	if err := s.getResult(ctx, data, &response); err != nil {
		return nil, err
	}
	return response.ReportRequestList, nil
}

// GetReportRequestListByNextToken 可通过之前请求提供给 GetReportRequestListByNextToken 或 GetReportRequestList 的 NextToken 值，返回报告请求列表，其中前一请求中的 HasNext 值为 true。
// GetReportRequestListByNextToken 操作返回与查询参数相匹配的报告请求列表。该操作使用之前请求提供给 GetReportRequestListByNextToken 或 GetReportRequestList 的 NextToken 值,其中前一请求中的 HasNext 值为 true。
// GetReportRequestListByNextToken 操作的最大请求限额为 30 个，恢复速率为每 2 秒 1 个请求。有关限制术语的定义，请参阅限制。
func (s *ReportClient) GetReportRequestListByNextToken(ctx context.Context, nextToken string) (*ReportRequestListResult, error) {
	data := Param{}.SetAction("GetReportRequestListByNextToken").Set("NextToken", nextToken)
	var response struct {
		ResponseMetadata
		ReportRequestList *ReportRequestListResult `xml:"GetReportRequestLisByNextTokentResult"`
	}
	if err := s.getResult(ctx, data, &response); err != nil {
		return nil, err
	}
	return response.ReportRequestList, nil
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
