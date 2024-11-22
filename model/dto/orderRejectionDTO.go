package dto

// OrderRejectionDTO 拒单接收数据模型
type OrderRejectionDTO struct {
	OrderId         int    `json:"id"`
	RejectionReason string `json:"rejectionReason"`
}
