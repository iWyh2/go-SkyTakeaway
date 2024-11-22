package config

import (
	"go-SkyTakeaway/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

// 数据库配置
func dbConfig() {
	// 从配置文件读取dsn
	dsn := ServerConfig.Database.Dsn
	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	// 配置数据库连接池
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalln(err)
	}
	// 设置连接池中空闲连接的最大数量
	sqlDB.SetMaxIdleConns(ServerConfig.Database.MaxIdleConns)
	// 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(ServerConfig.Database.MaxOpenConns)
	// 设置连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)
	// 将数据库连接赋给全局变量
	global.Db = db
}
