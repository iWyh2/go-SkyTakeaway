package vo

import (
	"go-SkyTakeaway/model"
	"go-SkyTakeaway/model/entity"
)

// SetmealVO 套餐查询返回数据模型
type SetmealVO struct {
	ID           int             `json:"id"`
	Name         string          `json:"name"`
	CategoryID   int             `json:"categoryId"`
	Price        float64         `json:"price"`
	Image        string          `json:"image"`
	Description  string          `json:"description"`
	Status       int             `json:"status"`
	UpdateTime   model.LocalTime `json:"updateTime"`
	CategoryName string          `json:"categoryName"`
	// gorm:"-" -> 数据库读写时忽略该字段
	SetmealDishes []entity.SetmealDish `json:"setmealDishes" gorm:"-"`
}
