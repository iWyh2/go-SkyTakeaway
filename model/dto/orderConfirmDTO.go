package dto

// OrderConfirmDTO 接单接收数据模型
type OrderConfirmDTO struct {
	OrderId any `json:"id"`
	Status  int `json:"status"`
}
