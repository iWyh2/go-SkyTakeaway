package service

import (
	"errors"
	errs "go-SkyTakeaway/common/errors"
	"go-SkyTakeaway/global"
	"go-SkyTakeaway/model/entity"
	"gorm.io/gorm"
)

// 根据菜品id获取关联的套餐id
func getSetmealIDsByDishIDs(ids []string) []int {
	// 准备数据容器
	setmealIDs := make([]int, 0)
	// 准备查询条件
	query := global.Db.Table("setmeal_dish").
		Select("setmeal_id").
		Where("dish_id in (?)", ids)
	// 执行查询
	if err := query.Find(&setmealIDs).Error; err != nil {
		panic(errs.DBError)
	}
	return setmealIDs
}

// 根据套餐id获取套餐菜品关联数据
func querySetmealDishesBySetmealID(id string) *[]entity.SetmealDish {
	// 准备数据容器
	var setmealDishes []entity.SetmealDish
	// 执行查询
	if err := global.Db.Where("setmeal_id = ?", id).Find(&setmealDishes).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(errs.RecordNotFoundError)
		}
		panic(errs.DBError)
	}
	return &setmealDishes
}

// 保存套餐和菜品的关联关系
func insertSetmealDishes(setmealDishes []entity.SetmealDish, setmealId int) {
	// 填充套餐id
	for i := 0; i < len(setmealDishes); i++ {
		setmealDishes[i].SetmealID = setmealId
	}
	// 创建数据
	if err := global.Db.Create(&setmealDishes).Error; err != nil {
		panic(errs.DBError)
	}
}

// 删除套餐和菜品的关联关系
func deleteSetmealDishBySetmealIDs(ids []string) {
	// 执行删除
	if err := global.Db.Where("setmeal_id in (?)", ids).
		Delete(&entity.SetmealDish{}).Error; err != nil {
		panic(errs.DBError)
	}
}
