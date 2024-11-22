package vo

import "go-SkyTakeaway/model/entity"

// OrderVO 查询订单详情返回数据模型
type OrderVO struct {
	entity.Order
	OrderDishes     string               `json:"orderDishes"`
	OrderDetailList []entity.OrderDetail `json:"orderDetailList"`
}
