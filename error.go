package mws

import "fmt"

// ErrorResponse 错误响应
type ErrorResponse struct {
	Type      string `xml:"Error>Type"`
	Code      string `xml:"Error>Code"`
	Message   string `xml:"Error>Message"`
	Detail    string `xml:"Error>Detail"`
	RequestID string `xml:"RequestID"`
}

func (err ErrorResponse) Error() string {
	return fmt.Sprintf("Request ERROR: [%s.%s] %s -> %s (RequestID: %s)", err.Type, err.Code, err.Message, err.Detail, err.RequestID)
}

//IsErrorResponse 是否错误响应
func IsErrorResponse(err error) bool {
	if err != nil {
		_, ok := err.(*ErrorResponse)
		return ok
	}
	return false
}

//IsServiceUnavailable 亚马逊 MWS 服务是否不可用, http响应代码应为 503, 错误表明亚马逊 MWS 服务不可用。。应当使用“指数退避”方法重试请求。
func IsServiceUnavailable(err error) bool {
	if err != nil {
		if resp, ok := err.(*ErrorResponse); ok {
			return resp.Code == "ServiceUnavailable"
		}
	}
	return false
}

//IsRequestThrottled 判断请求是否被限制, http响应代码应为 503, 错误表明您的请求已被限制。请查看您所提交请求类型的相关限制。请设定重试逻辑，以便在适当时间后重新发送请求，以免触发限制。
func IsRequestThrottled(err error) bool {
	if err != nil {
		if resp, ok := err.(*ErrorResponse); ok {
			return resp.Code == "RequestThrottled"
		}
	}
	return false
}

//IsUserAgentHeaderMalformed 判断是否是错误的UserAgent标头, 错误表明请求中所含的 User-Agent 标头为无效格式。请使用亚马逊MWS 客户端库中的代码创建 User-Agent 标头，或参阅相关文档，以了解可接受的 User-Agent 标头格式。
func IsUserAgentHeaderMalformed(err error) bool {
	if err != nil {
		if resp, ok := err.(*ErrorResponse); ok {
			return resp.Code == "UserAgentHeaderMissing" ||
				resp.Code == "UserAgentHeaderMalformed" ||
				resp.Code == "UserAgentHeaderLanguageAttributeMissing" ||
				resp.Code == "UserAgentHeaderMaximumLengthExceeded"
		}
	}
	return false
}
