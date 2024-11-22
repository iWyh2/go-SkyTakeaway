package service

import (
	"errors"
	errs "go-SkyTakeaway/common/errors"
	"go-SkyTakeaway/global"
	"go-SkyTakeaway/model/entity"
	"gorm.io/gorm"
)

// 向订单明细表插入多条数据
func insertBatchOrderDetail(orderDetailList []entity.OrderDetail) {
	if err := global.Db.Create(&orderDetailList).Error; err != nil {
		panic(errs.DBError)
	}
}

// GetOrderDetailByOrderId 根据订单id查询菜品/套餐明细
func GetOrderDetailByOrderId(orderId string) (orderDetailList []entity.OrderDetail) {
	if err := global.Db.Where("order_id = ?", orderId).
		Find(&orderDetailList).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		panic(errs.DBError)
	}
	return orderDetailList
}
