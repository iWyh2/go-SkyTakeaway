package adminController

import (
	"github.com/gin-gonic/gin"
	errs "go-SkyTakeaway/common/errors"
	dto2 "go-SkyTakeaway/model/dto"
	model "go-SkyTakeaway/model/result"
	pageResult "go-SkyTakeaway/model/result/page"
	"go-SkyTakeaway/model/vo"
	"go-SkyTakeaway/service"
	"log"
	"net/http"
	"strings"
)

// AddSetmeal 新增套餐
func AddSetmeal(ctx *gin.Context) {
	// 创建接收数据模型
	var data dto2.SetmealDTO
	// 模型绑定
	if err := ctx.ShouldBindJSON(&data); err != nil {
		panic(errs.ServerInternalError)
	}
	// 日志打印
	log.Printf("新增套餐: %v", data)
	// 调用service层进行处理
	service.AddSetmealWithDish(&data, ctx)
	// 创建统一返回结果
	var result model.Result[string]
	// 响应
	ctx.JSON(http.StatusOK, result.Success())
}

// SetmealPageQuery 套餐分页查询
func SetmealPageQuery(ctx *gin.Context) {
	// 创建分页查询数据模型
	var data dto2.SetmealPageQueryDTO
	// 模型绑定，Query - `form` tag
	if err := ctx.ShouldBindQuery(&data); err != nil {
		panic(errs.ServerInternalError)
	}
	// 日志打印
	log.Printf("套餐分页查询: %v", data)
	// 调用service层进行处理
	page := service.SetmealPageQuery(&data, ctx)
	// 创建统一返回结果
	var result model.Result[pageResult.Page[vo.SetmealVO]]
	// 响应
	ctx.JSON(http.StatusOK, result.SuccessByData(*page))
}

// DeleteSetmealByIDs 删除套餐
func DeleteSetmealByIDs(ctx *gin.Context) {
	// 获得id数据
	data := ctx.QueryArray("ids")
	ids := strings.Split(data[0], ",")
	if len(ids) == 0 {
		panic(errs.InvalidIDError)
	}
	// 日志打印
	log.Printf("根据id删除套餐: [%v]", ids)
	// 调用service层处理
	service.DeleteSetmealByIDs(ids)
	// 创建统一返回结果
	var result model.Result[string]
	// 响应
	ctx.JSON(http.StatusOK, result.Success())
}

// QuerySetmealByID 根据id查询套餐
func QuerySetmealByID(ctx *gin.Context) {
	// 接收Uri传递的数据
	id := ctx.Param("id")
	// 日志打印
	log.Printf("根据id查询套餐: [%v]", id)
	// 调用service层处理
	setmealVO := service.QuerySetmealByIDWithDish(id)
	// 创建统一返回结果
	var result model.Result[vo.SetmealVO]
	// 响应
	ctx.JSON(http.StatusOK, result.SuccessByData(*setmealVO))
}

// UpdateSetmeal 修改套餐
func UpdateSetmeal(ctx *gin.Context) {
	// 接收更新数据
	var info dto2.SetmealDTO
	if err := ctx.ShouldBindJSON(&info); err != nil {
		panic(errs.ServerInternalError)
	}
	// 日志打印
	log.Printf("修改套餐: %v", info)
	// 调用service层进行处理
	service.UpdateSetmeal(&info, ctx)
	// 创建统一返回结果
	var result model.Result[string]
	// 响应
	ctx.JSON(http.StatusOK, result.Success())
}

// StartOrStopSetmeal 起售/停售套餐
func StartOrStopSetmeal(ctx *gin.Context) {
	// 接收Uri传递的数据
	status := ctx.Param("status")
	// 接收query传递的数据
	id := ctx.Query("id")
	// 日志打印
	log.Printf("起售/禁售套餐: [%v:%v]", id, status)
	// 调用service层处理
	service.StartOrStopSetmeal(id, status)
	// 创建统一返回结果
	var result model.Result[string]
	// 响应
	ctx.JSON(http.StatusOK, result.Success())
}
