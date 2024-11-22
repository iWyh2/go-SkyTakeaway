package entity

import (
	"go-SkyTakeaway/model"
)

// Dish 菜品数据模型
type Dish struct {
	ID          int    `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Name        string `json:"name"`
	CategoryID  int    `json:"categoryId"`
	Price       string `json:"price"`
	Image       string `json:"image"`
	Description string `json:"description"`
	Status      int    `json:"status"`
	// gorm:"autoCreateTime" -> 自动填充当前时间为创建时间
	CreateTime model.LocalTime `json:"createTime" gorm:"autoCreateTime"`
	// gorm:"autoUpdateTime" -> 自动填充当前时间为修改时间
	UpdateTime model.LocalTime `json:"updateTime" gorm:"autoUpdateTime"`
	CreateUser int             `json:"createUser"`
	UpdateUser int             `json:"updateUser"`
}

func (Dish) TableName() string {
	return "dish"
}
