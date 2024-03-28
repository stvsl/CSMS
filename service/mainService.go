package service

import (
	"github.com/gin-gonic/gin"
	"stvsljl.com/CSMS/cos"
	"stvsljl.com/CSMS/db"
	"stvsljl.com/CSMS/utils"
)

func Start() {
	// 数据库事务初始化
	db.Connect()
	// 日志组件初始化
	utils.Log.Init()
	// 对象存储初始化
	cos.Init()
	// 服务器服务启动
	router := gin.Default()
	router.SetTrustedProxies(nil)
	// 允许跨域
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})
	/**********************
	 * 加载路由
	 **********************/
	router.POST("api/admin/login", handleAdminLogin)
	router.POST("api/user/register", handleUserRegister)
	// 通信加密相关
	router.Run(":6521")
}
