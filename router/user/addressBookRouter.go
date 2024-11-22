package user

import (
	"github.com/gin-gonic/gin"
	controller "go-SkyTakeaway/controller/user"
	"go-SkyTakeaway/middleware"
)

// 地址簿路由
func addressBookRouter(r *gin.Engine) {
	// 路由分组
	addressBook := r.Group("/user/addressBook")
	// 使用JWT中间件进行登录校验
	addressBook.Use(middleware.JwtUser)
	// 注册路由
	{
		// 新增地址
		addressBook.POST("", controller.AddAddressBook)
		// 查询登录用户所有地址
		addressBook.GET("list", controller.QueryAddressBooks)
		// 查询默认地址
		addressBook.GET("default", controller.GetDefaultAddressBook)
		// 设置默认地址
		addressBook.PUT("default", controller.SetDefaultAddressBook)
		// 根据id修改地址
		addressBook.PUT("", controller.UpdateAddressBookById)
		// 根据id删除地址
		addressBook.DELETE("", controller.DeleteAddressBookById)
		// 根据id查询地址
		addressBook.GET(":id", controller.QueryAddressBookById)
	}
}
