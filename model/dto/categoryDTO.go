package dto

// CategoryDTO 新增 / 修改分类数据模型
type CategoryDTO struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
	Name string `json:"name" binding:"required"`
	Sort string `json:"sort" binding:"required"`
}
