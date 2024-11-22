package vo

// OrderReportVO 订单统计返回数据模型
type OrderReportVO struct {
	DateList            string  `json:"dateList"`
	OrderCountList      string  `json:"orderCountList"`
	ValidOrderCountList string  `json:"validOrderCountList"`
	TotalOrderCount     int     `json:"totalOrderCount"`
	ValidOrderCount     int     `json:"validOrderCount"`
	OrderCompletionRate float64 `json:"orderCompletionRate"`
}
