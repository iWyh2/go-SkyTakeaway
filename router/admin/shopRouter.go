package admin

import (
	"github.com/gin-gonic/gin"
	controller "go-SkyTakeaway/controller/admin"
	"go-SkyTakeaway/middleware"
)

// 店铺路由
func shopRouter(r *gin.Engine) {
	// 路由分组
	shop := r.Group("/admin/shop")
	// 使用JWT中间件进行登录校验
	shop.Use(middleware.JwtAdmin)
	// 注册路由
	{
		// 设置营业状态
		shop.PUT(":status", controller.SetShopStatus)
		// 查询营业状态
		shop.GET("status", controller.GetShopStatus)
	}
}
