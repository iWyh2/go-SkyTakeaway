package adminController

import (
	"github.com/gin-gonic/gin"
	errs "go-SkyTakeaway/common/errors"
	"go-SkyTakeaway/model/dto"
	model "go-SkyTakeaway/model/result"
	pageResult "go-SkyTakeaway/model/result/page"
	"go-SkyTakeaway/model/vo"
	"go-SkyTakeaway/service"
	"log"
	"net/http"
)

// OrderConditionSearch 订单搜索
func OrderConditionSearch(ctx *gin.Context) {
	// 创建接收数据模型
	var data dto.OrderPageQueryDTO
	// 模型绑定
	if err := ctx.ShouldBindQuery(&data); err != nil {
		panic(errs.ServerInternalError)
	}
	// 日志打印
	log.Printf("订单搜索: %v", data)
	// 调用service层进行处理
	page := service.OrderConditionSearch(&data)
	// 创建统一返回结果
	var result model.Result[pageResult.Page[vo.OrderVO]]
	// 响应
	ctx.JSON(http.StatusOK, result.SuccessByData(*page))
}

// OrderStatistics 各个状态的订单数量统计
func OrderStatistics(ctx *gin.Context) {
	// 调用service层进行处理
	orderStatisticsVO := service.OrderStatistics()
	// 日志打印
	log.Printf("各个状态的订单数量统计: %v", *orderStatisticsVO)
	// 创建统一返回结果
	var result model.Result[vo.OrderStatisticsVO]
	// 响应
	ctx.JSON(http.StatusOK, result.SuccessByData(*orderStatisticsVO))
}

// OrderDetail 查询订单详情
func OrderDetail(ctx *gin.Context) {
	// 接收路径参数
	orderId := ctx.Param("id")
	// 日志打印
	log.Printf("查询订单详情: [%v]", orderId)
	// 调用service层进行处理
	orderVO := service.OrderDetail(orderId)
	// 创建统一返回结果
	var result model.Result[vo.OrderVO]
	// 响应
	ctx.JSON(http.StatusOK, result.SuccessByData(*orderVO))
}

// OrderConfirm 接单
func OrderConfirm(ctx *gin.Context) {
	// 创建接收数据模型
	var data dto.OrderConfirmDTO
	// 模型绑定
	if err := ctx.ShouldBindJSON(&data); err != nil {
		panic(errs.ServerInternalError)
	}
	// 日志打印
	log.Printf("接单: [%v]", data.OrderId)
	// 调用service层进行处理
	service.OrderConfirm(&data)
	// 创建统一返回结果
	var result model.Result[string]
	// 响应
	ctx.JSON(http.StatusOK, result.Success())
}

// OrderRejection 拒单
func OrderRejection(ctx *gin.Context) {
	// 创建接收数据模型
	var data dto.OrderRejectionDTO
	// 模型绑定
	if err := ctx.ShouldBindJSON(&data); err != nil {
		panic(errs.ServerInternalError)
	}
	// 日志打印
	log.Printf("拒单: [%v]", data.OrderId)
	// 调用service层进行处理
	service.OrderRejection(&data)
	// 创建统一返回结果
	var result model.Result[string]
	// 响应
	ctx.JSON(http.StatusOK, result.Success())
}

// CancelOrder 商家取消订单
func CancelOrder(ctx *gin.Context) {
	// 创建接收数据模型
	var data dto.OrderCancelDTO
	// 模型绑定
	if err := ctx.ShouldBindJSON(&data); err != nil {
		panic(errs.ServerInternalError)
	}
	// 日志打印
	log.Printf("商家取消订单: [%v]", data.OrderId)
	// 调用service层进行处理
	service.CancelOrderByBusiness(&data)
	// 创建统一返回结果
	var result model.Result[string]
	// 响应
	ctx.JSON(http.StatusOK, result.Success())
}

// OrderDelivery 订单派送
func OrderDelivery(ctx *gin.Context) {
	// 接收路径参数
	orderId := ctx.Param("id")
	// 日志打印
	log.Printf("订单派送: [%v]", orderId)
	// 调用service层进行处理
	service.OrderDelivery(orderId)
	// 创建统一返回结果
	var result model.Result[string]
	// 响应
	ctx.JSON(http.StatusOK, result.Success())
}

// OrderComplete 完成订单
func OrderComplete(ctx *gin.Context) {
	// 接收路径参数
	orderId := ctx.Param("id")
	// 日志打印
	log.Printf("完成订单: [%v]", orderId)
	// 调用service层进行处理
	service.OrderComplete(orderId)
	// 创建统一返回结果
	var result model.Result[string]
	// 响应
	ctx.JSON(http.StatusOK, result.Success())
}
