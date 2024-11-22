package middleware

import (
	"github.com/gin-gonic/gin"
	"go-SkyTakeaway/common/constant"
	errs "go-SkyTakeaway/common/errors"
	"go-SkyTakeaway/common/utils"
	"go-SkyTakeaway/config"
)

// JwtAdmin 管理端校验Jwt令牌中间件
func JwtAdmin(ctx *gin.Context) {
	// 获得token
	token := ctx.GetHeader(config.ServerConfig.Jwt.AdminTokenName)
	if token == "" {
		// 未登录 / 未认证
		panic(errs.AuthError)
	}
	// jwt校验
	empId, err := utils.ParseJWT(token, constant.EmpID)
	if err != nil {
		// token解析错误
		panic(errs.JwtParseError)
	}
	// 存储当前已登录员工id
	ctx.Set(constant.EmpID, empId)
	// 继续调用
	ctx.Next()
}

// JwtUser 用户端校验Jwt令牌中间件
func JwtUser(ctx *gin.Context) {
	// 获得token
	token := ctx.GetHeader(config.ServerConfig.Jwt.UserTokenName)
	if token == "" {
		// 缺少Authorization标头，未登录 / 未认证
		panic(errs.AuthError)
	}
	// jwt校验
	userId, err := utils.ParseJWT(token, constant.UserID)
	if err != nil {
		// token解析错误
		panic(errs.JwtParseError)
	}
	// 存储当前已登录用户id
	ctx.Set(constant.UserID, userId)
	// 继续调用
	ctx.Next()
}
