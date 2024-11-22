package dto

import (
	"go-SkyTakeaway/model/entity"
)

// SetmealDTO 新增 / 修改套餐数据模型
type SetmealDTO struct {
	ID            int                  `json:"id"`
	CategoryID    int                  `json:"categoryId"`
	Name          string               `json:"name"`
	Price         string               `json:"price"`
	Image         string               `json:"image"`
	Description   string               `json:"description"`
	Status        int                  `json:"status"`
	SetmealDishes []entity.SetmealDish `json:"setmealDishes"`
}
