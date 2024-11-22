package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"go-SkyTakeaway/common/constant"
	errs "go-SkyTakeaway/common/errors"
	"go-SkyTakeaway/model/vo"
	"strconv"
	"strings"
	"time"
)

// TurnoverStatistics 营业额数据统计
func TurnoverStatistics(begin, end string) *vo.TurnoverReportVO {
	// 创建数据容器
	dateList := make([]string, 0)
	turnoverList := make([]string, 0)
	// 解析时间区间
	beginTime, _ := time.Parse(time.DateOnly, begin)
	endTime, _ := time.Parse(time.DateOnly, end)
	// 从起始时间遍历至终止时间
	for !beginTime.After(endTime) {
		// 解析这一天最小时间 00:00:00
		beginTIME := time.Date(beginTime.Year(),
			beginTime.Month(), beginTime.Day(), 0, 0, 0, 0, time.Local)
		// 解析这一天最大时间 23:59:59.999999999
		endTIME := time.Date(beginTime.Year(),
			beginTime.Month(), beginTime.Day(), 23, 59, 59, 999999999, time.Local)
		// 获取每日营业额
		turnover := GetDailyTurnover(beginTIME, endTIME)
		// 添加日期
		dateList = append(dateList, beginTime.String())
		// 添加营业额
		turnoverList = append(turnoverList, strconv.FormatFloat(turnover, 'f', 2, 64))
		// 日期+1
		beginTime = beginTime.AddDate(0, 0, 1)
	}
	// 处理数据
	for i := 0; i < len(dateList); i++ {
		dateList[i] = strings.Split(dateList[i], " ")[0]
	}
	return &vo.TurnoverReportVO{
		DateList:     strings.Join(dateList, ","),
		TurnoverList: strings.Join(turnoverList, ","),
	}
}

// UserStatistics 用户统计
func UserStatistics(begin, end string) *vo.UserReportVO {
	// 创建数据容器
	dateList := make([]string, 0)
	totalUserList := make([]string, 0)
	newUserList := make([]string, 0)
	// 解析时间区间
	beginTime, _ := time.Parse(time.DateOnly, begin)
	endTime, _ := time.Parse(time.DateOnly, end)
	// 从起始时间遍历至终止时间
	for !beginTime.After(endTime) {
		// 解析这一天最小时间 00:00:00
		beginTIME := time.Date(beginTime.Year(),
			beginTime.Month(), beginTime.Day(), 0, 0, 0, 0, time.Local)
		// 解析这一天最大时间 23:59:59.999999999
		endTIME := time.Date(beginTime.Year(),
			beginTime.Month(), beginTime.Day(), 23, 59, 59, 999999999, time.Local)
		// 获取总用户数
		totalUser := GetUserCount(time.Time{}, endTIME)
		// 获取新增用户数
		newUser := GetUserCount(beginTIME, endTIME)
		// 添加日期
		dateList = append(dateList, beginTime.String())
		// 添加该日新增用户数 / 总用户数
		totalUserList = append(totalUserList, strconv.Itoa(totalUser))
		newUserList = append(newUserList, strconv.Itoa(newUser))
		// 日期+1
		beginTime = beginTime.AddDate(0, 0, 1)
	}
	// 处理数据
	for i := 0; i < len(dateList); i++ {
		dateList[i] = strings.Split(dateList[i], " ")[0]
	}
	return &vo.UserReportVO{
		DateList:      strings.Join(dateList, ","),
		TotalUserList: strings.Join(totalUserList, ","),
		NewUserList:   strings.Join(newUserList, ","),
	}
}

// ReportOrderStatistics 订单统计
func ReportOrderStatistics(begin, end string) *vo.OrderReportVO {
	// 创建数据容器
	dateList := make([]string, 0)
	orderCountList := make([]string, 0)
	validOrderCountList := make([]string, 0)
	totalCount, validCount := 0, 0
	// 解析时间区间
	beginTime, _ := time.Parse(time.DateOnly, begin)
	endTime, _ := time.Parse(time.DateOnly, end)
	// 从起始时间遍历至终止时间
	for !beginTime.After(endTime) {
		// 解析这一天最小时间 00:00:00
		beginTIME := time.Date(beginTime.Year(),
			beginTime.Month(), beginTime.Day(), 0, 0, 0, 0, time.Local)
		// 解析这一天最大时间 23:59:59.999999999
		endTIME := time.Date(beginTime.Year(),
			beginTime.Month(), beginTime.Day(), 23, 59, 59, 999999999, time.Local)
		// 查询每天的总订单数
		orderCount := GetDailyOrderCount(beginTIME, endTIME, 0)
		totalCount += orderCount
		// 查询每天的有效订单数
		validOrderCount := GetDailyOrderCount(beginTIME, endTIME, constant.Completed)
		validCount += validOrderCount
		// 添加日期
		dateList = append(dateList, beginTime.String())
		// 添加该日总订单数 / 有效订单数
		orderCountList = append(orderCountList, strconv.Itoa(orderCount))
		validOrderCountList = append(validOrderCountList, strconv.Itoa(validOrderCount))
		// 日期+1
		beginTime = beginTime.AddDate(0, 0, 1)
	}
	// 处理数据
	for i := 0; i < len(dateList); i++ {
		dateList[i] = strings.Split(dateList[i], " ")[0]
	}
	var orderCompletionRate float64
	if totalCount == 0 {
		orderCompletionRate = 0
	} else {
		orderCompletionRate = float64(validCount) / float64(totalCount)
	}
	return &vo.OrderReportVO{
		DateList:            strings.Join(dateList, ","),
		OrderCountList:      strings.Join(orderCountList, ","),
		ValidOrderCountList: strings.Join(validOrderCountList, ","),
		TotalOrderCount:     totalCount,
		ValidOrderCount:     validCount,
		OrderCompletionRate: orderCompletionRate,
	}
}

// Top10Statistics 销量排名
func Top10Statistics(begin, end string) *vo.SalesTop10ReportVO {
	// 创建数据容器
	nameList := make([]string, 0)
	numberList := make([]string, 0)
	// 解析时间区间
	beginTime, _ := time.Parse(time.DateOnly, begin)
	endTime, _ := time.Parse(time.DateOnly, end)
	// 查询前十商品
	top10Goods := GetSalesTop10(beginTime, endTime)
	// 处理数据
	for _, goods := range top10Goods {
		nameList = append(nameList, goods.Name)
		numberList = append(numberList, strconv.Itoa(goods.Number))
	}
	return &vo.SalesTop10ReportVO{
		NameList:   strings.Join(nameList, ","),
		NumberList: strings.Join(numberList, ","),
	}
}

// ExportExcel 导出运营数据Excel报表
func ExportExcel(ctx *gin.Context) {
	// 查询概览运营数据，提供给Excel模板文件
	begin := time.Now().AddDate(0, 0, -30)
	beginTime := time.Date(begin.Year(), begin.Month(), begin.Day(),
		0, 0, 0, 0, time.Local)
	endTime := time.Date(time.Now().AddDate(0, 0, -1).Year(),
		time.Now().AddDate(0, 0, -1).Month(),
		time.Now().AddDate(0, 0, -1).Day(),
		23, 59, 59, 999999999, time.Local)
	businessData := GetBusinessData(beginTime, endTime)
	// 基于提供好的模板文件创建一个新的Excel表格对象
	excel, err := excelize.OpenFile("./template/运营数据报表模板.xlsx")
	if err != nil {
		panic(errs.ServerInternalError)
	}
	// 关闭文件流
	defer excel.Close()
	// 获得第2行，填写时间
	if err := excel.SetCellValue("Sheet1", "B2",
		beginTime.Format("2006-01-02")+"至"+
			endTime.Format("2006-01-02")); err != nil {
		panic(errs.ServerInternalError)
	}
	// 创建样式：水平居中 + 垂直居中
	styleID, err := excel.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
	})
	// 应用样式到单元格 B2
	if err := excel.SetCellStyle("Sheet1", "B2",
		"B2", styleID); err != nil {
		panic(errs.ServerInternalError)
	}
	// 获得第4行，填写数据
	if err := excel.SetCellValue("Sheet1", "C4",
		businessData.Turnover); err != nil {
		panic(errs.ServerInternalError)
	}
	if err := excel.SetCellValue("Sheet1", "E4",
		businessData.OrderCompletionRate); err != nil {
		panic(errs.ServerInternalError)
	}
	if err := excel.SetCellValue("Sheet1", "G4",
		businessData.NewUsers); err != nil {
		panic(errs.ServerInternalError)
	}
	// 获得第5行，填写数据
	if err := excel.SetCellValue("Sheet1", "C5",
		businessData.ValidOrderCount); err != nil {
		panic(errs.ServerInternalError)
	}
	if err := excel.SetCellValue("Sheet1", "E5",
		businessData.UnitPrice); err != nil {
		panic(errs.ServerInternalError)
	}
	// 填写详细数据
	for i := 0; i < 30; i++ {
		// 准备明细数据
		date := begin.AddDate(0, 0, i)
		beginTime = time.Date(date.Year(), date.Month(), date.Day(),
			0, 0, 0, 0, time.Local)
		endTime = time.Date(date.Year(), date.Month(),
			date.Day(), 23, 59, 59, 999999999, time.Local)
		data := GetBusinessData(beginTime, endTime)
		// 从第 8 行开始写入
		rowNum := 8 + i
		// 填写数据
		_ = excel.SetCellValue("Sheet1", fmt.Sprintf("B%d", rowNum), date.Format("2006-01-02"))
		_ = excel.SetCellValue("Sheet1", fmt.Sprintf("C%d", rowNum), data.Turnover)
		_ = excel.SetCellValue("Sheet1", fmt.Sprintf("D%d", rowNum), data.ValidOrderCount)
		_ = excel.SetCellValue("Sheet1", fmt.Sprintf("E%d", rowNum), data.OrderCompletionRate)
		_ = excel.SetCellValue("Sheet1", fmt.Sprintf("F%d", rowNum), data.UnitPrice)
		_ = excel.SetCellValue("Sheet1", fmt.Sprintf("G%d", rowNum), data.NewUsers)
	}
	// 传输到浏览器让管理员下载
	// 前端控制了下载名称时，后端只需要设置 Content-Type
	ctx.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	// 将 Excel 数据直接写入响应流
	if err := excel.Write(ctx.Writer); err != nil {
		panic(errs.ServerInternalError)
	}
}
