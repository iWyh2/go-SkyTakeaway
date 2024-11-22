package admin

import (
	"github.com/gin-gonic/gin"
	controller "go-SkyTakeaway/controller/admin"
	"go-SkyTakeaway/middleware"
)

// 订单路由
func orderRouter(r *gin.Engine) {
	// 路由分组
	order := r.Group("/admin/order")
	// 使用JWT中间件进行登录校验
	order.Use(middleware.JwtAdmin)
	// 注册路由
	{
		// 接单
		order.PUT("confirm", controller.OrderConfirm)
		// 拒单
		order.PUT("rejection", controller.OrderRejection)
		// 商家取消订单
		order.PUT("cancel", controller.CancelOrder)
		// 派送订单
		order.PUT("delivery/:id", controller.OrderDelivery)
		// 完成订单
		order.PUT("complete/:id", controller.OrderComplete)
		// 订单搜索
		order.GET("conditionSearch", controller.OrderConditionSearch)
		// 查询订单详情
		order.GET("details/:id", controller.OrderDetail)
		// 各个状态的订单数量统计
		order.GET("statistics", controller.OrderStatistics)
	}
}
