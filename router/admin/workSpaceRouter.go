package admin

import (
	"github.com/gin-gonic/gin"
	controller "go-SkyTakeaway/controller/admin"
	"go-SkyTakeaway/middleware"
)

// 工作台路由
func workSpaceRouter(r *gin.Engine) {
	// 路由分组
	workspace := r.Group("/admin/workspace")
	// 使用JWT中间件进行登录校验
	workspace.Use(middleware.JwtAdmin)
	// 注册路由
	{
		// 今日数据
		workspace.GET("businessData", controller.TodayBusinessData)
		// 订单管理
		workspace.GET("overviewOrders", controller.OverviewOrders)
		// 菜品总览
		workspace.GET("overviewDishes", controller.OverviewDishes)
		// 套餐总览
		workspace.GET("overviewSetmeals", controller.OverviewSetmeals)
	}
}
