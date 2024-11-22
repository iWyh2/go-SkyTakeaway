package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/ulule/deepcopier"
	"go-SkyTakeaway/common/constant"
	errs "go-SkyTakeaway/common/errors"
	"go-SkyTakeaway/global"
	"go-SkyTakeaway/model/dto"
	"go-SkyTakeaway/model/entity"
	"gorm.io/gorm"
	"strconv"
)

// AddShoppingCart 添加购物车
func AddShoppingCart(data *dto.ShoppingCartDTO, ctx *gin.Context) {
	// 创建分类数据模型
	var shoppingCart entity.ShoppingCart
	// 属性拷贝（使用deepcopier库）
	err := deepcopier.Copy(data).To(&shoppingCart)
	if err != nil {
		panic(errs.ServerInternalError)
	}
	// 只能查询自己的购物车数据
	userId, exist := ctx.Get("userId")
	if !exist {
		panic(errs.MissUserIdError)
	}
	id, _ := strconv.Atoi(userId.(string))
	shoppingCart.UserId = id
	// 查询购物车，判断当前商品是否在购物车中
	shoppingCartList := queryShoppingCart(shoppingCart)
	// 如果已经存在，就更新数量，数量加1
	if shoppingCartList != nil && len(shoppingCartList) == 1 {
		shoppingCart = shoppingCartList[0]
		shoppingCart.Number++
		// 更新购物车
		if err := global.Db.Updates(&shoppingCart).Error; err != nil {
			panic(errs.DBError)
		}
	} else {
		// 如果不存在，插入数据，数量就是1
		// 判断当前添加到购物车的是菜品还是套餐
		dishId := data.DishId
		// 添加到购物车的是菜品
		if dishId != 0 {
			// 查询添加的菜品信息
			dish := queryDishByID(strconv.Itoa(dishId))
			shoppingCart.Name = dish.Name
			shoppingCart.Image = dish.Image
			price, _ := strconv.ParseFloat(dish.Price, 64)
			shoppingCart.Amount = price
		} else {
			// 添加到购物车的是套餐
			// 查询添加的套餐信息
			setmeal := querySetmealByID(strconv.Itoa(data.SetmealId))
			shoppingCart.Name = setmeal.Name
			shoppingCart.Image = setmeal.Image
			price, _ := strconv.ParseFloat(setmeal.Price, 64)
			shoppingCart.Amount = price
		}
		shoppingCart.Number = 1
		// 插入购物车
		if err := global.Db.Create(&shoppingCart).Error; err != nil {
			panic(errs.DBError)
		}
	}
}

// QueryShoppingCart 查看购物车
func QueryShoppingCart(ctx *gin.Context) *[]entity.ShoppingCart {
	// 创建数据容器
	var shoppingCartList []entity.ShoppingCart
	// 获取已登录用户的userid
	userId, exist := ctx.Get(constant.UserID)
	if !exist {
		panic(errs.MissUserIdError)
	}
	// 查询
	id, _ := strconv.Atoi(userId.(string))
	shoppingCartList = queryShoppingCart(entity.ShoppingCart{UserId: id})
	return &shoppingCartList
}

// 查询购物车数据
func queryShoppingCart(shoppingCart entity.ShoppingCart) []entity.ShoppingCart {
	// 准备数据容器
	var shoppingCartList []entity.ShoppingCart
	// 查询购物车
	if err := global.Db.Table("shopping_cart").
		Where(&shoppingCart).
		Order("create_time desc").
		Find(&shoppingCartList).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		panic(errs.DBError)
	}
	return shoppingCartList
}

// CleanShoppingCart 清空购物车
func CleanShoppingCart(ctx *gin.Context) string {
	// 获取已登录用户id
	userId, exist := ctx.Get(constant.UserID)
	if !exist {
		panic(errs.MissUserIdError)
	}
	if err := global.Db.Table("shopping_cart").
		Where("user_id = ?", userId.(string)).
		Delete(&entity.ShoppingCart{}).Error; err != nil {
		panic(errs.DBError)
	}
	return userId.(string)
}

// SubShoppingCart 减少购物车
func SubShoppingCart(data *dto.ShoppingCartDTO, ctx *gin.Context) {
	// 创建分类数据模型
	var shoppingCart entity.ShoppingCart
	// 属性拷贝（使用deepcopier库）
	err := deepcopier.Copy(data).To(&shoppingCart)
	if err != nil {
		panic(errs.ServerInternalError)
	}
	// 只能查询自己的购物车数据
	userId, exist := ctx.Get("userId")
	if !exist {
		panic(errs.MissUserIdError)
	}
	id, _ := strconv.Atoi(userId.(string))
	shoppingCart.UserId = id
	// 查询购物车，判断当前商品是否在购物车中
	shoppingCartList := queryShoppingCart(shoppingCart)
	// 查询到当前这条商品的购物车数据 获取出来
	if shoppingCartList != nil && len(shoppingCartList) != 0 {
		shoppingCart = shoppingCartList[0]
		// 当前商品在购物车中的份数为1，直接删除当前记录
		if shoppingCart.Number == 1 {
			deleteShoppingCartById(shoppingCart)
		} else {
			// 当前商品在购物车中的份数不为1，修改份数即可
			shoppingCart.Number--
			// 更新购物车
			if err := global.Db.Updates(&shoppingCart).Error; err != nil {
				panic(errs.DBError)
			}
		}
	}
}

// 根据id删除购物车
func deleteShoppingCartById(shoppingCart entity.ShoppingCart) {
	// 删除购物车
	if err := global.Db.Delete(&shoppingCart).Error; err != nil {
		panic(errs.DBError)
	}
}

// InsertBatchShoppingCart 批量插入购物车数据
func InsertBatchShoppingCart(shoppingCartList []entity.ShoppingCart) {
	if err := global.Db.Create(&shoppingCartList).Error; err != nil {
		panic(errs.DBError)
	}
}
