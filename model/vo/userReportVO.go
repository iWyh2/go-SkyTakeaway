package vo

// UserReportVO 用户统计返回数据模型
type UserReportVO struct {
	DateList      string `json:"dateList"`
	TotalUserList string `json:"totalUserList"`
	NewUserList   string `json:"newUserList"`
}
