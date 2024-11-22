package adminController

import (
	"github.com/gin-gonic/gin"
	iUtils "github.com/iWyh2/go-myUtils/utils"
	errs "go-SkyTakeaway/common/errors"
	"go-SkyTakeaway/common/utils"
	model "go-SkyTakeaway/model/result"
	"io"
	"log"
	"net/http"
	"path/filepath"
)

// UploadFile 文件上传
func UploadFile(ctx *gin.Context) {
	// 获取上传的单个文件
	uploadFile, _ := ctx.FormFile("file")
	if uploadFile == nil {
		panic(errs.NoFileError)
	}
	// 日志打印
	log.Printf("文件上传: [%v]", uploadFile.Filename)
	// 打开文件
	file, err := uploadFile.Open()
	if err != nil {
		panic(errs.ServerInternalError)
	}
	defer file.Close()
	// 读取文件内容字节
	fileData, err := io.ReadAll(file)
	if err != nil {
		panic(errs.ServerInternalError)
	}
	// 截取原始文件名的后缀[.png]
	extension := filepath.Ext(uploadFile.Filename)
	// 构造新文件名称
	objectName := iUtils.UUID() + extension
	// 获取文件的请求路径
	path := utils.UploadFile(fileData, objectName)
	// 创建统一返回结果
	var result model.Result[string]
	// 响应
	ctx.JSON(http.StatusOK, result.SuccessByData(path))
}
