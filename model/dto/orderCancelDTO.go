package dto

// OrderCancelDTO 商家取消订单接收数据模型
type OrderCancelDTO struct {
	OrderId      int    `json:"id"`
	CancelReason string `json:"cancelReason"`
}
