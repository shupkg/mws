package mws

import (
	"context"
	"time"
)

type RequestReportRequest struct {
	//string	ReportType的值，指明待请求报告的类型。
	//	//类型：xs:string
	//	//
	//	//是	ReportType 值
	ReportType string
	//用于选择待报告数据日期范围的起始日期。
	//	//否	默认值：现在
	EndDate time.Time
	//用于选择待报告数据日期范围的结束日期。
	//	否	默认值：现在
	StartDate time.Time
	//	传递给报告的其他信息。
	//  如果报告接受 ReportOptions，ReportType 枚举 章节的报告描述中会介绍这些参数值。
	ReportOptions string
	//	一个包含您所注册的一个或多个商城编号的列表。生成的报告将包含您指定的所有商城的信息。
	//	示例： &MarketplaceIdList.Id.1=A13V1IB3VIYZZH &MarketplaceIdList.Id.2=A1PA6795UKMFR9
	//	请注意，MarketplaceIdList 请求参数不在日本和中国使用。
	//	您注册的商城编号。
	//	默认值：您注册的第一个商城。
	MarketplaceIdList []string `mws:"MarketplaceIdList.Id"`
}

//RequestReport 操作用于创建报告请求。亚马逊 MWS 处理报告请求，并在报告完成时，将报告请求的状态设置为 _DONE_。报告保留 90 天。
//
// 参考链接: http://docs.developer.amazonservices.com/en_US/reports/Reports_RequestReport.html
//
// 当调用 RequestReport 操作时，如果向可选请求参数 MarketplaceIdList 提供商城编号列表，则可以指定报告要涵盖的商城。
// 如果不指定商城编号，将使用本地商城编号。
// 请注意，MarketplaceIdList 请求参数不在日本和中国使用。
//
// 操作的最大请求限额为 15 个，恢复速率为每分钟 1 个请求。
func (s *ReportClient) RequestReport(ctx context.Context, request RequestReportRequest) (string, *ReportRequestInfo, error) {
	data := Param{}.SetAction("RequestReport").Load(request)

	var response struct {
		ResponseMetadata
		ReportRequestInfo *ReportRequestInfo `xml:"RequestReportResult>ReportRequestInfo"`
	}
	if err := s.getResult(ctx, data, &response); err != nil {
		return "", nil, err
	}
	return response.RequestID, response.ReportRequestInfo, nil
}
