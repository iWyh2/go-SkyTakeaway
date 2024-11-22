package admin

import (
	"github.com/gin-gonic/gin"
	controller "go-SkyTakeaway/controller/admin"
	"go-SkyTakeaway/middleware"
)

// 菜品路由
func dishRouter(r *gin.Engine) {
	// 路由分组
	dish := r.Group("/admin/dish")
	// 使用JWT中间件进行登录校验
	dish.Use(middleware.JwtAdmin)
	// 注册路由
	{
		// 新增菜品
		dish.POST("", controller.AddDishWithFlavor)
		// 菜品分页查询
		dish.GET("/page", controller.DishPageQuery)
		// 删除菜品
		dish.DELETE("", controller.DeleteDishByIDs)
		// 根据id查询菜品
		dish.GET("/:id", controller.QueryDishByID)
		// 修改菜品
		dish.PUT("", controller.UpdateDish)
		// 起售/停售菜品
		dish.POST("/status/:status", controller.StartOrStopDish)
		// 根据分类id查询菜品
		dish.GET("/list", controller.QueryDishByCategoryID)
	}
}
