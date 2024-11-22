package userController

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

// OrderSubmit 用户下单
func OrderSubmit(ctx *gin.Context) {
	// 创建接收数据模型
	var data dto.OrderSubmitDTO
	// 模型绑定
	if err := ctx.ShouldBindJSON(&data); err != nil {
		panic(errs.ServerInternalError)
	}
	// 日志打印
	log.Printf("用户下单: %v", data)
	// 调用service层进行处理
	orderVO := service.OrderSubmit(&data, ctx)
	// 创建统一返回结果
	var result model.Result[vo.OrderSubmitVO]
	// 响应
	ctx.JSON(http.StatusOK, result.SuccessByData(*orderVO))
}

// OrderPayment 订单支付
func OrderPayment(ctx *gin.Context) {
	// 创建接收数据模型
	var data dto.OrderPaymentDTO
	// 模型绑定
	if err := ctx.ShouldBindJSON(&data); err != nil {
		panic(errs.ServerInternalError)
	}
	// 日志打印
	log.Printf("订单支付: %v", data)
	// 调用service层进行处理
	paymentVO := service.OrderPayment(&data, ctx)
	// 创建统一返回结果
	var result model.Result[vo.OrderPaymentVO]
	// 响应
	ctx.JSON(http.StatusOK, result.SuccessByData(*paymentVO))
}

// OrderDetail 根据订单id查询订单详情
func OrderDetail(ctx *gin.Context) {
	orderId := ctx.Param("id")
	// 日志打印
	log.Printf("根据订单id查询订单详情: [%v]", orderId)
	// 调用service层进行处理
	orderVO := service.OrderDetail(orderId)
	// 创建统一返回结果
	var result model.Result[vo.OrderVO]
	// 响应
	ctx.JSON(http.StatusOK, result.SuccessByData(*orderVO))
}

// HistoryOrders 查询历史订单
func HistoryOrders(ctx *gin.Context) {
	// 调用service层进行处理
	page := service.HistoryOrders(ctx)
	// 日志打印
	log.Printf("查询历史订单: [%v]", page.Total)
	// 创建统一返回结果
	var result model.Result[pageResult.Page[vo.OrderVO]]
	// 响应
	ctx.JSON(http.StatusOK, result.SuccessByData(*page))
}

// CancelOrder 用户取消订单
func CancelOrder(ctx *gin.Context) {
	// 接收路径参数
	orderId := ctx.Param("id")
	// 日志打印
	log.Printf("取消订单: [%v]", orderId)
	// 调用service层进行处理
	service.CancelOrder(orderId)
	// 创建统一返回结果
	var result model.Result[string]
	// 响应
	ctx.JSON(http.StatusOK, result.Success())
}

// RepetitionOrder 再来一单
func RepetitionOrder(ctx *gin.Context) {
	// 接收路径参数
	orderId := ctx.Param("id")
	// 日志打印
	log.Printf("再来一单: [%v]", orderId)
	// 调用service层进行处理
	service.RepetitionOrder(orderId, ctx)
	// 创建统一返回结果
	var result model.Result[string]
	// 响应
	ctx.JSON(http.StatusOK, result.Success())
}

// OrderReminder 用户催单
func OrderReminder(ctx *gin.Context) {
	// 接收路径参数
	orderId := ctx.Param("id")
	// 日志打印
	log.Printf("用户催单: [%v]", orderId)
	// 调用service层进行处理
	service.OrderReminder(orderId)
	// 创建统一返回结果
	var result model.Result[string]
	// 响应
	ctx.JSON(http.StatusOK, result.Success())
}
