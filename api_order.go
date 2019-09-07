package mws

//OrderService 订单服务
type OrderService struct {
	*Client
}

//Orders 创建订单服务
func Orders() *OrderService {
	return &OrderService{
		Client: newClient("/Orders/2013-09-01", "2013-09-01"),
	}
}
