package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/ulule/deepcopier"
	"go-SkyTakeaway/common/constant"
	errs "go-SkyTakeaway/common/errors"
	"go-SkyTakeaway/common/utils"
	"go-SkyTakeaway/global"
	"go-SkyTakeaway/model/dto"
	"go-SkyTakeaway/model/entity"
	pageResult "go-SkyTakeaway/model/result/page"
	"gorm.io/gorm"
	"strings"
)

// Login 员工登录service层处理
func Login(input *dto.EmployeeLoginDTO) *entity.Employee {
	// 获得输入的用户名与密码
	username := input.Username
	password := input.Password
	// 创建员工数据模型
	var employee entity.Employee
	// 查询数据库
	if err := global.Db.Where("username = ?", username).First(&employee).Error; err != nil {
		// 数据库未查询到
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(errs.AccountNotFoundError)
		} else {
			// 数据库异常
			panic(errs.DBError)
		}
	}
	// 对比密码
	password = utils.MD5encrypt(password)
	if password != employee.Password {
		// 密码错误
		panic(errs.PasswordError)
	}
	// 检查状态
	if employee.Status == constant.Disable {
		// 账号状态为禁用
		panic(errs.AccountLockedError)
	}
	// 返回
	return &employee
}

// Register 新增员工
func Register(info *dto.EmployeeDTO, ctx *gin.Context) {
	// 创建员工数据模型
	var employee entity.Employee
	// 属性拷贝（使用deepcopier库）
	err := deepcopier.Copy(info).To(&employee)
	if err != nil {
		panic(errs.ServerInternalError)
	}
	// 设置账号状态
	employee.Status = constant.Enable
	// 设置账号密码
	employee.Password = utils.MD5encrypt(constant.DefaultPassword)
	// 填充公共字段
	utils.AutoFillEmpID(constant.Create, ctx, &employee)
	// 创建数据库表数据
	if err := global.Db.Create(&employee).Error; err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			// 用户名已存在
			panic(errs.UsernameAlreadyExistsError)
		} else {
			// 数据库异常，创建失败
			panic(errs.DBError)
		}
	}
}

// EmpPageQuery 员工分页查询
func EmpPageQuery(pageData *dto.EmployeePageQueryDTO, ctx *gin.Context) *pageResult.Page[entity.Employee] {
	// 获取分页参数
	pageIndex := pageData.Page
	pageSize := pageData.PageSize
	// 准备数据容器
	emps := make([]entity.Employee, 0)
	// 准备返回数据
	page := &pageResult.Page[entity.Employee]{}
	// 指定查询模型，方便后续操作
	query := global.Db.Model(&entity.Employee{})
	// 设置模糊查询条件
	if _, isExist := ctx.GetQuery("name"); isExist {
		query = query.Where("name like ?", "%"+ctx.Query("name")+"%")
	}
	// 设置按 create_time desc 排序
	if err := query.Order("create_time desc").Error; err != nil {
		panic(errs.DBError)
	}
	// 统计数据数量，执行分页查询
	if err := query.Count(&page.Total).
		Limit(pageSize).Offset((pageIndex - 1) * pageSize).
		Find(&emps).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(errs.RecordNotFoundError)
		}
		panic(errs.DBError)
	}
	// 装入查询数据
	page.Records = emps
	// 返回数据
	return page
}

// StartOrStopEmp 启用 / 禁用员工账号
func StartOrStopEmp(status, id string) {
	// 更新status
	if err := global.Db.Table("employee").Where("id = ?", id).Update("status", status).Error; err != nil {
		panic(errs.DBError)
	}
}

// EmpQueryByID 根据id查询员工信息
func EmpQueryByID(id int) *entity.Employee {
	// 创建数据模型
	var employee entity.Employee
	// 查询数据库
	if err := global.Db.Where("id = ?", id).First(&employee).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(errs.RecordNotFoundError)
		}
		panic(errs.DBError)
	}
	// 隐藏密码
	employee.Password = "********"
	// 返回
	return &employee
}

// EmpUpdate 编辑员工信息
func EmpUpdate(info *dto.EmployeeDTO, ctx *gin.Context) {
	// 创建员工数据模型
	var employee entity.Employee
	// 属性拷贝（使用deepcopier库）
	err := deepcopier.Copy(info).To(&employee)
	if err != nil {
		panic(errs.ServerInternalError)
	}
	// 填充公共字段
	utils.AutoFillEmpID(constant.Update, ctx, &employee)
	// 更新信息
	if err := global.Db.Updates(&employee).Error; err != nil {
		panic(errs.DBError)
	}
}
