package mws

import (
	"errors"
	"fmt"
)

// Error 错误响应
type Error struct {
	Type      string `xml:"Error>Type"`
	Code      string `xml:"Error>Code"`
	Message   string `xml:"Error>Message"`
	Detail    string `xml:"Error>Detail"`
	RequestID string `xml:"RequestID"`
}

func (err *Error) HasError() bool {
	return err.Code != ""
}

func (err *Error) Error() string {
	s := ErrorCodes[err.Code]
	if s != "" {
		return fmt.Sprintf("[%s.%s]%s - %s", err.Type, err.Code, s, err.Message)
	}
	return fmt.Sprintf("[%s.%s] %s -> %s (RequestID: %s)", err.Type, err.Code, err.Message, err.Detail, err.RequestID)
}

//IsErrorResponse 是否错误响应
func IsErrorResponse(err error) bool {
	if err != nil {
		var resp *Error
		return errors.As(err, &resp)
	}
	return false
}

//IsServiceUnavailable 亚马逊 MWS 服务是否不可用, http响应代码应为 503, 错误表明亚马逊 MWS 服务不可用。。应当使用“指数退避”方法重试请求。
func IsServiceUnavailable(err error) bool {
	if err != nil {
		var resp *Error
		if errors.As(err, &resp) {
			return resp.Code == "ServiceUnavailable"
		}
	}
	return false
}

//IsRequestThrottled 判断请求是否被限制, http响应代码应为 503, 错误表明您的请求已被限制。请查看您所提交请求类型的相关限制。请设定重试逻辑，以便在适当时间后重新发送请求，以免触发限制。
func IsRequestThrottled(err error) bool {
	if err != nil {
		var resp *Error
		if errors.As(err, &resp) {
			return resp.Code == "RequestThrottled"
		}
	}
	return false
}

//IsUserAgentHeaderMalformed 判断是否是错误的UserAgent标头, 错误表明请求中所含的 User-Agent 标头为无效格式。请使用亚马逊MWS 客户端库中的代码创建 User-Agent 标头，或参阅相关文档，以了解可接受的 User-Agent 标头格式。
func IsUserAgentHeaderMalformed(err error) bool {
	if err != nil {
		var resp *Error
		if errors.As(err, &resp) {
			return resp.Code == "UserAgentHeaderMissing" ||
				resp.Code == "UserAgentHeaderMalformed" ||
				resp.Code == "UserAgentHeaderLanguageAttributeMissing" ||
				resp.Code == "UserAgentHeaderMaximumLengthExceeded"
		}
	}
	return false
}

var ErrorCodes = map[string]string{
	"AccessDenied":                            "客户端尝试通过 HTTP 而不是 HTTPS 与亚马逊 MWS 连接。",
	"AccessToFeedProcessingResultDenied":      "没有足够权限访问上传数据处理结果。",
	"AccessToReportDenied":                    "没有足够权限访问所请求的报告。",
	"ContentMD5Missing":                       "缺少 Content-MD5 标头值。",
	"ContentMD5DoesNotMatch":                  "计算出的 MD5 哈希值与所提供的 Content-MD5 值不一致。",
	"FeedCanceled":                            "当请求已取消的上传数据的处理报告时返回。",
	"FeedProcessingResultNoLongerAvailable":   "无法下载上传数据处理结果。",
	"FeedProcessingResultNotReady":            "处理报告尚未生成。",
	"InputDataError":                          "上传数据内容包含错误。",
	"InternalError":                           "发生了未知的服务器错误。",
	"InvalidAccessKeyId":                      "提供的 AWSAccessKeyId 请求参数无效或过期。",
	"InvalidFeedSubmissionId":                 "提供的上传数据 Submission Id 无效。",
	"InvalidFeedType":                         "所提交的 Feed Type 无效。",
	"InvalidParameterValue":                   "提供的查询参数无效。例如，Timestamp 参数的格式不正确。",
	"InvalidQueryParameter":                   "提交了多余的参数。",
	"InvalidReportId":                         "提供的 Report Id 无效。",
	"InvalidReportType":                       "所提交的 Report Type 无效。",
	"InvalidRequest":                          "请求中由于缺少参数或参数无效，导致请求无法解析。",
	"InvalidScheduleFrequency":                "所提交的计划频率无效。",
	"MissingClientTokenId":                    "缺少 MerchantModel Id 参数或为空。",
	"MissingParameter":                        "查询中缺少必需的参数。",
	"ReportNoLongerAvailable":                 "无法下载指定的报告。",
	"ReportNotReady":                          "报告尚未生成。",
	"SignatureDoesNotMatch":                   "所提供的请求签名与服务器计算的签名值不一致。",
	"UserAgentHeaderLanguageAttributeMissing": "缺少 User-Agent 标头的 Language 属性。",
	"UserAgentHeaderMalformed":                "User-Agent 值不符合所需格式。",
	"UserAgentHeaderMaximumLengthExceeded":    "User-Agent 值超过 500 个字符。",
	"UserAgentHeaderMissing":                  "缺少 User-Agent 标头值。",
}
