package userController

import (
	"github.com/gin-gonic/gin"
	model "go-SkyTakeaway/model/result"
	"go-SkyTakeaway/model/vo"
	"go-SkyTakeaway/service"
	"log"
	"net/http"
)

// UserQuerySetmealByCategoryID 根据分类id查询套餐
func UserQuerySetmealByCategoryID(ctx *gin.Context) {
	// 接收Uri传递的数据
	categoryId := ctx.Query("categoryId")
	// 日志打印
	log.Printf("根据分类id查询套餐: [%v]", categoryId)
	// 调用service层处理
	setmeals := service.UserQuerySetmealByCategoryId(categoryId)
	// 创建统一返回结果
	var result model.Result[any]
	// 响应
	ctx.JSON(http.StatusOK, result.SuccessByData(setmeals))
}

// UserQueryDishBySetmealID 根据套餐id查询具体信息列表
func UserQueryDishBySetmealID(ctx *gin.Context) {
	// 接收Uri传递的数据
	setmealId := ctx.Param("id")
	// 日志打印
	log.Printf("根据套餐id查询包含的菜品列表: [%v]", setmealId)
	// 调用service层处理
	dishes := service.UserQueryDishBySetmealID(setmealId)
	// 创建统一返回结果
	var result model.Result[[]vo.DishItemVO]
	// 响应
	ctx.JSON(http.StatusOK, result.SuccessByData(*dishes))
}
