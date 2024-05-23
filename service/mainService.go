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
	router.POST("api/admin/register", handleAdminRegister)
	router.POST("api/admin/update", handleAdminUpdate)
	router.POST("api/admin/passwd", handleAdminPasswd)
	router.POST("api/user/login", handleUserLogin)
	router.POST("api/user/register", handleUserRegister)
	router.GET("api/user/fetchlastlogintime", handleLastLoginTime)
	router.GET("api/user/count", handleUserCount)
	router.GET("api/user/details", handleUserDetails)

	router.POST("api/service/image/upload", handleImageUpload)
	router.POST("api/service/content/upload", handleCOSContentUpload)

	router.POST("api/article/upload", handleArticleUpload)
	router.GET("api/article/count", handleArticleCount)
	router.GET("api/article/id/list", handleArticleIdList)
	router.POST("api/article/details", handleArticleDetails)
	router.POST("api/article/delete", handleArticleDelete)
	router.POST("api/article/update", handleArticleUpdate)
	router.POST("api/article/visible", handleArticleVisible)
	router.GET("api/article/id/recent", handleArticleRecent)
	router.GET("api/article/id/all", handleArticleAll)
	router.GET("api/article/id/count", handleArticleIdCount)
	router.GET("api/articleanounce/admin/overview", handleArticleOverview)

	router.GET("api/anounce/count", handleAncouceCount)
	router.POST("api/anounce/upload", handleAnounceUpload)
	router.GET("api/anounce/id/list", handleAnounceIdList)
	router.POST("api/anounce/details", handleAnounceDetails)
	router.POST("api/anounce/delete", handleAnounceDelete)
	router.POST("api/anounce/update", handleAnounceUpdate)
	router.POST("api/anounce/visible", handleAnounceVisible)
	router.GET("api/anounce/id/recent", handleAnounceRecent)
	router.GET("api/anounce/id/all", handleAncouceAll)
	router.GET("api/anounce/id/count", handleAncouceIdCount)

	router.POST("api/activity/upload", handleActivityUpload)
	router.GET("api/activity/count", handleActivityCount)
	router.GET("api/activity/id/list", handleActivityIdList)
	router.POST("api/activity/details", handleActivityDetails)
	router.POST("api/activity/status", handleActivityStatus)
	router.POST("api/activity/delete", handleActivityDelete)
	router.POST("api/activity/peoplestatus", handleActivityPeopleStatus)
	router.POST("api/activity/update", handleActivityUpdate)
	router.POST("api/activity/recent/ids", handleActivityRecentIdList)
	router.POST("api/activity/hasjoin", handleActivityHasjion)
	router.POST("api/activity/user/status", handleActivityUserStatus)
	router.POST("api/activity/join", handleActivityJoin)
	router.POST("api/activity/exit", handleActivityExit)
	router.GET("api/activity/user/count", handleActivityUserCount)
	router.GET("api/activity/user/id/list", handleActivityUserIdList)
	router.GET("api/activity/admin/overview", handleActivityOverview)

	router.POST("api/feed/upload", handleFeedUpload)
	router.GET("api/feed/count", handleFeedCount)
	router.GET("api/feed/id/list", handleFeedIdList)
	router.POST("api/feed/details", handleFeedDetails)
	router.POST("api/feed/delete", handleFeedDelete)
	router.POST("api/feed/update", handleFeedUpdate)
	router.POST("api/feed/status", handleFeedStatus)
	router.POST("api/feed/process", handleFeedProcess)
	router.POST("api/feed/do", handleFeedDo)
	router.GET("api/feed/admin/overview", handleFeedAdminOverview)
	router.GET("api/feed/user/overview", handleFeedUserOverview)
	router.GET("api/feed/user/count", handleFeedUserCount)
	router.GET("api/feed/user/id/list", handleFeedUserIdList)

	router.POST("api/fix/upload", handleFixUpload)
	router.GET("api/fix/count", handleFixCount)
	router.GET("api/fix/id/list", handleFixIdList)
	router.POST("api/fix/details", handleFixDetails)
	router.POST("api/fix/delete", handleFixDelete)
	router.POST("api/fix/update", handleFixUpdate)
	router.POST("api/fix/status", handleFixStatus)
	router.POST("api/fix/process", handleFixProcess)
	router.POST("api/fix/do", handleFixDo)
	router.GET("api/fix/admin/overview", handleFixAdminOverview)
	router.GET("api/fix/user/overview", handleFixUserOverview)
	router.GET("api/fix/user/count", handleFixUserCount)
	router.GET("api/fix/user/id/list", handleFixUserIdList)

	router.GET("api/account/count", handleAccountCount)
	router.GET("api/account/id/list", handleAccountIdList)
	router.POST("api/account/detail", handleAccountDetail)
	router.POST("api/account/update/passwd", handleAccountUpdatePasswd)
	router.POST("api/account/update/info", handleAccountUpdateInfo)
	router.POST("api/account/delete", handleAccountDelete)
	router.POST("api/account/getbykey", handleAccountIdsGetByKey)
	router.GET("api/account/admin/overview", handleAccountOverview)
	router.POST("api/account/add",handleAccountAdd)

	router.GET("api/account/3rd/count", handleAccount3rdCount)
	router.GET("api/account/3rd/id/list", handleAccount3rdIdList)
	router.POST("api/account/3rd/detail", handleAccount3rdDetail)
	router.POST("api/account/3rd/delete", handleAccount3rdDelete)
	router.POST("api/account/3rd/register", handleAccount3rdRegister)
	router.POST("api/account/3rd/cancelRight", handleAccount3rdCancleRight)
	router.GET("api/account/3rd/id/fetchlist", handleAccountFetchList)

	// 打印总路由条数
	utils.Log.Info("总路由条数：", len(router.Routes()))
	router.Run(":6521")
}
