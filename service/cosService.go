package service

import (
	"github.com/gin-gonic/gin"
	"stvsljl.com/CSMS/cos"
)

func handleImageUpload(c *gin.Context) {
	// 获取文件
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "上传失败",
		})
		return
	}
	// 调用腾讯云COS存储上传文件
	str, err := cos.UploadFile(file)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "上传失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "上传成功",
		"url":  str,
	})
}

func handleCOSContentUpload(c *gin.Context) {
	// 获取文件
	file, err := c.FormFile("wangeditor-uploaded-image")
	if err != nil {
		c.JSON(400, gin.H{
			"errno":   400,
			"message": "上传失败",
		})
		return
	}
	// 调用腾讯云COS存储上传文件
	str, err := cos.UploadFile(file)
	if err != nil {
		c.JSON(400, gin.H{
			"errno":  301,
			"mssage": "上传失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"errno": 0,
		"msg":   "上传成功",
		"data": gin.H{
			"url": str,
		},
	})
}
