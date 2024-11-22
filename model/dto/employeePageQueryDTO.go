package dto

// EmployeePageQueryDTO 员工分页查询请求参数模型
type EmployeePageQueryDTO struct {
	Name     string `form:"name"`
	Page     int    `form:"page" binding:"required"`
	PageSize int    `form:"pageSize" binding:"required"`
}
