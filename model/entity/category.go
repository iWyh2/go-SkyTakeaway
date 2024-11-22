package entity

import "go-SkyTakeaway/model"

// Category 分类数据模型
type Category struct {
	ID int `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	// 1菜品分类 2套餐分类
	Type string `json:"type"`
	Name string `json:"name"`
	// 用于分类数据的排序
	Sort   string `json:"sort"`
	Status int    `json:"status"`
	// gorm:"autoCreateTime" -> 自动填充当前时间为创建时间
	CreateTime model.LocalTime `json:"createTime" gorm:"autoCreateTime"`
	// gorm:"autoUpdateTime" -> 自动填充当前时间为修改时间
	UpdateTime model.LocalTime `json:"updateTime" gorm:"autoUpdateTime"`
	CreateUser int             `json:"createUser"`
	UpdateUser int             `json:"updateUser"`
}

// TableName 指定表名
func (Category) TableName() string {
	return "category"
}
