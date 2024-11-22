package service

import (
	"errors"
	errs "go-SkyTakeaway/common/errors"
	"go-SkyTakeaway/global"
	"go-SkyTakeaway/model/entity"
	"gorm.io/gorm"
)

// 根据菜品id插入口味数据
func addFlavorByDishID(flavors []entity.DishFlavor, dishID int) {
	if len(flavors) == 0 {
		return
	}
	// 向口味数据添加菜品id
	for i := 0; i < len(flavors); i++ {
		flavors[i].DishID = dishID
	}
	// 创建数据库表数据，口味
	if err := global.Db.Create(flavors).Error; err != nil {
		panic(errs.DBError)
	}
}

// 根据菜品id查询口味数据
func getFlavorByDishID(id string) []entity.DishFlavor {
	// 准备数据容器
	var flavors []entity.DishFlavor
	// 执行查询
	if err := global.Db.Where("dish_id = ?", id).Find(&flavors).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(errs.RecordNotFoundError)
		}
		panic(errs.DBError)
	}
	return flavors
}

// 根据菜品id删除口味数据
func deleteFlavorByDishID(ids []string) {
	// 执行删除
	if err := global.Db.Where("dish_id in (?)", ids).
		Delete(&entity.DishFlavor{}).Error; err != nil {
		panic(errs.DBError)
	}
}
