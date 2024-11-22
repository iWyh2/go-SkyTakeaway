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

// AddSetmealWithDish 新增套餐
func AddSetmealWithDish(data *dto.SetmealDTO, ctx *gin.Context) {
	// 创建分类数据模型
	var setmeal entity.Setmeal
	// 属性拷贝（使用deepcopier库）
	err := deepcopier.Copy(data).To(&setmeal)
	if err != nil {
		panic(errs.ServerInternalError)
	}
	// 设置分类状态，默认为禁用0
	setmeal.Status = constant.Disable
	// 填充公共字段
	utils.AutoFillEmpID(constant.Create, ctx, &setmeal)
	// 创建数据库表数据，套餐
	if err := global.Db.Create(&setmeal).Error; err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			// 套餐已存在
			panic(errs.SetmealAlreadyExistsError)
		} else {
			// 数据库异常，创建失败
			panic(errs.DBError)
		}
	}
	// 获得套餐菜品关系数据
	setmealDishes := data.SetmealDishes
	// 保存套餐和菜品的关联关系
	insertSetmealDishes(setmealDishes, setmeal.ID)
	// 将所有的套餐缓存数据清理掉
	utils.CleanCache(constant.SetmealCacheKey + "*")
}

// SetmealPageQuery 套餐分页查询
func SetmealPageQuery(pageData *dto.SetmealPageQueryDTO, ctx *gin.Context) *pageResult.Page[vo.SetmealVO] {
	// 获取分页参数
	pageIndex := pageData.Page
	pageSize := pageData.PageSize
	// 准备数据容器
	setmealVOs := make([]vo.SetmealVO, 0)
	// 准备返回数据
	page := &pageResult.Page[vo.SetmealVO]{}
	// 指定查询模型，联合查询
	query := global.Db.Model(&entity.Setmeal{}).
		Select("setmeal.*, category.name as category_name").
		Joins("left join category on category.id = setmeal.category_id")
	// 设置模糊查询条件1
	if _, isExist := ctx.GetQuery("name"); isExist {
		query = query.Where("setmeal.name like ?", "%"+ctx.Query("name")+"%")
	}
	// 设置模糊查询条件2
	if _, isExist := ctx.GetQuery("categoryId"); isExist {
		query = query.Where("setmeal.category_id = ?", ctx.Query("categoryId"))
	}
	// 设置模糊查询条件3
	if _, isExist := ctx.GetQuery("status"); isExist {
		if ctx.Query("status") != "" {
			query = query.Where("setmeal.status = ?", ctx.Query("status"))
		}
	}
	// 设置按 create_time desc 排序
	if err := query.Order("setmeal.create_time desc").Error; err != nil {
		panic(errs.DBError)
	}
	// 统计数据数量，执行分页查询
	if err := query.Count(&page.Total).
		Limit(pageSize).Offset((pageIndex - 1) * pageSize).
		Scan(&setmealVOs).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(errs.RecordNotFoundError)
		}
		panic(errs.DBError)
	}
	// 处理查询数据
	for i := 0; i < len(setmealVOs); i++ {
		setmealVOs[i].SetmealDishes = make([]entity.SetmealDish, 0)
	}
	// 装入查询数据
	page.Records = setmealVOs
	// 返回数据
	return page
}

// DeleteSetmealByIDs 删除套餐
func DeleteSetmealByIDs(ids []string) {
	// 判断当前套餐是否能够删除
	// 是否存在起售中的套餐
	for _, id := range ids {
		setmeal := querySetmealByID(id)
		if setmeal.Status != constant.Disable {
			panic(errs.SetmealOnSaleError)
		}
	}
	// 删除套餐数据
	if err := global.Db.Where("id in (?)", ids).Delete(&entity.Setmeal{}).Error; err != nil {
		panic(errs.DBError)
	}
	// 删除套餐菜品关系表中的数据
	deleteSetmealDishBySetmealIDs(ids)
	// 将所有的套餐缓存数据清理掉
	utils.CleanCache(constant.SetmealCacheKey + "*")
}

// 根据id查询套餐
func querySetmealByID(id string) *entity.Setmeal {
	// 准备数据
	var setmeal entity.Setmeal
	// 执行查询
	if err := global.Db.First(&setmeal, id).Error; err != nil {
		panic(errs.DBError)
	}
	return &setmeal
}

// QuerySetmealByIDWithDish 根据id查询套餐&菜品
func QuerySetmealByIDWithDish(id string) *vo.SetmealVO {
	// 根据id查询套餐数据
	setmeal := querySetmealByID(id)
	// 根据id查询套餐菜品关系数据
	setmealDishes := querySetmealDishesBySetmealID(id)
	// 将查询到的数据封装到VO
	var setmealVO vo.SetmealVO
	err := deepcopier.Copy(setmeal).To(&setmealVO)
	if err != nil {
		panic(errs.ServerInternalError)
	}
	setmealVO.SetmealDishes = *setmealDishes
	setmealVO.Price, _ = strconv.ParseFloat(setmeal.Price, 64)
	return &setmealVO
}

// UpdateSetmeal 修改套餐
func UpdateSetmeal(info *dto.SetmealDTO, ctx *gin.Context) {
	// 创建套餐数据模型
	var setmeal entity.Setmeal
	// 属性拷贝（使用deepcopier库）
	err := deepcopier.Copy(info).To(&setmeal)
	if err != nil {
		panic(errs.ServerInternalError)
	}
	// 填充公共字段
	utils.AutoFillEmpID(constant.Update, ctx, &setmeal)
	// 修改套餐基本信息
	if err := global.Db.Updates(&setmeal).Error; err != nil {
		panic(errs.DBError)
	}
	// 删除原有的套餐菜品关联数据
	deleteSetmealDishBySetmealIDs([]string{strconv.Itoa(info.ID)})
	// 重新插入关联数据
	setmealDishes := info.SetmealDishes
	insertSetmealDishes(setmealDishes, info.ID)
	// 将所有的套餐缓存数据清理掉
	utils.CleanCache(constant.SetmealCacheKey + "*")
}

// StartOrStopSetmeal 起售/停售套餐
func StartOrStopSetmeal(id, status string) {
	// 起售套餐时，判断套餐内是否有停售菜品，有停售菜品提示"套餐内包含未启售菜品，无法启售"
	if status == strconv.Itoa(constant.Enable) {
		// 查询关联的菜品
		dishes := queryDishBySetmealID(id)
		if len(dishes) > 0 {
			for _, dish := range dishes {
				// 有停售菜品
				if dish.Status == constant.Disable {
					panic(errs.DishDisableInSetmealError)
				}
			}
		}
	}
	// 更新菜品status
	if err := global.Db.
		Table("setmeal").
		Where("id = ?", id).
		Update("status", status).Error; err != nil {
		panic(errs.DBError)
	}
	// 将所有的套餐缓存数据清理掉
	utils.CleanCache(constant.SetmealCacheKey + "*")
}

// UserQuerySetmealByCategoryId 根据分类id查询套餐
func UserQuerySetmealByCategoryId(categoryId string) any {
	// 准备数据容器
	setmealVOs := make([]vo.SetmealVO, 0)
	// 首先查询缓存是否存在数据
	cacheData, err := global.RedisClient.Get(constant.SetmealCacheKey + categoryId).Result()
	if err == nil {
		// 有缓存，直接返回缓存数据
		// 日志打印
		log.Printf("查询Redis套餐缓存数据: [%v]", constant.SetmealCacheKey+categoryId)
		// 反序列化缓存
		if err := json.Unmarshal([]byte(cacheData), &setmealVOs); err != nil {
			panic(errs.ServerInternalError)
		}
		// 返回
		return &setmealVOs
	} else if !errors.Is(err, redis.Nil) {
		panic(errs.RedisError)
	}
	// 没有缓存数据再查询数据库
	// 设置查询条件
	query := global.Db.Table("setmeal").
		Where("category_id = ?", categoryId).
		Where("status = ?", constant.Enable)
	// 执行查询套餐信息
	if err := query.Find(&setmealVOs).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(errs.RecordNotFoundError)
		}
		panic(errs.DBError)
	}
	// 序列化为json
	setmealVOsJSON, err := json.Marshal(setmealVOs)
	if err != nil {
		panic(errs.ServerInternalError)
	}
	// 设置缓存，过期时间为0表示不会过期
	if err := global.RedisClient.Set(constant.SetmealCacheKey+categoryId, setmealVOsJSON, 0).Err(); err != nil {
		panic(errs.RedisError)
	}
	return setmealVOs
}

// UserQueryDishBySetmealID 根据套餐id查询具体菜品信息
func UserQueryDishBySetmealID(setmealId string) *[]vo.DishItemVO {
	// 准备数据容器
	dishItemVOs := make([]vo.DishItemVO, 0)
	// 指定查询模型，联合查询
	if err := global.Db.Table("setmeal_dish").
		Select("setmeal_dish.name, setmeal_dish.copies, dish.image, dish.description").
		Joins("left join dish on setmeal_dish.dish_id = dish.id").
		Where("setmeal_dish.setmeal_id = ?", setmealId).
		Scan(&dishItemVOs).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(errs.RecordNotFoundError)
		}
		panic(errs.DBError)
	}
	return &dishItemVOs
}

// GetSetmealCount 获取某状态套餐数量
func GetSetmealCount(status int) int {
	var setmealCount int64
	if err := global.Db.Table("setmeal").
		Where("status = ?", status).
		Count(&setmealCount).Error; err != nil {
		panic(errs.DBError)
	}
	return int(setmealCount)
}
