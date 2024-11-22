package dto

import (
	"go-SkyTakeaway/model/entity"
)

// DishDTO 新增 / 修改菜品数据模型
type DishDTO struct {
	ID          int                 `json:"id"`
	Name        string              `json:"name"`
	CategoryID  int                 `json:"categoryId"`
	Price       string              `json:"price"`
	Image       string              `json:"image"`
	Description string              `json:"description"`
	Status      int                 `json:"status"`
	Flavors     []entity.DishFlavor `json:"flavors"`
}
