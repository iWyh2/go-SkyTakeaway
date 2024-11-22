# 《苍穹外卖》Go语言重构版

## 项目介绍

此项目来自于B站黑马程序员开源Java后端教学项目《苍穹外卖》。

项目简略介绍为：外卖管理系统+微信小程序点单下单程序。

此项目是一个前后端分离的项目。我使用go语言重构了项目的后端代码。

## 技术介绍

我进行重构使用的技术栈如下：

编程语言 -> golang

web框架 -> gin

数据库交互工具 -> gorm

缓存交互工具 -> go-redis

权限校验工具 -> golang-jwt

websocket工具 -> gorilla

数据对象拷贝工具 -> deepcopier

excel文件处理工具 - > excelize

oss存储工具 -> aliyun-oss-go-sdk

配置文件读取工具 -> viper

等等...

## 项目启动

我的golang版本：go1.22.2

1.下载这个仓库源码

2.下载项目所需全部依赖

```go
// 根据实际情况
go mod tidy
```

3.进入./config目录下，打开config.ymal，填写必要配置（数据库，redis...）

```yaml
# 服务器端口配置
server:
  port: :8080
# 数据库配置
database:
  dsn: root:root@tcp(127.0.0.1:3306)/db_name?charset=utf8mb4&parseTime=True&loc=Local
  MaxIdleConns: 10
  MaxOpenConns: 100
# JWT配置
jwt:
  # 设置jwt签名加密时使用的秘钥
  secretKey: your_secretKey
  # 设置前端传递过来的令牌名称
  adminTokenName: your_adminTokenName
  # 微信小程序用户端登录设置令牌
  userTokenName: your_userTokenName
# 阿里OSS配置
aliOSS:
  endpoint: your_endpoint
  accessKeyId: your_accessKeyId
  accessKeySecret: your_accessKeySecret
  bucketName: your_bucketName
# redis配置
redis:
  host: localhost
  # password:
  port: 6379
  database: 0
# 微信小程序所需配置
wechat:
  # 微信登录所需配置
  # 小程序的appid
  appid: your_appid
  # 小程序的秘钥
  secret: your_secret
  # 微信支付所需配置
  # 商户号
  mchid: your_mchid
  # 商户API证书的证书序列号
  mchSerialNo: your_mchSerialNo
  # 商户私钥文件
  privateKeyFilePath: your_privateKeyFilePath
  # 证书解密的密钥
  apiV3Key: your_apiV3Key
  # 平台证书
  weChatPayCertFilePath: your_weChatPayCertFilePath
  # 支付成功的回调地址
  # cpolar生成的临时公网地址 https://xxxxxx.cpolar.top
  # 每次测试使用到微信支付功能都需要重新生成并改写
  notifyUrl: https://xxxxxx.cpolar.top/notify/paySuccess
  # 退款成功的回调地址
  refundNotifyUrl: https://xxxxxx.cpolar.top/notify/refundSuccess
```

4.再打开./config/config.go，更改需要读取的配置文件

```go
// 设置文件名
viper.SetConfigName("config-dev")
// 更改为如下
viper.SetConfigName("config")
```

5.启动服务器即可

```go
// 你也可以编译后启动
// 启动前需要确保配置正确，以及mysql与redis启动成功
go run ./server.go
```

------

> 注：此项目只用于让学习go语言的同学了解go语言后端开发技术栈
>
> ©2024 iWyh2