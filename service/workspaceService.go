package service

import (
	"go-SkyTakeaway/common/constant"
	"go-SkyTakeaway/model/vo"
	"time"
)

// TodayBusinessData 今日数据
func TodayBusinessData() *vo.BusinessDataVO {
	// 今日时间区间
	beginTime := time.Date(time.Now().Year(), time.Now().Month(),
		time.Now().Day(), 0, 0, 0, 0, time.Local)
	endTime := time.Date(time.Now().Year(), time.Now().Month(),
		time.Now().Day(), 23, 59, 59, 999999999, time.Local)
	return GetBusinessData(beginTime, endTime)
}

// GetBusinessData 获取营业数据
func GetBusinessData(beginTime, endTime time.Time) *vo.BusinessDataVO {
	// 查询总订单数
	totalOrderCount := GetDailyOrderCount(beginTime, endTime, 0)
	// 营业额
	turnover := GetDailyTurnover(beginTime, endTime)
	// 有效订单数
	validOrderCount := GetDailyOrderCount(beginTime, endTime, constant.Completed)
	// 订单完成率
	var orderCompletionRate float64
	// 平均客单价
	var unitPrice float64
	if totalOrderCount != 0 && validOrderCount != 0 {
		orderCompletionRate = float64(validOrderCount) / float64(totalOrderCount)
		unitPrice = turnover / float64(validOrderCount)
	}
	// 新增用户数
	newUsers := GetUserCount(beginTime, endTime)
	return &vo.BusinessDataVO{
		Turnover:            turnover,
		ValidOrderCount:     validOrderCount,
		OrderCompletionRate: orderCompletionRate,
		UnitPrice:           unitPrice,
		NewUsers:            newUsers,
	}
}

// OverviewOrders 订单管理
func OverviewOrders() *vo.OrderOverViewVO {
	// 今日时间区间
	beginTime := time.Date(time.Now().Year(), time.Now().Month(),
		time.Now().Day(), 0, 0, 0, 0, time.Local)
	endTime := time.Date(time.Now().Year(), time.Now().Month(),
		time.Now().Day(), 23, 59, 59, 999999999, time.Local)
	// 待接单订单数
	waitingOrders := GetDailyOrderCount(beginTime, endTime, constant.ToBeConfirmed)
	// 待派送订单数
	deliveredOrders := GetDailyOrderCount(beginTime, endTime, constant.Confirmed)
	// 已完成订单数
	completedOrders := GetDailyOrderCount(beginTime, endTime, constant.Completed)
	// 已取消订单数
	cancelledOrders := GetDailyOrderCount(beginTime, endTime, constant.Cancelled)
	// 全部订单订单数
	allOrders := GetDailyOrderCount(beginTime, endTime, 0)
	return &vo.OrderOverViewVO{
		WaitingOrders:   waitingOrders,
		DeliveredOrders: deliveredOrders,
		CompletedOrders: completedOrders,
		CancelledOrders: cancelledOrders,
		AllOrders:       allOrders,
	}
}

// OverviewDishes 菜品总览
func OverviewDishes() *vo.DishOverViewVO {
	sold := GetDishCount(constant.Enable)
	discontinued := GetDishCount(constant.Disable)
	return &vo.DishOverViewVO{
		Sold:         sold,
		Discontinued: discontinued,
	}
}

// OverviewSetmeals 套餐总览
func OverviewSetmeals() *vo.SetmealOverViewVO {
	sold := GetSetmealCount(constant.Enable)
	discontinued := GetSetmealCount(constant.Disable)
	return &vo.SetmealOverViewVO{
		Sold:         sold,
		Discontinued: discontinued,
	}
}
