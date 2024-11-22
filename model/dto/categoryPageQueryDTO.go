package dto

// CategoryPageQueryDTO 分类分页查询数据模型
type CategoryPageQueryDTO struct {
	Type     string `form:"type"`
	Name     string `form:"name"`
	Page     int    `form:"page" binding:"required"`
	PageSize int    `form:"pageSize" binding:"required"`
}
