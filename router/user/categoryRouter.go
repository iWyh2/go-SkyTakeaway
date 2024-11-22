package user

import (
	"github.com/gin-gonic/gin"
	controller "go-SkyTakeaway/controller/user"
	"go-SkyTakeaway/middleware"
)

// 分类路由
func categoryRouter(r *gin.Engine) {
	// 路由分组
	category := r.Group("/user/category")
	// 使用JWT中间件进行登录校验
	category.Use(middleware.JwtUser)
	// 注册路由
	{
		// 根据类型查询分类
		category.GET("/list", controller.QueryCategoryByType)
	}
}
