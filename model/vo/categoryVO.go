package vo

import "go-SkyTakeaway/model"

// CategoryVO 返回分类数据模型
type CategoryVO struct {
	ID int `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	// 1菜品分类 2套餐分类
	Type int    `json:"type"`
	Name string `json:"name"`
	// 用于分类数据的排序
	Sort   string `json:"sort"`
	Status int    `json:"status"`
	// gorm:"autoCreateTime" -> 自动填充当前时间为创建时间
	CreateTime model.LocalTime `json:"createTime" gorm:"autoCreateTime"`
}
