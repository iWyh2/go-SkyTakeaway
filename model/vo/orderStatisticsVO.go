package vo

// OrderStatisticsVO 订单数量统计返回数据模型
type OrderStatisticsVO struct {
	ToBeConfirmed      int `json:"toBeConfirmed"`
	Confirmed          int `json:"confirmed"`
	DeliveryInProgress int `json:"deliveryInProgress"`
}
