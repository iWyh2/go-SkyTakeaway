package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/ulule/deepcopier"
	"go-SkyTakeaway/common/constant"
	errs "go-SkyTakeaway/common/errors"
	"go-SkyTakeaway/common/utils"
	"go-SkyTakeaway/global"
	dto "go-SkyTakeaway/model/dto"
	"go-SkyTakeaway/model/entity"
	pageResult "go-SkyTakeaway/model/result/page"
	"go-SkyTakeaway/model/vo"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

// AddCategory 新增分类
func AddCategory(info *dto.CategoryDTO, ctx *gin.Context) {
	// 创建分类数据模型
	var category entity.Category
	// 属性拷贝（使用deepcopier库）
	err := deepcopier.Copy(info).To(&category)
	if err != nil {
		panic(errs.ServerInternalError)
	}
	// 设置分类状态，默认为禁用0
	category.Status = constant.Disable
	// 填充公共字段
	utils.AutoFillEmpID(constant.Create, ctx, &category)
	// 创建数据库表数据
	if err := global.Db.Create(&category).Error; err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			// 分类已存在
			panic(errs.CategoryAlreadyExistsError)
		} else {
			// 数据库异常，创建失败
			panic(errs.DBError)
		}
	}
}

// CategoryPageQuery 分类分页查询
func CategoryPageQuery(pageData *dto.CategoryPageQueryDTO, ctx *gin.Context) *pageResult.Page[entity.Category] {
	// 获取分页参数
	pageIndex := pageData.Page
	pageSize := pageData.PageSize
	// 准备数据容器
	categories := make([]entity.Category, 0)
	// 准备返回数据
	page := &pageResult.Page[entity.Category]{}
	// 指定查询模型，方便后续操作
	query := global.Db.Model(&entity.Category{})
	// 设置模糊查询条件1
	if _, isExist := ctx.GetQuery("name"); isExist {
		query = query.Where("name like ?", "%"+ctx.Query("name")+"%")
	}
	// 设置模糊查询条件2
	if _, isExist := ctx.GetQuery("type"); isExist {
		query = query.Where("type = ?", ctx.Query("type"))
	}
	// 设置按 create_time desc 排序
	if err := query.Order("update_time desc, sort asc").Error; err != nil {
		panic(errs.DBError)
	}
	// 统计数据数量，执行分页查询
	if err := query.Count(&page.Total).
		Limit(pageSize).Offset((pageIndex - 1) * pageSize).
		Find(&categories).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(errs.RecordNotFoundError)
		}
		panic(errs.DBError)
	}
	// 装入查询数据
	page.Records = categories
	// 返回数据
	return page
}

// DeleteCategoryByID 根据id删除分类
func DeleteCategoryByID(id string) {
	// 执行删除
	if err := global.Db.Delete(&entity.Category{}, id).Error; err != nil {
		panic(errs.DBError)
	}
}

// UpdateCategory 修改分类
func UpdateCategory(info *dto.CategoryDTO, ctx *gin.Context) {
	// 创建分类数据模型
	var category entity.Category
	// 属性拷贝（使用deepcopier库）
	err := deepcopier.Copy(info).To(&category)
	if err != nil {
		panic(errs.ServerInternalError)
	}
	// 填充公共字段
	utils.AutoFillEmpID(constant.Update, ctx, &category)
	// 更新信息
	if err := global.Db.Updates(&category).Error; err != nil {
		panic(errs.DBError)
	}
}

// StartOrStopCategory 启用/禁用分类
func StartOrStopCategory(status, id string) {
	// 更新status
	if err := global.Db.Table("category").Where("id = ?", id).Update("status", status).Error; err != nil {
		panic(errs.DBError)
	}
}

// QueryCategoryByType 根据类型查询分类
func QueryCategoryByType(typeData string) *[]entity.Category {
	// 创建数据模型
	var categories []entity.Category
	// 设置查询条件
	query := global.Db.Table("category")
	if typeData != "" {
		query = query.Where("type = ?", typeData)
	}
	// 执行查询
	if err := query.Find(&categories).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(errs.RecordNotFoundError)
		}
		panic(errs.DBError)
	}
	// 返回
	return &categories
}

// UserQueryCategoryByType 根据类型查询分类
func UserQueryCategoryByType(typeData string) *[]vo.CategoryVO {
	// 创建数据模型
	var categories []entity.Category
	categoryVOs := make([]vo.CategoryVO, 0)
	// 设置查询条件
	query := global.Db.Table("category")
	if typeData != "" {
		query = query.Where("type = ?", typeData)
	}
	query = query.Where("status", 1).
		Order("sort, create_time desc")
	// 执行查询
	if err := query.Find(&categories).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(errs.RecordNotFoundError)
		}
		panic(errs.DBError)
	}
	// 处理数据
	for _, category := range categories {
		var categoryVO vo.CategoryVO
		// 属性拷贝（使用deepcopier库）
		err := deepcopier.Copy(&category).To(&categoryVO)
		if err != nil {
			panic(errs.ServerInternalError)
		}
		t, _ := strconv.Atoi(category.Type)
		categoryVO.Type = t
		categoryVOs = append(categoryVOs, categoryVO)
	}
	// 返回
	return &categoryVOs
}
