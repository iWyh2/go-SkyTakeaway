package global

import (
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

var (
	// Db 数据库连接
	Db *gorm.DB
	// RedisClient redis连接客户端
	RedisClient *redis.Client
)
