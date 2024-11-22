package admin

import (
	"github.com/gin-gonic/gin"
	controller "go-SkyTakeaway/controller/admin"
	"go-SkyTakeaway/middleware"
)

// 员工路由
func employeeRouter(r *gin.Engine) {
	// 注册路由，员工登录，无需jwt校验
	r.POST("/admin/employee/login", controller.Login)
	// 路由分组，需校验jwt
	employee := r.Group("/admin/employee")
	// 使用JWT中间件进行登录校验
	employee.Use(middleware.JwtAdmin)
	// 注册路由
	{
		// 员工退出
		employee.POST("/logout", controller.Logout)
		// 新增员工
		employee.POST("", controller.Register)
		// 分页查询员工
		employee.GET("/page", controller.EmpPageQuery)
		// 启用禁用员工账号
		employee.POST("/status/:status", controller.StartOrStopEmp)
		// 根据id查询员工信息
		employee.GET("/:id", controller.EmpQueryByID)
		// 编辑员工信息
		employee.PUT("", controller.EmpUpdate)
	}
}
