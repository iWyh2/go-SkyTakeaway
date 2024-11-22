package dto

// EmployeeDTO 新增 / 修改员工数据模型
type EmployeeDTO struct {
	ID       int    `json:"id"`
	Username string `json:"username" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Sex      string `json:"sex" binding:"required"`
	IdNumber string `json:"idNumber" binding:"required"`
}
