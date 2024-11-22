package adminController

import (
	"github.com/gin-gonic/gin"
	model "go-SkyTakeaway/model/result"
	"go-SkyTakeaway/model/vo"
	"go-SkyTakeaway/service"
	"log"
	"net/http"
)

// TurnoverStatistics 营业额数据统计
func TurnoverStatistics(ctx *gin.Context) {
	// 接收query参数
	begin := ctx.Query("begin")
	end := ctx.Query("end")
	// 调用service层进行处理
	turnover := service.TurnoverStatistics(begin, end)
	// 日志打印
	log.Printf("营业额数据统计: %v", turnover)
	// 创建统一返回结果
	var result model.Result[vo.TurnoverReportVO]
	// 响应
	ctx.JSON(http.StatusOK, result.SuccessByData(*turnover))
}

// UserStatistics 用户统计
func UserStatistics(ctx *gin.Context) {
	// 接收query参数
	begin := ctx.Query("begin")
	end := ctx.Query("end")
	// 调用service层进行处理
	userStatistics := service.UserStatistics(begin, end)
	// 日志打印
	log.Printf("用户统计: %v", userStatistics)
	// 创建统一返回结果
	var result model.Result[vo.UserReportVO]
	// 响应
	ctx.JSON(http.StatusOK, result.SuccessByData(*userStatistics))
}

// ReportOrderStatistics 订单统计
func ReportOrderStatistics(ctx *gin.Context) {
	// 接收query参数
	begin := ctx.Query("begin")
	end := ctx.Query("end")
	// 调用service层进行处理
	orderStatistics := service.ReportOrderStatistics(begin, end)
	// 日志打印
	log.Printf("订单统计: %v", orderStatistics)
	// 创建统一返回结果
	var result model.Result[vo.OrderReportVO]
	// 响应
	ctx.JSON(http.StatusOK, result.SuccessByData(*orderStatistics))
}

// Top10Statistics 销量排名
func Top10Statistics(ctx *gin.Context) {
	// 接收query参数
	begin := ctx.Query("begin")
	end := ctx.Query("end")
	// 调用service层进行处理
	top10 := service.Top10Statistics(begin, end)
	// 日志打印
	log.Printf("销量排名: %v", top10)
	// 创建统一返回结果
	var result model.Result[*vo.SalesTop10ReportVO]
	// 响应
	ctx.JSON(http.StatusOK, result.SuccessByData(top10))
}

// ExportExcel 导出运营数据Excel报表
func ExportExcel(ctx *gin.Context) {
	// 调用service层进行处理
	service.ExportExcel(ctx)
	// 日志打印
	log.Printf("导出运营数据Excel报表")
}
