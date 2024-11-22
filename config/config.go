package config

import (
	"github.com/spf13/viper"
)

// Config config.yaml配置类
type Config struct {
	Server struct {
		Port string
	}
	Database struct {
		Dsn          string
		MaxIdleConns int
		MaxOpenConns int
	}
	Jwt struct {
		SecretKey      string
		AdminTokenName string
		UserTokenName  string
	}
	AliOSS struct {
		Endpoint        string
		AccessKeyId     string
		AccessKeySecret string
		BucketName      string
	}
	Redis struct {
		Host     string
		Port     string
		Database int
	}
	WeChat struct {
		Appid                 string
		Secret                string
		Mchid                 string
		MchSerialNo           string
		PrivateKeyFilePath    string
		ApiV3Key              string
		WeChatPayCertFilePath string
		NotifyUrl             string
		RefundNotifyUrl       string
	}
}

// ServerConfig 服务器配置信息
var ServerConfig = new(Config)

// Init 初始化服务器配置信息
func Init() {
	// 设置文件名
	viper.SetConfigName("config-dev")
	// 设置文件类型
	viper.SetConfigType("yaml")
	// 添加配置文件路劲
	viper.AddConfigPath("./config")
	// 读取配置文件内容
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	// 将配置文件内容解析到结构体（定义好的符合配置文件的结构体）
	if err := viper.Unmarshal(ServerConfig); err != nil {
		panic(err)
	}
	// 初始化db配置
	dbConfig()
	// 初始化redis配置
	redisConfig()
}
