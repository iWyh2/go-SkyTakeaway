package adminController

import (
	"github.com/gin-gonic/gin"
	model "go-SkyTakeaway/model/result"
	"go-SkyTakeaway/service"
	"log"
	"net/http"
	"strconv"
)

// SetShopStatus 设置营业状态
func SetShopStatus(ctx *gin.Context) {
	// 接收Uri传递的数据
	status := ctx.Param("status")
	// 日志打印
	log.Printf("设置店铺营业状态: [%v]", status)
	// 调用service层处理
	service.SetShopStatus(status)
	// 创建统一返回结果
	var result model.Result[string]
	// 响应
	ctx.JSON(http.StatusOK, result.Success())
}

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
