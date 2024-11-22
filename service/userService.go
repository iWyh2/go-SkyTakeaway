package service

import (
	"errors"
	errs "go-SkyTakeaway/common/errors"
	"go-SkyTakeaway/common/utils"
	"go-SkyTakeaway/global"
	"go-SkyTakeaway/model"
	"go-SkyTakeaway/model/dto"
	"go-SkyTakeaway/model/entity"
	"gorm.io/gorm"
	"strings"
	"time"
)

// WxLogin 用户微信登陆
func WxLogin(input *dto.UserLoginDTO) *entity.User {
	// 获取用户的openid
	openID := utils.GetOpenID(input.Code)
	// 判断openid是否为空，如果为空表示登录失败
	if openID == "" {
		panic(errs.WxLoginFailedError)
	}
	// 查询用户
	// 创建用户数据模型
	var user entity.User
	// 查询数据库
	if err := global.Db.Where("openid = ?", openID).First(&user).Error; err != nil {
		// 数据库未查询到，当前用户是新用户，自动完成注册
		if errors.Is(err, gorm.ErrRecordNotFound) {
			user.OpenID = openID
			// 创建数据库表数据
			if err := global.Db.Create(&user).Error; err != nil {
				if strings.Contains(err.Error(), "Duplicate entry") {
					// 用户已存在
					panic(errs.UserAlreadyExistsError)
				} else {
					// 数据库异常，创建失败
					panic(errs.DBError)
				}
			}
			// 返回这个用户对象
			return &user
		} else {
			// 数据库异常
			panic(errs.DBError)
		}
	}
	// 返回这个用户对象
	return &user
}

// GetUserCount 获取用户数量
func GetUserCount(begin, end time.Time) int {
	var userCount int64
	if begin.IsZero() {
		// 查询总用户数量
		if err := global.Db.Table("user").
			Where("create_time <= ?", model.LocalTime(end)).
			Count(&userCount).Error; err != nil {
			panic(errs.DBError)
		}
	} else {
		// 查询新增用户数量
		if err := global.Db.Table("user").
			Where("create_time >= ?", model.LocalTime(begin)).
			Where("create_time <= ?", model.LocalTime(end)).
			Count(&userCount).Error; err != nil {
			panic(errs.DBError)
		}
	}
	return int(userCount)
}
