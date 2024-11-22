package user

import (
	"github.com/gin-gonic/gin"
	controller "go-SkyTakeaway/controller/user"
	"go-SkyTakeaway/middleware"
)

// 菜品路由
func dishRouter(r *gin.Engine) {
	// 路由分组
	dish := r.Group("/user/dish")
	// 使用JWT中间件进行登录校验
	dish.Use(middleware.JwtUser)
	// 注册路由
	{
		// 根据分类id查询菜品
		dish.GET("/list", controller.QueryDishByCategoryID)
	}
}
