package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/ulule/deepcopier"
	"go-SkyTakeaway/common/constant"
	errs "go-SkyTakeaway/common/errors"
	"go-SkyTakeaway/global"
	"go-SkyTakeaway/model/dto"
	"go-SkyTakeaway/model/entity"
	"gorm.io/gorm"
	"strconv"
)

// AddAddressBook 新增地址簿
func AddAddressBook(data *dto.AddressBookDTO, ctx *gin.Context) {
	// 取出登录用户id
	userId, exist := ctx.Get(constant.UserID)
	if !exist {
		panic(errs.MissUserIdError)
	}
	id, _ := strconv.Atoi(userId.(string))
	// 设置登录用户id
	data.UserId = id
	// 设置该地址不为默认地址
	data.IsDefault = 0
	// 准备数据容器
	var addressBook entity.AddressBook
	// 属性拷贝（使用deepcopier库）
	err := deepcopier.Copy(data).To(&addressBook)
	if err != nil {
		panic(errs.ServerInternalError)
	}
	// 处理label数据
	addressBook.Label = strconv.Itoa(data.Label)
	// 插入数据库
	if err := global.Db.Table("address_book").
		Create(&addressBook).Error; err != nil {
		panic(errs.DBError)
	}
}

// QueryAddressBooks 查询登录用户所有地址
func QueryAddressBooks(ctx *gin.Context) *[]entity.AddressBook {
	// 创建数据容器
	var addressBooks []entity.AddressBook
	// 获取已登录用户的userid
	userId, exist := ctx.Get(constant.UserID)
	if !exist {
		panic(errs.MissUserIdError)
	}
	// 查询数据库
	id, _ := strconv.Atoi(userId.(string))
	if err := global.Db.Table("address_book").
		Where("user_id = ?", id).
		Find(&addressBooks).Error; err != nil {
		panic(errs.DBError)
	}
	return &addressBooks
}

// GetDefaultAddressBook 查询默认地址
func GetDefaultAddressBook(ctx *gin.Context) (*entity.AddressBook, bool) {
	// 创建数据容器
	var addressBook entity.AddressBook
	// 获取已登录用户的userid
	userId, exist := ctx.Get(constant.UserID)
	if !exist {
		panic(errs.MissUserIdError)
	}
	id, _ := strconv.Atoi(userId.(string))
	// 查询数据库
	if err := global.Db.Table("address_book").
		Where("user_id = ?", id).
		Where("is_default = ?", 1).
		Find(&addressBook).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, false
		} else {
			panic(errs.DBError)
		}
	}
	return &addressBook, true
}

// SetDefaultAddressBook 设置默认地址
func SetDefaultAddressBook(data *entity.AddressBook, ctx *gin.Context) {
	// 获取已登录用户的userid
	userId, exist := ctx.Get(constant.UserID)
	if !exist {
		panic(errs.MissUserIdError)
	}
	// 将当前用户的所有地址修改为非默认地址
	if err := global.Db.Table("address_book").
		Where("user_id = ?", userId.(string)).
		Update("is_default", 0).Error; err != nil {
		panic(errs.DBError)
	}
	// 将当前地址改为默认地址
	if err := global.Db.Table("address_book").
		Where("id = ?", data.Id).
		Update("is_default", 1).Error; err != nil {
		panic(errs.DBError)
	}
}

// UpdateAddressBookById 根据id修改地址
func UpdateAddressBookById(data *dto.AddressBookDTO) {
	// 准备数据容器
	var addressBook entity.AddressBook
	// 属性拷贝（使用deepcopier库）
	err := deepcopier.Copy(data).To(&addressBook)
	if err != nil {
		panic(errs.ServerInternalError)
	}
	// 处理label数据
	addressBook.Label = strconv.Itoa(data.Label)
	// 修改数据库
	if err := global.Db.Table("address_book").
		Updates(&addressBook).Error; err != nil {
		panic(errs.DBError)
	}
}

// DeleteAddressBookById 根据id删除地址
func DeleteAddressBookById(id string) {
	// 处理数据
	Id, _ := strconv.Atoi(id)
	// 删除
	if err := global.Db.Delete(&entity.AddressBook{Id: Id}).Error; err != nil {
		panic(errs.DBError)
	}
}

// QueryAddressBookById 根据id查询地址
func QueryAddressBookById(id string) *entity.AddressBook {
	// 准备数据容器
	var addressBook entity.AddressBook
	// 处理数据
	Id, _ := strconv.Atoi(id)
	addressBook.Id = Id
	// 查询
	if err := global.Db.First(&addressBook).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		panic(errs.DBError)
	}
	return &addressBook
}
