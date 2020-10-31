package mws

import (
	"errors"
	"strings"
)

//IsMwsEx 是否错误响应
func IsMwsEx(err error) bool {
	_, ok := AsMwsEx(err)
	return ok
}

func AsMwsEx(err error) (ex Error, ok bool) {
	if err != nil {
		ok = errors.As(err, &ex)
	}
	return
}

func FindMwsEx(err error, finder func(ex Error) bool) bool {
	if ex, ok := AsMwsEx(err); ok {
		return finder(ex)
	}
	return false
}

func IsSDKEx(err error) bool {
	return FindMwsEx(err, func(ex Error) bool { return ex.Type == "SDK" })
}

//IsServiceUnavailable 亚马逊 MWS 服务是否不可用, http响应代码应为 503, 错误表明亚马逊 MWS 服务不可用。。应当使用“指数退避”方法重试请求。
func IsServiceUnavailable(err error) bool {
	return FindMwsEx(err, func(ex Error) bool { return ex.Code == "ServiceUnavailable" })
}

func IsAccessDenied(err error) bool {
	return FindMwsEx(err, func(ex Error) bool {
		return strings.Contains(ex.Code, "Access") && strings.Contains(ex.Code, "Denied")
	})
}

//IsThrottled 判断请求是否被限制, http响应代码应为 503, 错误表明您的请求已被限制。请查看您所提交请求类型的相关限制。请设定重试逻辑，以便在适当时间后重新发送请求，以免触发限制。
func IsThrottled(err error) bool {
	return FindMwsEx(err, func(ex Error) bool { return ex.Code == "RequestThrottled" })
}

//IsUserAgentMalformed 判断是否是错误的UserAgent标头, 错误表明请求中所含的 User-Agent 标头为无效格式。请使用亚马逊MWS 客户端库中的代码创建 User-Agent 标头，或参阅相关文档，以了解可接受的 User-Agent 标头格式。
func IsUserAgentMalformed(err error) bool {
	return FindMwsEx(err, func(ex Error) bool {
		return ex.Code == "UserAgentHeaderMissing" ||
			ex.Code == "UserAgentHeaderMalformed" ||
			ex.Code == "UserAgentHeaderLanguageAttributeMissing" ||
			ex.Code == "UserAgentHeaderMaximumLengthExceeded"
	})
}
