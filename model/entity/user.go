package entity

import "go-SkyTakeaway/model"

type User struct {
	ID       int    `json:"id" gorm:"primary_key"`
	OpenID   string `json:"openid" gorm:"column:openid"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Sex      string `json:"sex"`
	IdNumber string `json:"idNumber"`
	// 头像
	Avatar string `json:"avatar" `
	// gorm:"autoCreateTime" -> 自动填充当前时间为创建时间
	CreateTime model.LocalTime `json:"createTime" gorm:"autoCreateTime"`
}

// TableName 指定表名
func (User) TableName() string {
	return "user"
}
