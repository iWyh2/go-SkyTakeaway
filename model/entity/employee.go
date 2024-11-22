package entity

import (
	"go-SkyTakeaway/model"
)

// Employee 员工数据模型
type Employee struct {
	ID       int    `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Sex      string `json:"sex"`
	IdNumber string `json:"idNumber"`
	Status   int    `json:"status" `
	// gorm:"autoCreateTime" -> 自动填充当前时间为创建时间
	CreateTime model.LocalTime `json:"createTime" gorm:"autoCreateTime"`
	// gorm:"autoUpdateTime" -> 自动填充当前时间为修改时间
	UpdateTime model.LocalTime `json:"updateTime" gorm:"autoUpdateTime"`
	CreateUser int             `json:"createUser"`
	UpdateUser int             `json:"updateUser"`
}

// TableName 指定表名
func (Employee) TableName() string {
	return "employee"
}
