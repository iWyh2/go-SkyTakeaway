package user

import "github.com/gin-gonic/gin"

// RouterUser 注册用户端路由
func RouterUser(r *gin.Engine) {
	// 注册店铺路由
	shopRouter(r)
	// 注册微信用户路由
	wechatUserRouter(r)
	// 注册分类路由
	categoryRouter(r)
	// 注册菜品路由
	dishRouter(r)
	// 注册套餐路由
	setmealRouter(r)
	// 注册购物车路由
	shoppingCartRouter(r)
	// 注册地址簿路由
	addressBookRouter(r)
	// 注册订单路由
	orderRouter(r)
}
