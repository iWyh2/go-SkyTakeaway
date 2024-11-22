package userController

import (
	"github.com/gin-gonic/gin"
	errs "go-SkyTakeaway/common/errors"
	"go-SkyTakeaway/model/dto"
	"go-SkyTakeaway/model/entity"
	model "go-SkyTakeaway/model/result"
	"go-SkyTakeaway/service"
	"log"
	"net/http"
)

// AddAddressBook 新增地址簿
func AddAddressBook(ctx *gin.Context) {
	// 创建接收数据模型
	var data dto.AddressBookDTO
	// 模型绑定
	if err := ctx.ShouldBindJSON(&data); err != nil {
		panic(errs.ServerInternalError)
	}
	// 日志打印
	log.Printf("新增地址簿: %v", data)
	// 调用service层进行处理
	service.AddAddressBook(&data, ctx)
	// 创建统一返回结果
	var result model.Result[string]
	// 响应
	ctx.JSON(http.StatusOK, result.Success())
}

// QueryAddressBooks 查询登录用户所有地址
func QueryAddressBooks(ctx *gin.Context) {
	// 调用service层进行处理
	addressBooks := service.QueryAddressBooks(ctx)
	// 日志打印
	log.Printf("查询登录用户所有地址: [%v]", (*addressBooks)[0].UserId)
	// 创建统一返回结果
	var result model.Result[[]entity.AddressBook]
	// 响应
	ctx.JSON(http.StatusOK, result.SuccessByData(*addressBooks))
}

// GetDefaultAddressBook 查询默认地址
func GetDefaultAddressBook(ctx *gin.Context) {
	// 调用service层进行处理
	addressBook, exist := service.GetDefaultAddressBook(ctx)
	// 日志打印
	log.Printf("查询默认地址: [%v]", addressBook.Id)
	// 创建统一返回结果
	var result model.Result[entity.AddressBook]
	// 响应
	if exist {
		ctx.JSON(http.StatusOK, result.SuccessByData(*addressBook))
	} else {
		ctx.JSON(http.StatusOK, result.Error("没有查询到默认地址"))
	}
}

// SetDefaultAddressBook 设置默认地址
func SetDefaultAddressBook(ctx *gin.Context) {
	// 创建接收数据模型
	var data entity.AddressBook
	// 模型绑定
	if err := ctx.ShouldBindJSON(&data); err != nil {
		panic(errs.ServerInternalError)
	}
	// 日志打印
	log.Printf("设置默认地址: [%v]", data.Id)
	// 调用service层进行处理
	service.SetDefaultAddressBook(&data, ctx)
	// 创建统一返回结果
	var result model.Result[string]
	// 响应
	ctx.JSON(http.StatusOK, result.Success())
}

// UpdateAddressBookById 根据id修改地址
func UpdateAddressBookById(ctx *gin.Context) {
	// 创建接收数据模型
	var data dto.AddressBookDTO
	// 模型绑定
	if err := ctx.ShouldBindJSON(&data); err != nil {
		panic(errs.ServerInternalError)
	}
	// 日志打印
	log.Printf("根据id修改地址: [%v]", data.Id)
	// 调用service层进行处理
	service.UpdateAddressBookById(&data)
	// 创建统一返回结果
	var result model.Result[string]
	// 响应
	ctx.JSON(http.StatusOK, result.Success())
}

// DeleteAddressBookById 根据id删除地址
func DeleteAddressBookById(ctx *gin.Context) {
	// 获取query参数
	id := ctx.Query("id")
	// 日志打印
	log.Printf("根据id删除地址: [%v]", id)
	// 调用service层进行处理
	service.DeleteAddressBookById(id)
	// 创建统一返回结果
	var result model.Result[string]
	// 响应
	ctx.JSON(http.StatusOK, result.Success())
}

// QueryAddressBookById 根据id查询地址
func QueryAddressBookById(ctx *gin.Context) {
	// 获取query参数
	id := ctx.Param("id")
	// 日志打印
	log.Printf("根据id查询地址: [%v]", id)
	// 调用service层进行处理
	addressBook := service.QueryAddressBookById(id)
	// 创建统一返回结果
	var result model.Result[entity.AddressBook]
	// 响应
	ctx.JSON(http.StatusOK, result.SuccessByData(*addressBook))
}
