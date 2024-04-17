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
	router.POST("api/user/login", handleUserLogin)
	router.POST("api/user/register", handleUserRegister)

	router.POST("api/service/image/upload", handleImageUpload)
	router.POST("api/service/content/upload", handleCOSContentUpload)

	router.POST("api/article/upload", handleArticleUpload)
	router.GET("api/article/count", handleArticleCount)
	router.GET("api/article/id/list", handleArticleIdList)
	router.POST("api/article/details", handleArticleDetails)
	router.POST("api/article/delete", handleArticleDelete)
	router.POST("api/article/update", handleArticleUpdate)
	router.POST("api/article/visible", handleArticleVisible)

	router.GET("api/anounce/count", handleAncouceCount)
	router.POST("api/anounce/upload", handleAnounceUpload)
	router.GET("api/anounce/id/list", handleAnounceIdList)
	router.POST("api/anounce/details", handleAnounceDetails)
	router.POST("api/anounce/delete", handleAnounceDelete)
	router.POST("api/anounce/update", handleAnounceUpdate)
	router.POST("api/anounce/visible", handleAnounceVisible)

	// 通信加密相关
	router.Run(":6521")
}
