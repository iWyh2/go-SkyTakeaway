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

// AddShoppingCart 添加购物车
func AddShoppingCart(ctx *gin.Context) {
	// 创建接收数据模型
	var data dto.ShoppingCartDTO
	// 模型绑定
	if err := ctx.ShouldBindJSON(&data); err != nil {
		panic(errs.ServerInternalError)
	}
	// 日志打印
	log.Printf("添加购物车: %v", data)
	// 调用service层进行处理
	service.AddShoppingCart(&data, ctx)
	// 创建统一返回结果
	var result model.Result[string]
	// 响应
	ctx.JSON(http.StatusOK, result.Success())
}

// QueryShoppingCart 查看购物车
func QueryShoppingCart(ctx *gin.Context) {
	// 调用service层进行处理
	shoppingCartList := service.QueryShoppingCart(ctx)
	// 日志打印
	log.Printf("查看购物车: %v", *shoppingCartList)
	// 创建统一返回结果
	var result model.Result[[]entity.ShoppingCart]
	// 响应
	ctx.JSON(http.StatusOK, result.SuccessByData(*shoppingCartList))
}

// CleanShoppingCart 清空购物车
func CleanShoppingCart(ctx *gin.Context) {
	// 调用service层进行处理
	userId := service.CleanShoppingCart(ctx)
	// 日志打印
	log.Printf("清空购物车: [%v]", userId)
	// 创建统一返回结果
	var result model.Result[string]
	// 响应
	ctx.JSON(http.StatusOK, result.Success())
}

// SubShoppingCart 减少购物车某商品
func SubShoppingCart(ctx *gin.Context) {
	// 创建接收数据模型
	var data dto.ShoppingCartDTO
	// 模型绑定
	if err := ctx.ShouldBindJSON(&data); err != nil {
		panic(errs.ServerInternalError)
	}
	// 日志打印
	log.Printf("减少购物车某商品: %v", data)
	// 调用service层进行处理
	service.SubShoppingCart(&data, ctx)
	// 创建统一返回结果
	var result model.Result[string]
	// 响应
	ctx.JSON(http.StatusOK, result.Success())
}
