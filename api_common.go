package mws

import "context"

//GetServiceStatus 获取服务状态
func (c *Client) GetServiceStatus(ctx context.Context) (string, error) {
	result := struct {
		Status string `xml:"GetServiceStatusResult>Status"`
	}{}

	if err := c.getResult(ctx, "GetServiceStatus", nil, &result); err != nil {
		return "", err
	}
	return result.Status, nil
}
