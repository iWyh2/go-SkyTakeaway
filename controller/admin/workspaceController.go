package adminController

import (
	"github.com/gin-gonic/gin"
	model "go-SkyTakeaway/model/result"
	"go-SkyTakeaway/model/vo"
	"go-SkyTakeaway/service"
	"log"
	"net/http"
)

// TodayBusinessData 今日数据
func TodayBusinessData(ctx *gin.Context) {
	// 调用service层进行处理
	businessData := service.TodayBusinessData()
	// 日志打印
	log.Printf("今日数据: %v", *businessData)
	// 创建统一返回结果
	var result model.Result[vo.BusinessDataVO]
	// 响应
	ctx.JSON(http.StatusOK, result.SuccessByData(*businessData))
}

// OverviewOrders 订单管理
func OverviewOrders(ctx *gin.Context) {
	// 调用service层进行处理
	ordersData := service.OverviewOrders()
	// 日志打印
	log.Printf("订单管理: %v", *ordersData)
	// 创建统一返回结果
	var result model.Result[vo.OrderOverViewVO]
	// 响应
	ctx.JSON(http.StatusOK, result.SuccessByData(*ordersData))
}

// OverviewDishes 菜品总览
func OverviewDishes(ctx *gin.Context) {
	// 调用service层进行处理
	dishesData := service.OverviewDishes()
	// 日志打印
	log.Printf("菜品总览: %v", *dishesData)
	// 创建统一返回结果
	var result model.Result[vo.DishOverViewVO]
	// 响应
	ctx.JSON(http.StatusOK, result.SuccessByData(*dishesData))
}

// OverviewSetmeals 套餐总览
func OverviewSetmeals(ctx *gin.Context) {
	// 调用service层进行处理
	setmealsData := service.OverviewSetmeals()
	// 日志打印
	log.Printf("套餐总览: %v", *setmealsData)
	// 创建统一返回结果
	var result model.Result[vo.SetmealOverViewVO]
	// 响应
	ctx.JSON(http.StatusOK, result.SuccessByData(*setmealsData))
}
