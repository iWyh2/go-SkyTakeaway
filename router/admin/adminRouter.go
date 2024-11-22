package admin

import (
	"github.com/gin-gonic/gin"
)

// RouterAdmin 注册管理端路由
func RouterAdmin(r *gin.Engine) {
	// 注册员工路由
	employeeRouter(r)
	// 注册分类路由
	categoryRouter(r)
	// 注册文件上传路由
	uploadRouter(r)
	// 注册菜品路由
	dishRouter(r)
	// 注册套餐路由
	setmealRouter(r)
	// 注册店铺路由
	shopRouter(r)
	// 注册订单路由
	orderRouter(r)
	// 注册报表路由
	reportRouter(r)
	// 注册工作台路由
	workSpaceRouter(r)
}
