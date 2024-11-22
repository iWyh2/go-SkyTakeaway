package userController

import (
	"github.com/gin-gonic/gin"
	model "go-SkyTakeaway/model/result"
	"go-SkyTakeaway/model/vo"
	"go-SkyTakeaway/service"
	"log"
	"net/http"
)

// QueryDishByCategoryID 根据分类id查询菜品
func QueryDishByCategoryID(ctx *gin.Context) {
	// 接收query传递的数据
	categoryId := ctx.Query("categoryId")
	// 日志打印
	log.Printf("根据分类id查询菜品: [%v]", categoryId)
	// 调用service层处理
	dishes := service.UserQueryDishByCategoryID(categoryId)
	// 创建统一返回结果
	var result model.Result[[]vo.DishVO]
	// 响应
	ctx.JSON(http.StatusOK, result.SuccessByData(*dishes))
}
