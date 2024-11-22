package user

import (
	"github.com/gin-gonic/gin"
	controller "go-SkyTakeaway/controller/user"
	"go-SkyTakeaway/middleware"
)

// 微信用户路由
func wechatUserRouter(r *gin.Engine) {
	// 路由分组
	user := r.Group("/user/user")
	// 用户登录路由不需要登录校验
	user.POST("login", controller.UserLogin)
	// 使用JWT中间件进行登录校验
	user.Use(middleware.JwtUser)
	// 注册其他路由
	{
	}
}
