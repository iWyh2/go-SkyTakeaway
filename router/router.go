package router

import (
	"github.com/gin-gonic/gin"
	"go-SkyTakeaway/middleware"
	"go-SkyTakeaway/router/admin"
	"go-SkyTakeaway/router/user"
	"go-SkyTakeaway/router/websocket"
)

// Router 返回系统总路由
func Router() *gin.Engine {
	// 创建默认路由
	r := gin.Default()
	// 使用全局错误处理中间件
	r.Use(middleware.Recover)
	// 启用websocket服务
	websocket.WSRouter(r)
	// 注册管理端路由
	admin.RouterAdmin(r)
	// 注册用户端路由
	user.RouterUser(r)
	return r
}
