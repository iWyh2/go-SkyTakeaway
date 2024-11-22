package entity

type DishFlavor struct {
	ID     int    `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	DishID int    `json:"dish_id"`
	Name   string `json:"name"`
	Value  string `json:"value"`
}

func (DishFlavor) TableName() string {
	return "dish_flavor"
}
