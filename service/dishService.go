package service

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/ulule/deepcopier"
	"go-SkyTakeaway/common/constant"
	errs "go-SkyTakeaway/common/errors"
	"go-SkyTakeaway/common/utils"
	"go-SkyTakeaway/global"
	"go-SkyTakeaway/model/dto"
	"go-SkyTakeaway/model/entity"
	pageResult "go-SkyTakeaway/model/result/page"
	"go-SkyTakeaway/model/vo"
	"gorm.io/gorm"
	"log"
	"strconv"
	"strings"
)

// AddDishWithFlavor 新增菜品
func AddDishWithFlavor(data *dto.DishDTO, ctx *gin.Context) {
	// 创建分类数据模型
	var dish entity.Dish
	// 属性拷贝（使用deepcopier库）
	err := deepcopier.Copy(data).To(&dish)
	if err != nil {
		panic(errs.ServerInternalError)
	}
	// 设置分类状态，默认为禁用0
	dish.Status = constant.Disable
	// 填充公共字段
	utils.AutoFillEmpID(constant.Create, ctx, &dish)
	// 创建数据库表数据，菜品
	if err := global.Db.Create(&dish).Error; err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			// 分类已存在
			panic(errs.DishAlreadyExistsError)
		} else {
			// 数据库异常，创建失败
			panic(errs.DBError)
		}
	}
	// 获得口味数据
	flavors := data.Flavors
	if flavors != nil {
		// 插入菜品口味数据
		addFlavorByDishID(flavors, dish.ID)
	}
	// 清理缓存数据
	utils.CleanCache(constant.DishCacheKey + strconv.Itoa(data.CategoryID))
}

// DishPageQuery 菜品分页查询
func DishPageQuery(pageData *dto.DishPageQueryDTO, ctx *gin.Context) *pageResult.Page[vo.DishVO] {
	// 获取分页参数
	pageIndex := pageData.Page
	pageSize := pageData.PageSize
	// 准备数据容器
	dishVOs := make([]vo.DishVO, 0)
	// 准备返回数据
	page := &pageResult.Page[vo.DishVO]{}
	// 指定查询模型，联合查询
	query := global.Db.Model(&entity.Dish{}).
		Select("dish.*, category.name as category_name").
		Joins("left join category on category.id = dish.category_id")
	// 设置模糊查询条件1
	if _, isExist := ctx.GetQuery("name"); isExist {
		query = query.Where("dish.name like ?", "%"+ctx.Query("name")+"%")
	}
	// 设置模糊查询条件2
	if _, isExist := ctx.GetQuery("categoryId"); isExist {
		query = query.Where("dish.category_id = ?", ctx.Query("categoryId"))
	}
	// 设置模糊查询条件3
	if _, isExist := ctx.GetQuery("status"); isExist {
		if ctx.Query("status") != "" {
			query = query.Where("dish.status = ?", ctx.Query("status"))
		}
	}
	// 设置按 create_time desc 排序
	if err := query.Order("dish.create_time desc").Error; err != nil {
		panic(errs.DBError)
	}
	// 统计数据数量，执行分页查询
	if err := query.Count(&page.Total).
		Limit(pageSize).Offset((pageIndex - 1) * pageSize).
		Scan(&dishVOs).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(errs.RecordNotFoundError)
		}
		panic(errs.DBError)
	}
	// 处理查询数据
	for i := 0; i < len(dishVOs); i++ {
		dishVOs[i].Flavors = make([]entity.DishFlavor, 0)
	}
	// 装入查询数据
	page.Records = dishVOs
	// 返回数据
	return page
}

// DeleteDishByIDs 删除菜品
func DeleteDishByIDs(ids []string) {
	// 判断当前菜品是否能够删除
	// 是否存在起售中的菜品
	for _, id := range ids {
		dish := queryDishByID(id)
		if dish.Status != constant.Disable {
			panic(errs.DishOnSaleError)
		}
	}
	// 是否被套餐关联
	setmealIDs := getSetmealIDsByDishIDs(ids)
	if len(setmealIDs) > 0 {
		panic(errs.DishBeRelatedBySetmealError)
	}
	// 执行删除菜品
	if err := global.Db.Where("id in (?)", ids).Delete(&entity.Dish{}).Error; err != nil {
		panic(errs.DBError)
	}
	// 执行删除菜品关联的口味数据
	deleteFlavorByDishID(ids)
	// 将所有的菜品缓存数据清理掉
	utils.CleanCache(constant.DishCacheKey + "*")
}

// 根据id查询菜品
func queryDishByID(id string) *entity.Dish {
	// 准备数据
	var dish entity.Dish
	// 执行查询
	if err := global.Db.First(&dish, id).Error; err != nil {
		panic(errs.DBError)
	}
	return &dish
}

// QueryDishByIDWithFlavor 根据id查询菜品和口味数据
func QueryDishByIDWithFlavor(id string) *vo.DishVO {
	// 根据id查询菜品数据
	dish := queryDishByID(id)
	// 根据菜品id查询口味数据
	dishFlavors := getFlavorByDishID(id)
	// 将查询到的数据封装到VO
	var dishVO vo.DishVO
	err := deepcopier.Copy(dish).To(&dishVO)
	if err != nil {
		panic(errs.ServerInternalError)
	}
	dishVO.Flavors = dishFlavors
	dishVO.Price, _ = strconv.ParseFloat(dish.Price, 64)
	return &dishVO
}

// UpdateDish 修改菜品
func UpdateDish(info *dto.DishDTO, ctx *gin.Context) {
	// 创建菜品数据模型
	var dish entity.Dish
	// 属性拷贝（使用deepcopier库）
	err := deepcopier.Copy(info).To(&dish)
	if err != nil {
		panic(errs.ServerInternalError)
	}
	// 填充公共字段
	utils.AutoFillEmpID(constant.Update, ctx, &dish)
	// 修改菜品基本信息
	if err := global.Db.Updates(&dish).Error; err != nil {
		panic(errs.DBError)
	}
	// 删除原有的口味数据
	deleteFlavorByDishID([]string{strconv.Itoa(info.ID)})
	// 重新插入口味数据
	flavors := info.Flavors
	if flavors != nil {
		addFlavorByDishID(flavors, info.ID)
	}
	// 将所有的菜品缓存数据清理掉
	utils.CleanCache(constant.DishCacheKey + "*")
}

// StartOrStopDish 起售/禁售菜品
func StartOrStopDish(id, status string) {
	// 更新菜品status
	if err := global.Db.
		Table("dish").
		Where("id = ?", id).
		Update("status", status).Error; err != nil {
		panic(errs.DBError)
	}
	// 如果是禁售操作，还需要将包含当前菜品的套餐也停售
	if status == strconv.Itoa(constant.Disable) {
		ids := []string{id}
		// 查询关联的套餐
		setmealIDs := getSetmealIDsByDishIDs(ids)
		if len(setmealIDs) > 0 {
			for _, id := range setmealIDs {
				// 更新套餐status
				if err := global.Db.
					Table("setmeal").
					Where("id = ?", id).
					Update("status", status).Error; err != nil {
					panic(errs.DBError)
				}
			}
		}
	}
	// 将所有的菜品缓存数据清理掉
	utils.CleanCache(constant.DishCacheKey + "*")
}

// QueryDishByCategoryID 根据分类id查询菜品
func QueryDishByCategoryID(categoryId string) *[]entity.Dish {
	// 创建数据模型
	var dishes []entity.Dish
	// 设置查询条件
	query := global.Db.Table("dish")
	if categoryId != "" {
		query = query.Where("category_id = ?", categoryId)
	}
	query = query.Where("status = ?", constant.Enable).Order("create_time desc")
	// 执行查询
	if err := query.Find(&dishes).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(errs.RecordNotFoundError)
		}
		panic(errs.DBError)
	}
	// 返回
	return &dishes
}

// 根据套餐id获得菜品数据
func queryDishBySetmealID(setmealID string) []entity.Dish {
	// 准备数据容器
	dishes := make([]entity.Dish, 0)
	// 指定查询模型，联合查询
	query := global.Db.Model(&entity.Dish{}).
		Select("dish.*").
		Joins("left join setmeal_dish on dish.id = setmeal_dish.dish_id")
	// 查询数据
	query.Where("setmeal_dish.setmeal_id = ?", setmealID).Scan(&dishes)
	return dishes
}

// UserQueryDishByCategoryID 根据分类id查询菜品
func UserQueryDishByCategoryID(categoryId string) *[]vo.DishVO {
	// 创建数据容器
	dishVOs := make([]vo.DishVO, 0)
	// 首先查询缓存是否存在数据
	cacheData, err := global.RedisClient.Get(constant.DishCacheKey + categoryId).Result()
	if err == nil {
		// 有缓存，直接返回缓存数据
		// 日志打印
		log.Printf("查询Redis菜品缓存数据: [%v]", constant.DishCacheKey+categoryId)
		// 反序列化缓存
		if err := json.Unmarshal([]byte(cacheData), &dishVOs); err != nil {
			panic(errs.ServerInternalError)
		}
		// 返回
		return &dishVOs
	} else if !errors.Is(err, redis.Nil) {
		panic(errs.RedisError)
	}
	// 没有缓存数据再查询数据库
	// 创建数据模型
	var dishes []entity.Dish
	// 设置查询条件
	query := global.Db.Table("dish").
		Where("category_id = ?", categoryId).
		Where("status = ?", constant.Enable).
		Order("create_time desc")
	// 执行查询
	if err := query.Find(&dishes).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(errs.RecordNotFoundError)
		}
		panic(errs.DBError)
	}
	// 处理数据
	for _, dish := range dishes {
		var dishVO vo.DishVO
		// 属性拷贝（使用deepcopier库）
		err := deepcopier.Copy(&dish).To(&dishVO)
		if err != nil {
			panic(errs.ServerInternalError)
		}
		// 根据菜品id查询对应的口味
		flavors := getFlavorByDishID(strconv.Itoa(dish.ID))
		dishVO.Flavors = flavors
		// 处理价格
		price, _ := strconv.ParseFloat(dish.Price, 64)
		dishVO.Price = price
		dishVOs = append(dishVOs, dishVO)
	}
	// 序列化为json
	dishVOsJSON, err := json.Marshal(dishVOs)
	if err != nil {
		panic(errs.ServerInternalError)
	}
	// 设置缓存，过期时间为0表示不会过期
	if err := global.RedisClient.Set(constant.DishCacheKey+categoryId, dishVOsJSON, 0).Err(); err != nil {
		panic(errs.RedisError)
	}
	// 返回
	return &dishVOs
}

// GetDishCount 获取某状态菜品数量
func GetDishCount(status int) int {
	var dishCount int64
	if err := global.Db.Table("dish").
		Where("status = ?", status).
		Count(&dishCount).Error; err != nil {
		panic(errs.DBError)
	}
	return int(dishCount)
}
