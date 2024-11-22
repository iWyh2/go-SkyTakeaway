package timer

import (
	"github.com/robfig/cron/v3"
	"go-SkyTakeaway/common/constant"
	errs "go-SkyTakeaway/common/errors"
	"go-SkyTakeaway/model"
	"go-SkyTakeaway/service"
	"log"
	"time"
)

// 初始化定时器
func init() {
	log.Printf("启动定时器: [%s]", time.Now().Format("2006-01-02 15:04:05"))
	// 获得定时器
	timerTask := cron.New(cron.WithSeconds())
	// 添加定时器任务
	if _, err := timerTask.AddFunc("0 * * * * ?", processTimeoutOrder); err != nil {
		panic(errs.TimerTaskError)
	}
	if _, err := timerTask.AddFunc("0 0 1 * * ?", processDeliveryOrder); err != nil {
		panic(errs.TimerTaskError)
	}
	// 启动定时器
	timerTask.Start()
}

// 处理支付超时订单
func processTimeoutOrder() {
	// 当前时间
	nowTime := time.Now()
	log.Printf("处理支付超时订单: [%s]", time.Now().Format("2006-01-02 15:04:05"))
	// 15分钟前的时间
	duration, _ := time.ParseDuration("-1m")
	nowTime.Add(15 * duration)
	// 查询出超时订单
	ordersList := service.GetOrderByStatusAndOrderTime(constant.PendingPayment, model.LocalTime(nowTime))
	if ordersList != nil && len(ordersList) > 0 {
		for _, order := range ordersList {
			order.Status = constant.Cancelled
			order.CancelReason = "支付超时，自动取消"
			order.CancelTime = model.LocalTime(time.Now())
			service.UpdateOrder(order)
		}
	}
}

// 处理派送中订单
func processDeliveryOrder() {
	// 当前时间
	nowTime := time.Now()
	log.Printf("处理派送中订单: [%s]", time.Now().Format("2006-01-02 15:04:05"))
	// 一小时前的时间
	duration, _ := time.ParseDuration("-1h")
	nowTime.Add(duration)
	// 查询出仍在派送中的订单
	ordersList := service.GetOrderByStatusAndOrderTime(constant.DeliveryInProgress, model.LocalTime(nowTime))
	if ordersList != nil && len(ordersList) > 0 {
		for _, order := range ordersList {
			order.Status = constant.Completed
			service.UpdateOrder(order)
		}
	}
}
