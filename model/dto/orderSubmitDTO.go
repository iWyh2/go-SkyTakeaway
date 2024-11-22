package dto

import "go-SkyTakeaway/model"

// OrderSubmitDTO 用户下单接口参数
type OrderSubmitDTO struct {
	AddressBookId         int             `json:"addressBookId"`
	Amount                float64         `json:"amount"`
	DeliveryStatus        int             `json:"deliveryStatus"`
	EstimatedDeliveryTime model.LocalTime `json:"estimatedDeliveryTime"`
	PackAmount            float64         `json:"packAmount"`
	PayMethod             int             `json:"payMethod"`
	Remark                string          `json:"remark"`
	TablewareNumber       int             `json:"tablewareNumber"`
	TablewareStatus       int             `json:"tablewareStatus"`
}
