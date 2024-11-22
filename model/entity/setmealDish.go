package entity

// SetmealDish 套餐菜品关系数据模型
type SetmealDish struct {
	ID        int `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	SetmealID int `json:"setmealId"`
	DishID    int `json:"dishId"`
	// 菜品名称（冗余字段）
	Name string `json:"name"`
	// 菜品原价
	Price  string `json:"price"`
	Copies int    `json:"copies"`
}

func (SetmealDish) TableName() string {
	return "setmeal_dish"
}
