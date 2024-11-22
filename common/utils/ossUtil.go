package utils

import (
	"bytes"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	errs "go-SkyTakeaway/common/errors"
	"go-SkyTakeaway/config"
	"strings"
)

// UploadFile 文件上传
func UploadFile(fileData []byte, objectName string) string {
	// 创建OSSClient实例
	ossClient, err := oss.New(config.ServerConfig.AliOSS.Endpoint,
		config.ServerConfig.AliOSS.AccessKeyId,
		config.ServerConfig.AliOSS.AccessKeySecret)
	if err != nil {
		panic(errs.UploadFileError)
	}
	// 填写存储空间名称
	bucket, err := ossClient.Bucket(config.ServerConfig.AliOSS.BucketName)
	if err != nil {
		panic(errs.UploadFileError)
	}
	// 上传
	err = bucket.PutObject(objectName, bytes.NewReader(fileData))
	if err != nil {
		panic(errs.UploadFileError)
	}
	// 返回文件访问路径
	path := "https://" + config.ServerConfig.AliOSS.BucketName + "." +
		strings.Split(config.ServerConfig.AliOSS.Endpoint, "//")[1] +
		"/" + objectName
	return path
}
