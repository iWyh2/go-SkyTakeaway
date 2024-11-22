package vo

// UserLoginVO 返回的用户数据模型
type UserLoginVO struct {
	ID     int    `json:"id"`
	OpenID string `json:"openid"`
	Token  string `json:"token"`
}
