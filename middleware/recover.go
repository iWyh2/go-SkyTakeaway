package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	errs "go-SkyTakeaway/common/errors"
	model "go-SkyTakeaway/model/result"
	"log"
	"net/http"
)

// Recover 全局错误处理中间件
func Recover(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			// 打印错误信息
			log.Printf("panic: %v\n", r)
			// 错误处理
			errorHandle(r.(error), ctx)
			// 终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
			ctx.Abort()
		}
	}()
	// 加载完 defer recover，继续后续接口调用
	ctx.Next()
}

// 根据error类型进行不同处理
func errorHandle(err error, ctx *gin.Context) {
	// 封装通用json返回
	var result model.Result[any]
	if errors.Is(err, errs.AuthError) ||
		errors.Is(err, errs.PasswordError) {
		ctx.JSON(http.StatusUnauthorized, result.Error(errorToString(err)))
	} else if errors.Is(err, errs.UsernameAlreadyExistsError) ||
		errors.Is(err, errs.RecordNotFoundError) ||
		errors.Is(err, errs.InvalidIDError) ||
		errors.Is(err, errs.NoFileError) {
		ctx.JSON(http.StatusBadRequest, result.Error(errorToString(err)))
	} else if errors.Is(err, errs.DishOnSaleError) ||
		errors.Is(err, errs.DishBeRelatedBySetmealError) ||
		errors.Is(err, errs.SetmealOnSaleError) ||
		errors.Is(err, errs.DishDisableInSetmealError) ||
		errors.Is(err, errs.DishAlreadyExistsError) ||
		errors.Is(err, errs.SetmealAlreadyExistsError) ||
		errors.Is(err, errs.CategoryAlreadyExistsError) {
		ctx.JSON(http.StatusOK, result.Error(errorToString(err)))
	} else {
		ctx.JSON(http.StatusInternalServerError, result.Error(errorToString(err)))
	}
}

// recover错误转string
func errorToString(r any) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}
