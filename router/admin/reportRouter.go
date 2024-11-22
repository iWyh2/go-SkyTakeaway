package admin

import (
	"github.com/gin-gonic/gin"
	controller "go-SkyTakeaway/controller/admin"
	"go-SkyTakeaway/middleware"
)

// 报表路由
func reportRouter(r *gin.Engine) {
	// 路由分组
	report := r.Group("/admin/report")
	// 使用JWT中间件进行登录校验
	report.Use(middleware.JwtAdmin)
	// 注册路由
	{
		// 营业额数据统计
		report.GET("turnoverStatistics", controller.TurnoverStatistics)
		// 用户统计
		report.GET("userStatistics", controller.UserStatistics)
		// 订单统计
		report.GET("ordersStatistics", controller.ReportOrderStatistics)
		// 销量排名
		report.GET("top10", controller.Top10Statistics)
		// 导出运营数据Excel报表
		report.GET("export", controller.ExportExcel)
	}
}
