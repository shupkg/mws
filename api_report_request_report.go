package mws

import "context"

//RequestReport 操作用于创建报告请求。亚马逊 MWS 处理报告请求，并在报告完成时，将报告请求的状态设置为 _DONE_。报告保留 90 天。
//
// 参考链接: http://docs.developer.amazonservices.com/en_US/reports/Reports_RequestReport.html
//
// 当调用 RequestReport 操作时，如果向可选请求参数 MarketplaceIdList 提供商城编号列表，则可以指定报告要涵盖的商城。
// 如果不指定商城编号，将使用本地商城编号。
// 请注意，MarketplaceIdList 请求参数不在日本和中国使用。
//
// 操作的最大请求限额为 15 个，恢复速率为每分钟 1 个请求。
func (s *ReportService) RequestReport(ctx context.Context, c *Credential, reportType string, params ...Values) (string, *ReportRequestInfo, error) {
	data := ActionValues("RequestReport")
	data.Set("ReportType", reportType)
	data.SetAll(params...)
	// data.Set("ReportOptions", reportOptions)
	// data.SetTime("StartDate", startDate)
	// data.SetTime("EndDate", endDate)
	// data.Sets("MarketplaceIdList.Id", marketplaceIDList...)

	var response struct {
		BaseResponse
		ReportRequestInfo *ReportRequestInfo `xml:"RequestReportResult>ReportRequestInfo"`
	}
	if _, err := s.FetchStruct(ctx, c, data, &response); err != nil {
		return "", nil, err
	}
	return response.RequestID, response.ReportRequestInfo, nil
}
