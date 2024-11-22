package vo

// DishOverViewVO 菜品数据总览
type DishOverViewVO struct {
	Sold         int `json:"sold"`
	Discontinued int `json:"discontinued"`
}
