package mws

//ReportService 报表服务
type ReportService struct {
	*Client
}

//Reports 创建报表服务
func Reports() *ReportService {
	return &ReportService{
		Client: newClient("/", "2009-01-01"),
	}
}

//GetReport 操作返回报告的内容及所返回报告正文的 Content-MD5 标头。报告从生成之日起保留 90 天。
//
// 参考链接: http://docs.developer.amazonservices.com/en_US/reports/Reports_GetReport.html
//
// 您应计算 HTTP 正文的 MD5 哈希值，并将该值与所返回的 Content-MD5 标头值进行比较。
// 如果二者不一致，则表明正文在传送过程中已损坏。如果报告已损坏，您应当放弃此结果，并自动重试请求，但最多不可超过三次。
// 请告知亚马逊 MWS 您所接收的报告正文是否已损坏。
// 亚马逊 MWS 网站上“报告 API”部分的客户端库中包含处理和比较 Content-MD5 标头的相关代码。
// 有关使用 Content-MD5 标头的更多信息，请参阅亚马逊 MWS 开发者指南。
//
// 操作的最大请求限额为 15 个，恢复速率为每分钟 1 个请求。
func (s *ReportService) GetReport(c *Credential, ReportID string) ([]byte, error) {
	data := ActionValues("GetReport")
	data.Set("ReportId", ReportID)
	return s.GetBytes(c, data)
}
