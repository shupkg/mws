package mws

import (
	"context"
	"time"
)

func (c *Client) errorFrom(action string, params string, typ, code string, ex error) error {
	if ex == context.DeadlineExceeded || ex == context.Canceled {
		return ex
	}
	return Error{
		SellerID:      c.Credential.SellerID,
		AccessKeyID:   c.Credential.AWSAccessKeyID,
		Token:         c.Credential.MWSAuthToken,
		Action:        action,
		RequestParams: params,
		RequestID:     c.lastRID,
		Type:          typ,
		Code:          code,
		Message:       ex.Error(),
		Raw:           string(c.lastRaw),
		CreatedAt:     time.Now().Unix(),
	}
}

func (c *Client) errorFromSDK(action string, params string, code string, ex error) error {
	return c.errorFrom(action, params, "SDK", code, ex)
}

func (c *Client) errorFromXmlParser(action string, params string, ex error) error {
	return c.errorFromSDK(action, params, "XmlParseError", ex)
}

func (c *Client) errorFromRequest(action string, params string, ex error) error {
	if netEx, yes := isNetError(ex); yes {
		if netEx.Timeout() {
			return c.errorFromSDK(action, params, "Timeout", netEx)
		}
		return c.errorFromSDK(action, params, "Network", netEx)
	}
	return c.errorFromSDK(action, params, "Request", ex)
}

func (c *Client) errorFill(action string, params string, ex Error) Error {
	ex.SellerID = c.Credential.SellerID
	ex.AccessKeyID = c.Credential.AWSAccessKeyID
	ex.Token = c.Credential.MWSAuthToken
	ex.Action = action
	ex.RequestParams = params
	ex.Raw = string(c.lastRaw)
	ex.CreatedAt = time.Now().Unix()
	return ex
}
