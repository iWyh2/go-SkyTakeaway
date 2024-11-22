package dto

// OrderPageQueryDTO 订单分页查询数据模型
type OrderPageQueryDTO struct {
	Page      int    `form:"page" binding:"required"`
	PageSize  int    `form:"pageSize" binding:"required"`
	UserId    int    `form:"userId"`
	Number    string `form:"number"`
	Phone     string `form:"phone"`
	Status    int    `form:"status"`
	BeginTime string `form:"beginTime"`
	EndTime   string `form:"endTime"`
}
