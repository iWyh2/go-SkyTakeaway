package dto

// OrderPaymentDTO 订单支付传递数据模型
type OrderPaymentDTO struct {
	OrderNumber string `json:"orderNumber"`
	PayMethod   int    `json:"payMethod"`
}
