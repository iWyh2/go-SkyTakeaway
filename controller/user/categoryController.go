package userController

import (
	"github.com/gin-gonic/gin"
	model "go-SkyTakeaway/model/result"
	"go-SkyTakeaway/model/vo"
	"go-SkyTakeaway/service"
	"log"
	"net/http"
)

// QueryCategoryByType 根据类型查询分类
func QueryCategoryByType(ctx *gin.Context) {
	// 接收query传递的数据
	typeData := ctx.Query("type")
	// 日志打印
	log.Printf("根据类型查询分类: [%v]", typeData)
	// 调用service层处理
	categories := service.UserQueryCategoryByType(typeData)
	// 创建统一返回结果
	var result model.Result[[]vo.CategoryVO]
	// 响应
	ctx.JSON(http.StatusOK, result.SuccessByData(*categories))
}
