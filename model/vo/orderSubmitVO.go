package vo

import "go-SkyTakeaway/model"

// OrderSubmitVO 用户下单接口返回结果
type OrderSubmitVO struct {
	OrderId     int             `json:"id"`
	OrderNumber string          `json:"orderNumber"`
	OrderAmount float64         `json:"orderAmount"`
	OrderTime   model.LocalTime `json:"orderTime"`
}
