package admin

import (
	"github.com/gin-gonic/gin"
	controller "go-SkyTakeaway/controller/admin"
	"go-SkyTakeaway/middleware"
)

// 套餐路由
func setmealRouter(r *gin.Engine) {
	// 路由分组
	setmeal := r.Group("/admin/setmeal")
	// 使用JWT中间件进行登录校验
	setmeal.Use(middleware.JwtAdmin)
	// 注册路由
	{
		// 新增套餐
		setmeal.POST("", controller.AddSetmeal)
		// 套餐分页查询
		setmeal.GET("/page", controller.SetmealPageQuery)
		// 删除套餐
		setmeal.DELETE("", controller.DeleteSetmealByIDs)
		// 根据id查询套餐
		setmeal.GET("/:id", controller.QuerySetmealByID)
		// 修改套餐
		setmeal.PUT("", controller.UpdateSetmeal)
		// 起售/停售套餐
		setmeal.POST("/status/:status", controller.StartOrStopSetmeal)
	}
}
