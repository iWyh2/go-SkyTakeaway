package service

import (
	"errors"
	"github.com/go-redis/redis"
	"go-SkyTakeaway/common/constant"
	errs "go-SkyTakeaway/common/errors"
	"go-SkyTakeaway/global"
)

// SetShopStatus 设置营业状态
func SetShopStatus(status string) {
	// 设置缓存，过期时间为0表示不会过期
	if err := global.RedisClient.Set(constant.RedisKey, status, 0).Err(); err != nil {
		panic(errs.RedisError)
	}
}

// GetShopStatus 查询店铺营业状态
func GetShopStatus() string {
	// 查询缓存
	status, err := global.RedisClient.Get(constant.RedisKey).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			panic(errs.RedisNotFoundError)
		} else {
			panic(errs.RedisError)
		}
	}
	return status
}
