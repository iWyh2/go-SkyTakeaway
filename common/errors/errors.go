package errs

import (
	"errors"
	"go-SkyTakeaway/common/constant"
)

// 自定义错误
var (
	ServerInternalError         = errors.New(constant.ServerInternalError)
	AccountNotFoundError        = errors.New(constant.AccountNotFound)
	DBError                     = errors.New(constant.DbError)
	PasswordError               = errors.New(constant.PasswordError)
	AccountLockedError          = errors.New(constant.AccountLocked)
	ContextNoIDError            = errors.New(constant.ContextNoID)
	MissAuthHeaderError         = errors.New(constant.MissAuthHeader)
	JwtParseError               = errors.New(constant.JwtParseError)
	AuthError                   = errors.New(constant.UserNotLogin)
	UsernameAlreadyExistsError  = errors.New(constant.UsernameAlreadyExists)
	UserAlreadyExistsError      = errors.New(constant.UserAlreadyExists)
	RecordNotFoundError         = errors.New(constant.RecordNotFound)
	CategoryAlreadyExistsError  = errors.New(constant.CategoryAlreadyExists)
	DishAlreadyExistsError      = errors.New(constant.DishAlreadyExists)
	InvalidIDError              = errors.New(constant.InvalidID)
	UploadFileError             = errors.New(constant.UploadFailed)
	NoFileError                 = errors.New(constant.NoFile)
	DishOnSaleError             = errors.New(constant.DishOnSale)
	DishBeRelatedBySetmealError = errors.New(constant.DishBeRelatedBySetmeal)
	SetmealOnSaleError          = errors.New(constant.SetmealOnSale)
	SetmealAlreadyExistsError   = errors.New(constant.SetmealAlreadyExists)
	DishDisableInSetmealError   = errors.New(constant.DishDisableInSetmeal)
	RedisError                  = errors.New(constant.RedisError)
	RedisNotFoundError          = errors.New(constant.RedisNotFound)
	WxLoginFailedError          = errors.New(constant.WxLoginFailed)
	MissUserIdError             = errors.New(constant.MissUserId)
	AddressBookIsNilError       = errors.New(constant.AddressBookIsNull)
	ShoppingCartIsNilError      = errors.New(constant.ShoppingCartIsNull)
	TimerTaskError              = errors.New(constant.TimerTaskError)
	WebSocketUpgradeError       = errors.New(constant.WebSocketUpgradeError)
	WebSocketSendMsgError       = errors.New(constant.WebSocketSendMsgError)
	OrderNotExistError          = errors.New(constant.OrderNotFound)
	OrderStatusError            = errors.New(constant.OrderStatusError)
	RedisConnectError           = errors.New(constant.RedisConnectFail)
)
