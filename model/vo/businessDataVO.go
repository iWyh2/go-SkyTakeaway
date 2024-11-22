package vo

// BusinessDataVO 工作台今日数据概览
type BusinessDataVO struct {
	Turnover            float64 `json:"turnover"`
	ValidOrderCount     int     `json:"validOrderCount"`
	OrderCompletionRate float64 `json:"orderCompletionRate"`
	UnitPrice           float64 `json:"unitPrice"`
	NewUsers            int     `json:"newUsers"`
}
