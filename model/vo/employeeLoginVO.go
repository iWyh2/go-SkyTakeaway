package vo

// EmployeeLoginVO 返回的员工数据模型
type EmployeeLoginVO struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Token    string `json:"token"`
}
