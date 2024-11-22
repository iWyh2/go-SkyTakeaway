package admin

import (
	"github.com/gin-gonic/gin"
	controller "go-SkyTakeaway/controller/admin"
	"go-SkyTakeaway/middleware"
)

// 文件上传路由
func uploadRouter(r *gin.Engine) {
	// 路由分组
	upload := r.Group("/admin/common/upload")
	// 使用JWT中间件进行登录校验
	upload.Use(middleware.JwtAdmin)
	// 注册路由
	{
		// 文件上传
		upload.POST("", controller.UploadFile)
	}
}
