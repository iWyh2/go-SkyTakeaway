package dto

// SetmealPageQueryDTO 套餐分页查询数据模型
type SetmealPageQueryDTO struct {
	Page       int    `form:"page" binding:"required"`
	PageSize   int    `form:"pageSize" binding:"required"`
	Name       string `form:"name"`
	CategoryID string `form:"categoryId"`
	Status     string `form:"status"`
}
