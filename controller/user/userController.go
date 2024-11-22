package userController

import (
	"github.com/gin-gonic/gin"
	"go-SkyTakeaway/common/constant"
	errs "go-SkyTakeaway/common/errors"
	"go-SkyTakeaway/common/utils"
	"go-SkyTakeaway/model/dto"
	model "go-SkyTakeaway/model/result"
	"go-SkyTakeaway/model/vo"
	"go-SkyTakeaway/service"
	"log"
	"net/http"
	"strconv"
)

// UserLogin 用户微信登录
func UserLogin(ctx *gin.Context) {
	// 创建数据模型
	var input dto.UserLoginDTO
	// 模型绑定
	if err := ctx.ShouldBindJSON(&input); err != nil {
		panic(errs.ServerInternalError)
	}
	// 日志打印
	log.Printf("用户微信登录: %v", input)
	// 调用service层逻辑进行处理，最终获得用户数据
	user := service.WxLogin(&input)
	// 登录成功，生成JWT令牌
	token, err := utils.GenerateJWT(constant.UserID, strconv.Itoa(user.ID))
	if err != nil {
		panic(errs.ServerInternalError)
	}
	// 创建响应数据模型
	var data vo.UserLoginVO
	// 携带数据
	data.ID = user.ID
	data.OpenID = user.OpenID
	// 携带token
	data.Token = token
	// 创建统一返回结果
	var result model.Result[vo.UserLoginVO]
	// 响应
	ctx.JSON(http.StatusOK, result.SuccessByData(data))
}
