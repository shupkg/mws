package mws

//Orders 创建订单服务
func Orders(credential Credential) *OrderClient {
	return &OrderClient{createClient(ApiOption("/Orders/2013-09-01", "2013-09-01"), CredentialOption(credential))}
}

type OrderClient struct {
	*Client
}
