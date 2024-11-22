package user

import (
	"github.com/gin-gonic/gin"
	controller "go-SkyTakeaway/controller/user"
)

// 店铺路由
func shopRouter(r *gin.Engine) {
	// 路由分组
	shop := r.Group("/user/shop")
	// 注册路由
	{
		// 查询营业状态
		shop.GET("status", controller.GetShopStatus)
	}
}
