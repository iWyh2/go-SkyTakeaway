package dto

// ShoppingCartDTO 添加购物车传输数据模型
type ShoppingCartDTO struct {
	DishId     int    `json:"dishId"`
	SetmealId  int    `json:"setmealId"`
	DishFlavor string `json:"dishFlavor"`
}
