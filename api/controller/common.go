package controller

import (
	response "dpj-admin-api/support"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func UploadFile(c *gin.Context) {

	// 获取FormFile
	file, header, err := c.Request.FormFile("upload")
	if err != nil {
		response.WithContext(c).Error(400, fmt.Sprintf("上传文件失败: %s", err.Error()))
		return
	}

	// 获取文件名，并创建新的文件存储
	filename := header.Filename

	//临时存储路径
	path := "./upload/"

	out, err := os.Create(path + filename)
	if err != nil {
		response.WithContext(c).Error(400, fmt.Sprintf("创建文件: %s", err.Error()))
		return
	}

	defer out.Close()
	//将读取的文件流写到文件中
	_, err = io.Copy(out, file)
	if err != nil {
		response.WithContext(c).Error(400, fmt.Sprintf("读取文件失败: %s", err.Error()))
		return
	}

	response.WithContext(c).Success(gin.H{
		"msg": "上传成功",
	})
}
