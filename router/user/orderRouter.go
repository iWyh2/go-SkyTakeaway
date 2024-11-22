package user

import (
	"github.com/gin-gonic/gin"
	controller "go-SkyTakeaway/controller/user"
	"go-SkyTakeaway/middleware"
)

// 订单路由
func orderRouter(r *gin.Engine) {
	// 路由分组
	order := r.Group("/user/order")
	// 使用JWT中间件进行登录校验
	order.Use(middleware.JwtUser)
	// 注册路由
	{
		// 用户下单
		order.POST("submit", controller.OrderSubmit)
		// 订单支付
		order.PUT("payment", controller.OrderPayment)
		// 根据订单id查询订单详情
		order.GET("orderDetail/:id", controller.OrderDetail)
		// 查询历史订单
		order.GET("historyOrders", controller.HistoryOrders)
		// 用户取消订单
		order.PUT("cancel/:id", controller.CancelOrder)
		// 再来一单
		order.POST("repetition/:id", controller.RepetitionOrder)
		// 用户催单
		order.GET("reminder/:id", controller.OrderReminder)
	}
}
