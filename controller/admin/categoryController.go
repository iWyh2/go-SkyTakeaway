package adminController

import (
	"github.com/gin-gonic/gin"
	errs "go-SkyTakeaway/common/errors"
	dto2 "go-SkyTakeaway/model/dto"
	"go-SkyTakeaway/model/entity"
	model "go-SkyTakeaway/model/result"
	pageResult "go-SkyTakeaway/model/result/page"
	"go-SkyTakeaway/service"
	"log"
	"net/http"
)

// AddCategory 新增分类
func AddCategory(ctx *gin.Context) {
	// 创建接收数据模型
	var data dto2.CategoryDTO
	// 模型绑定
	if err := ctx.ShouldBindJSON(&data); err != nil {
		panic(errs.ServerInternalError)
	}
	// 日志打印
	log.Printf("新增分类: %v", data)
	// 调用service层进行处理
	service.AddCategory(&data, ctx)
	// 创建统一返回结果
	var result model.Result[string]
	// 响应
	ctx.JSON(http.StatusOK, result.Success())
}

// CategoryPageQuery 分类分页查询
func CategoryPageQuery(ctx *gin.Context) {
	// 创建分页查询数据模型
	var data dto2.CategoryPageQueryDTO
	// 模型绑定，Query - `form` tag
	if err := ctx.ShouldBindQuery(&data); err != nil {
		panic(errs.ServerInternalError)
	}
	// 日志打印
	log.Printf("分类分页查询: %v", data)
	// 调用service层进行处理
	page := service.CategoryPageQuery(&data, ctx)
	// 创建统一返回结果
	var result model.Result[pageResult.Page[entity.Category]]
	// 响应
	ctx.JSON(http.StatusOK, result.SuccessByData(*page))
}

// DeleteCategoryByID 根据id删除分类
func DeleteCategoryByID(ctx *gin.Context) {
	// 获得id数据
	id := ctx.Query("id")
	if id == "" {
		panic(errs.InvalidIDError)
	}
	// 日志打印
	log.Printf("根据id删除分类: [%v]", id)
	// 调用service层处理
	service.DeleteCategoryByID(id)
	// 创建统一返回结果
	var result model.Result[string]
	// 响应
	ctx.JSON(http.StatusOK, result.Success())
}

// UpdateCategory 修改分类
func UpdateCategory(ctx *gin.Context) {
	// 接收更新数据
	var info dto2.CategoryDTO
	if err := ctx.ShouldBindJSON(&info); err != nil {
		panic(errs.ServerInternalError)
	}
	// 日志打印
	log.Printf("修改分类: %v", info)
	// 调用service层进行处理
	service.UpdateCategory(&info, ctx)
	// 创建统一返回结果
	var result model.Result[string]
	// 响应
	ctx.JSON(http.StatusOK, result.Success())
}

// StartOrStopCategory 启用 / 禁用分类
func StartOrStopCategory(ctx *gin.Context) {
	// 接收Uri传递的数据
	var status struct {
		Status string `uri:"status"`
	}
	if err := ctx.ShouldBindUri(&status); err != nil {
		panic(errs.ServerInternalError)
	}
	// 接收query传递的数据
	id := ctx.Query("id")
	// 日志打印
	log.Printf("启用或禁用分类: [%v:%v]", id, status.Status)
	// 调用service层处理
	service.StartOrStopCategory(status.Status, id)
	// 创建统一返回结果
	var result model.Result[string]
	// 响应
	ctx.JSON(http.StatusOK, result.Success())
}

// QueryCategoryByType 根据类型查询分类
func QueryCategoryByType(ctx *gin.Context) {
	// 接收query传递的数据
	typeData := ctx.Query("type")
	// 日志打印
	log.Printf("根据类型查询分类: [%v]", typeData)
	// 调用service层处理
	categories := service.QueryCategoryByType(typeData)
	// 创建统一返回结果
	var result model.Result[[]entity.Category]
	// 响应
	ctx.JSON(http.StatusOK, result.SuccessByData(*categories))
}
