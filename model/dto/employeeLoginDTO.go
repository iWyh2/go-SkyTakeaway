package dto

// EmployeeLoginDTO 员工登录输入数据模型
type EmployeeLoginDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
