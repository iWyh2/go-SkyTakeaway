package constant

// 状态常量
const (
	// Disable 禁用
	Disable = 0
	// Enable 启用
	Enable = 1
)

// 订单状态
const (
	// PendingPayment 待付款
	PendingPayment = 1 + iota
	// ToBeConfirmed 待接单
	ToBeConfirmed
	// Confirmed 已接单
	Confirmed
	// DeliveryInProgress 派送中
	DeliveryInProgress
	// Completed 已完成
	Completed
	// Cancelled 已取消
	Cancelled
)

// 支付状态
const (
	// UnPaid 未支付
	UnPaid = iota
	// Paid 已支付
	Paid
	// Refund 已退款
	Refund
)
