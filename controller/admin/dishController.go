package adminController

import (
	"github.com/gin-gonic/gin"
	errs "go-SkyTakeaway/common/errors"
	"go-SkyTakeaway/model/dto"
	"go-SkyTakeaway/model/entity"
	model "go-SkyTakeaway/model/result"
	pageResult "go-SkyTakeaway/model/result/page"
	"go-SkyTakeaway/model/vo"
	"go-SkyTakeaway/service"
	"log"
	"net/http"
	"strings"
)

// AddDishWithFlavor 新增菜品
func AddDishWithFlavor(ctx *gin.Context) {
	// 创建接收数据模型
	var data dto.DishDTO
	// 模型绑定
	if err := ctx.ShouldBindJSON(&data); err != nil {
		panic(errs.ServerInternalError)
	}
	// 日志打印
	log.Printf("新增菜品: %v", data)
	// 调用service层进行处理
	service.AddDishWithFlavor(&data, ctx)
	// 创建统一返回结果
	var result model.Result[string]
	// 响应
	ctx.JSON(http.StatusOK, result.Success())
}

// DishPageQuery 菜品分页查询
func DishPageQuery(ctx *gin.Context) {
	// 创建分页查询数据模型
	var data dto.DishPageQueryDTO
	// 模型绑定，Query - `form` tag
	if err := ctx.ShouldBindQuery(&data); err != nil {
		panic(errs.ServerInternalError)
	}
	// 日志打印
	log.Printf("菜品分页查询: %v", data)
	// 调用service层进行处理
	page := service.DishPageQuery(&data, ctx)
	// 创建统一返回结果
	var result model.Result[pageResult.Page[vo.DishVO]]
	// 响应
	ctx.JSON(http.StatusOK, result.SuccessByData(*page))
}

// DeleteDishByIDs 删除菜品
func DeleteDishByIDs(ctx *gin.Context) {
	// 获得id数据
	data := ctx.QueryArray("ids")
	ids := strings.Split(data[0], ",")
	if len(ids) == 0 {
		panic(errs.InvalidIDError)
	}
	// 日志打印
	log.Printf("根据id删除菜品: [%v]", ids)
	// 调用service层处理
	service.DeleteDishByIDs(ids)
	// 创建统一返回结果
	var result model.Result[string]
	// 响应
	ctx.JSON(http.StatusOK, result.Success())
}

// QueryDishByID 根据id查询菜品
func QueryDishByID(ctx *gin.Context) {
	// 接收Uri传递的数据
	id := ctx.Param("id")
	// 日志打印
	log.Printf("根据id查询菜品: [%v]", id)
	// 调用service层处理
	dishVO := service.QueryDishByIDWithFlavor(id)
	// 创建统一返回结果
	var result model.Result[vo.DishVO]
	// 响应
	ctx.JSON(http.StatusOK, result.SuccessByData(*dishVO))
}

// UpdateDish 修改菜品
func UpdateDish(ctx *gin.Context) {
	// 接收更新数据
	var info dto.DishDTO
	if err := ctx.ShouldBindJSON(&info); err != nil {
		panic(errs.ServerInternalError)
	}
	// 日志打印
	log.Printf("修改菜品: %v", info)
	// 调用service层进行处理
	service.UpdateDish(&info, ctx)
	// 创建统一返回结果
	var result model.Result[string]
	// 响应
	ctx.JSON(http.StatusOK, result.Success())
}

// StartOrStopDish 起售/禁售菜品
func StartOrStopDish(ctx *gin.Context) {
	// 接收Uri传递的数据
	status := ctx.Param("status")
	// 接收query传递的数据
	id := ctx.Query("id")
	// 日志打印
	log.Printf("起售/禁售菜品: [%v:%v]", id, status)
	// 调用service层处理
	service.StartOrStopDish(id, status)
	// 创建统一返回结果
	var result model.Result[string]
	// 响应
	ctx.JSON(http.StatusOK, result.Success())
}

// QueryDishByCategoryID 根据分类id查询菜品
func QueryDishByCategoryID(ctx *gin.Context) {
	// 接收query传递的数据
	categoryId := ctx.Query("categoryId")
	// 日志打印
	log.Printf("根据分类id查询菜品: [%v]", categoryId)
	// 调用service层处理
	dishes := service.QueryDishByCategoryID(categoryId)
	// 创建统一返回结果
	var result model.Result[[]entity.Dish]
	// 响应
	ctx.JSON(http.StatusOK, result.SuccessByData(*dishes))
}
