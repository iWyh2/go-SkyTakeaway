package user

import (
	"github.com/gin-gonic/gin"
	controller "go-SkyTakeaway/controller/user"
	"go-SkyTakeaway/middleware"
)

// 套餐路由
func setmealRouter(r *gin.Engine) {
	// 路由分组
	setmeal := r.Group("/user/setmeal")
	// 使用JWT中间件进行登录校验
	setmeal.Use(middleware.JwtUser)
	// 注册路由
	{
		// 根据分类id查询套餐
		setmeal.GET("/list", controller.UserQuerySetmealByCategoryID)
		// 根据套餐id查询包含的菜品列表
		setmeal.GET("/dish/:id", controller.UserQueryDishBySetmealID)
	}
}
