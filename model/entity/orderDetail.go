package entity

// OrderDetail 订单明细数据模型
type OrderDetail struct {
	Id         int     `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Name       string  `json:"name"`
	OrderId    int     `json:"orderId"`
	DishId     int     `json:"dishId"`
	SetmealId  int     `json:"setmealId"`
	DishFlavor string  `json:"dishFlavor"`
	Number     int     `json:"number"`
	Amount     float64 `json:"amount"`
	Image      string  `json:"image"`
}

// TableName 指定表名
func (OrderDetail) TableName() string {
	return "order_detail"
}
