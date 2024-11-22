package utils

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"go-SkyTakeaway/common/constant"
	errs "go-SkyTakeaway/common/errors"
	"go-SkyTakeaway/global"
	"go-SkyTakeaway/model/entity"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

// MD5encrypt MD5加密
func MD5encrypt(pwd string) string {
	// 创建一个MD5哈希对象
	h := md5.New()
	// 计算字符串的MD5哈希值
	h.Write([]byte(pwd))
	sum := h.Sum(nil)
	// 将哈希值转为字符串返回
	return hex.EncodeToString(sum)
}

// AutoFillEmpID 自动填充公共字段（CreateUser/UpdateUser）
func AutoFillEmpID(OPType string, ctx *gin.Context, data any) {
	// 取出当前已登录员工的ID
	empId, exists := ctx.Get("empId")
	if !exists {
		panic(errs.ContextNoIDError)
	}
	// 将 字符串 id 转为 int id
	empIdd, err := strconv.Atoi(empId.(string))
	if err != nil {
		panic(errs.ServerInternalError)
	}
	// 根据 数据类型 & 操作类型 进行不同处理
	switch data.(type) {
	case *entity.Employee:
		employee := data.(*entity.Employee)
		if OPType == constant.Create {
			// 设置当前记录创建人id和修改人id
			employee.CreateUser = empIdd
			employee.UpdateUser = empIdd
		} else if OPType == constant.Update {
			// 设置当前记录修改人id
			employee.UpdateUser = empIdd
		} else {
			// 操作类型无效
			panic(errs.ServerInternalError)
		}
	case *entity.Category:
		category := data.(*entity.Category)
		if OPType == constant.Create {
			// 设置当前记录创建人id和修改人id
			category.CreateUser = empIdd
			category.UpdateUser = empIdd
		} else if OPType == constant.Update {
			// 设置当前记录修改人id
			category.UpdateUser = empIdd
		} else {
			// 操作类型无效
			panic(errs.ServerInternalError)
		}
	case *entity.Dish:
		dish := data.(*entity.Dish)
		if OPType == constant.Create {
			// 设置当前记录创建人id和修改人id
			dish.CreateUser = empIdd
			dish.UpdateUser = empIdd
		} else if OPType == constant.Update {
			// 设置当前记录修改人id
			dish.UpdateUser = empIdd
		} else {
			// 操作类型无效
			panic(errs.ServerInternalError)
		}
	case *entity.Setmeal:
		setmeal := data.(*entity.Setmeal)
		if OPType == constant.Create {
			// 设置当前记录创建人id和修改人id
			setmeal.CreateUser = empIdd
			setmeal.UpdateUser = empIdd
		} else if OPType == constant.Update {
			// 设置当前记录修改人id
			setmeal.UpdateUser = empIdd
		} else {
			// 操作类型无效
			panic(errs.ServerInternalError)
		}
	}
}

// CleanCache 清理Redis缓存
func CleanCache(pattern string) {
	// 根据正则获取符合的所有key
	keys, _ := global.RedisClient.Keys(pattern).Result()
	// 根据key删除缓存
	global.RedisClient.Del(keys...)
}

// DoGET 发送GET方式请求
func DoGET(targetUrl string, values map[string]string) string {
	// 将字符串url解析为URL结构
	u, _ := url.ParseRequestURI(targetUrl)
	// 准备query参数
	urlVal := url.Values{}
	// 遍历传入的query参数
	for k, v := range values {
		urlVal.Set(k, v)
	}
	// 为URL添加query参数
	u.RawQuery = urlVal.Encode()
	// 发送GET请求，获取响应
	resp, err := http.Get(u.String())
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// 返回响应数据
	return string(body)
}
