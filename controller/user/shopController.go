package userController

import (
	"github.com/gin-gonic/gin"
	model "go-SkyTakeaway/model/result"
	"go-SkyTakeaway/service"
	"log"
	"net/http"
	"strconv"
)

// GetShopStatus 查询店铺营业状态
func GetShopStatus(ctx *gin.Context) {
	// 调用service层处理
	data := service.GetShopStatus()
	// 处理得到的数据
	status, _ := strconv.Atoi(data)
	// 日志打印
	log.Printf("查询店铺营业状态: [%v]", status)
	// 创建统一返回结果
	var result model.Result[int]
	// 响应
	ctx.JSON(http.StatusOK, result.SuccessByData(status))
}
