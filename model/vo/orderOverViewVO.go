package vo

// OrderOverViewVO 订单概览数据
type OrderOverViewVO struct {
	WaitingOrders   int `json:"waitingOrders"`
	DeliveredOrders int `json:"deliveredOrders"`
	CompletedOrders int `json:"completedOrders"`
	CancelledOrders int `json:"cancelledOrders"`
	AllOrders       int `json:"allOrders"`
}
