package admin

import (
	"github.com/gin-gonic/gin"
	controller "go-SkyTakeaway/controller/admin"
	"go-SkyTakeaway/middleware"
)

// 分类路由
func categoryRouter(r *gin.Engine) {
	// 路由分组
	category := r.Group("/admin/category")
	// 使用JWT中间件进行登录校验
	category.Use(middleware.JwtAdmin)
	// 注册路由
	{
		// 分类分页查询
		category.GET("/page", controller.CategoryPageQuery)
		// 根据类型查询分类
		category.GET("/list", controller.QueryCategoryByType)
		// 启用&禁用分类
		category.POST("/status/:status", controller.StartOrStopCategory)
		// 新增分类
		category.POST("", controller.AddCategory)
		// 根据id删除分类
		category.DELETE("", controller.DeleteCategoryByID)
		// 修改分类
		category.PUT("", controller.UpdateCategory)
	}
}
