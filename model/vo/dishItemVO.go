package vo

// DishItemVO 用户端返回菜品项数据模型
type DishItemVO struct {
	Name        string `json:"name"`
	Copies      int    `json:"copies"`
	Image       string `json:"image"`
	Description string `json:"description"`
}
