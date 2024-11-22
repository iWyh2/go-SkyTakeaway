package user

import (
	"github.com/gin-gonic/gin"
	controller "go-SkyTakeaway/controller/user"
	"go-SkyTakeaway/middleware"
)

// 购物车路由
func shoppingCartRouter(r *gin.Engine) {
	// 路由分组
	shoppingCart := r.Group("/user/shoppingCart")
	// 使用JWT中间件进行登录校验
	shoppingCart.Use(middleware.JwtUser)
	// 注册路由
	{
		// 添加购物车
		shoppingCart.POST("add", controller.AddShoppingCart)
		// 查看购物车
		shoppingCart.GET("list", controller.QueryShoppingCart)
		// 清空购物车
		shoppingCart.DELETE("clean", controller.CleanShoppingCart)
		// 减少购物车某项
		shoppingCart.POST("sub", controller.SubShoppingCart)
	}
}
