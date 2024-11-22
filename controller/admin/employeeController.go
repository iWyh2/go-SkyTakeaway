package adminController

import (
	"github.com/gin-gonic/gin"
	"go-SkyTakeaway/common/constant"
	errs "go-SkyTakeaway/common/errors"
	"go-SkyTakeaway/common/utils"
	dto2 "go-SkyTakeaway/model/dto"
	"go-SkyTakeaway/model/entity"
	model "go-SkyTakeaway/model/result"
	pageResult "go-SkyTakeaway/model/result/page"
	"go-SkyTakeaway/model/vo"
	"go-SkyTakeaway/service"
	"log"
	"net/http"
	"strconv"
)

// Logout 员工退出
func Logout(ctx *gin.Context) {
	// 日志打印
	log.Printf("员工退出")
	// 创建统一返回结果
	var result model.Result[string]
	// 响应
	ctx.JSON(http.StatusOK, result.Success())
}

// Login 员工登录
func Login(ctx *gin.Context) {
	// 创建用户输入数据模型
	var input dto2.EmployeeLoginDTO
	// 模型绑定
	if err := ctx.ShouldBindJSON(&input); err != nil {
		panic(errs.ServerInternalError)
	}
	// 日志打印
	log.Printf("员工登录: %v", input)
	// 调用service层逻辑进行处理，最终获得员工数据
	employee := service.Login(&input)
	// 登录成功，生成JWT令牌
	token, err := utils.GenerateJWT(constant.EmpID, strconv.Itoa(employee.ID))
	if err != nil {
		panic(errs.ServerInternalError)
	}
	// 创建响应数据模型
	var data vo.EmployeeLoginVO
	// 携带员工数据
	data.ID = employee.ID
	data.Username = employee.Username
	data.Name = employee.Name
	// 携带token
	data.Token = token
	// 创建统一返回结果
	var result model.Result[vo.EmployeeLoginVO]
	// 响应
	ctx.JSON(http.StatusOK, result.SuccessByData(data))
}

// Register 新增员工
func Register(ctx *gin.Context) {
	// 创建前端提交数据模型
	var data dto2.EmployeeDTO
	// 模型绑定
	if err := ctx.ShouldBindJSON(&data); err != nil {
		panic(errs.ServerInternalError)
	}
	// 日志打印
	log.Printf("新增员工: %v", data)
	// 调用service层进行处理
	service.Register(&data, ctx)
	// 创建统一返回结果
	var result model.Result[string]
	// 响应
	ctx.JSON(http.StatusOK, result.Success())
}

// EmpPageQuery 员工分页查询
func EmpPageQuery(ctx *gin.Context) {
	// 创建分页查询数据模型
	var data dto2.EmployeePageQueryDTO
	// 模型绑定Query - `form` tag
	if err := ctx.ShouldBindQuery(&data); err != nil {
		panic(errs.ServerInternalError)
	}
	// 日志打印
	log.Printf("员工分页查询: %v", data)
	// 调用service层进行处理
	page := service.EmpPageQuery(&data, ctx)
	// 创建统一返回结果
	var result model.Result[pageResult.Page[entity.Employee]]
	// 响应
	ctx.JSON(http.StatusOK, result.SuccessByData(*page))
}

// StartOrStopEmp 启用或禁用员工账号
func StartOrStopEmp(ctx *gin.Context) {
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
	log.Printf("启用或禁用员工账号: [%v:%v]", id, status.Status)
	// 调用service层处理
	service.StartOrStopEmp(status.Status, id)
	// 创建统一返回结果
	var result model.Result[string]
	// 响应
	ctx.JSON(http.StatusOK, result.Success())
}

// EmpQueryByID 根据id查询员工信息
func EmpQueryByID(ctx *gin.Context) {
	// 接收Uri传递的数据
	var id struct {
		ID int `uri:"id"`
	}
	if err := ctx.ShouldBindUri(&id); err != nil {
		panic(errs.ServerInternalError)
	}
	// 日志打印
	log.Printf("根据id查询员工信息: [%v]", id.ID)
	// 调用service层处理
	employee := service.EmpQueryByID(id.ID)
	// 创建统一返回结果
	var result model.Result[entity.Employee]
	// 响应
	ctx.JSON(http.StatusOK, result.SuccessByData(*employee))
}

// EmpUpdate 编辑员工信息
func EmpUpdate(ctx *gin.Context) {
	// 接收更新数据
	var info dto2.EmployeeDTO
	if err := ctx.ShouldBindJSON(&info); err != nil {
		panic(errs.ServerInternalError)
	}
	// 日志打印
	log.Printf("编辑员工信息: %v", info)
	// 调用service层进行处理
	service.EmpUpdate(&info, ctx)
	// 创建统一返回结果
	var result model.Result[string]
	// 响应
	ctx.JSON(http.StatusOK, result.Success())
}
