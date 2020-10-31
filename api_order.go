package mws

//Orders 创建订单服务
func Orders(credential Credential, options ...Option) *Client {
	options = append(options, ApiOption("/Orders/2013-09-01", "2013-09-01"), CredentialOption(credential))
	return createClient(options...)
}
