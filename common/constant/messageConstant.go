package constant

// 消息常量
const (
	ServerInternalError          = "服务器异常"
	PasswordError                = "密码错误"
	AccountNotFound              = "账号不存在"
	AccountLocked                = "账号被锁定"
	ContextNoID                  = "没有已登录员工ID"
	MissAuthHeader               = "缺少token标头"
	MissUserId                   = "缺少User ID"
	JwtParseError                = "jwt解析token出现异常"
	UnknownError                 = "未知错误"
	UserNotLogin                 = "用户未登录"
	UsernameAlreadyExists        = "用户名已存在"
	UserAlreadyExists            = "用户名已存在"
	RecordNotFound               = "未找到记录"
	CategoryAlreadyExists        = "分类已存在"
	SetmealAlreadyExists         = "套餐已存在"
	DishAlreadyExists            = "菜品已存在"
	InvalidID                    = "无效ID"
	NoFile                       = "没有文件"
	DishDisableInSetmeal         = "套餐内包含未启售菜品,无法启售"
	CategoryBeRelatedBySetmeal   = "当前分类关联了套餐,不能删除"
	CategoryBeRelatedByDish      = "当前分类关联了菜品,不能删除"
	ShoppingCartIsNull           = "购物车数据为空，不能下单"
	AddressBookIsNull            = "用户地址为空，不能下单"
	WxLoginFailed                = "微信登录失败"
	UploadFailed                 = "文件上传失败"
	SetmealEnableFailed          = "套餐内包含未启售菜品，无法启售"
	PasswordEditFailed           = "密码修改失败"
	DishOnSale                   = "起售中的菜品不能删除"
	SetmealOnSale                = "起售中的套餐不能删除"
	DishBeRelatedBySetmeal       = "当前菜品关联了套餐,不能删除"
	OrderStatusError             = "订单状态错误"
	OrderNotFound                = "订单不存在"
	AlreadyExists                = "已存在"
	NotSelectDeleteObject        = "未选择删除对象"
	RedisNotFound                = "服务器查询缓存异常"
	RedisError                   = "服务器缓存异常"
	DbError                      = "数据库异常"
	TimerTaskError               = "定时器任务执行异常"
	WebSocketUpgradeError        = "webSocket upgrade 错误"
	WebSocketListenAndServeError = "webSocket listen and serve 错误"
	WebSocketSendMsgError        = "webSocket send message 错误"
	RedisConnectFail             = "redis连接失败"
)