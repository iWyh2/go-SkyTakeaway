package utils

import (
	"encoding/json"
	"go-SkyTakeaway/common/constant"
	errs "go-SkyTakeaway/common/errors"
	"go-SkyTakeaway/config"
)

// GetOpenID 调用微信接口服务，获取微信用户的openid
func GetOpenID(code string) string {
	// 准备微信接口服务参数
	values := map[string]string{
		"appid":      config.ServerConfig.WeChat.Appid,
		"secret":     config.ServerConfig.WeChat.Secret,
		"js_code":    code,
		"grant_type": "authorization_code",
	}
	// 调用微信接口服务，获得返回数据
	resultJSON := DoGET(constant.WxLogin, values)
	// 创建请求成功返回数据模型
	wxLoginResData := struct {
		SessionKey string `json:"session_key"`
		OpenID     string `json:"openid"`
	}{}
	// 解析返回数据
	if err := json.Unmarshal([]byte(resultJSON), &wxLoginResData); err != nil {
		panic(errs.WxLoginFailedError)
	}
	// 返回openid
	return wxLoginResData.OpenID
}
