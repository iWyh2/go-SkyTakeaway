package entity

import "go-SkyTakeaway/model"

// Order 订单数据模型
type Order struct {
	Id            int    `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Number        string `json:"number"`
	Status        int    `json:"status"`
	UserId        int    `json:"userId"`
	AddressBookId int    `json:"addressBookId"`
	// 下单时间，以当前时间自动插入
	OrderTime             model.LocalTime `json:"orderTime" gorm:"autoCreateTime"`
	CheckoutTime          model.LocalTime `json:"checkoutTime"`
	CancelTime            model.LocalTime `json:"cancelTime"`
	EstimatedDeliveryTime model.LocalTime `json:"estimatedDeliveryTime"`
	DeliveryTime          model.LocalTime `json:"deliveryTime"`
	PayMethod             int             `json:"payMethod"`
	PayStatus             int             `json:"payStatus"`
	Amount                float64         `json:"amount"`
	Remark                string          `json:"remark"`
	Username              string          `json:"username" gorm:"column:user_name"`
	Phone                 string          `json:"phone"`
	Address               string          `json:"address"`
	Consignee             string          `json:"consignee"`
	CancelReason          string          `json:"cancelReason"`
	RejectionReason       string          `json:"rejectionReason"`
	DeliveryStatus        int             `json:"deliveryStatus"`
	PackAmount            float64         `json:"packAmount"`
	TablewareNumber       int             `json:"tablewareNumber"`
	TablewareStatus       int             `json:"tablewareStatus"`
}

// TableName 指定表名
func (Order) TableName() string {
	return "orders"
}
