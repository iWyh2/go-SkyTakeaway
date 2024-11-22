package config

import (
	"github.com/go-redis/redis"
	errs "go-SkyTakeaway/common/errors"
	"go-SkyTakeaway/global"
)

// redis配置
func redisConfig() {
	// 创建redis客户端
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     ServerConfig.Redis.Host + ":" + ServerConfig.Redis.Port,
		Password: "",
		DB:       ServerConfig.Redis.Database,
	})
	// 检查是否连接成功
	_, err := RedisClient.Ping().Result()
	if err != nil {
		panic(errs.RedisConnectError)
	}
	// 赋值给全局变量
	global.RedisClient = RedisClient
}
