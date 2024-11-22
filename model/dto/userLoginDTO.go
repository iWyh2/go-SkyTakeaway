package dto

// UserLoginDTO 用户端用户登录传输数据模型
type UserLoginDTO struct {
	// 微信授权码
	Code string `json:"code" binding:"required"`
}
