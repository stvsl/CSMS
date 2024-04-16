package service

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"stvsljl.com/CSMS/db"
)

func handleArticleCount(c *gin.Context) {
	var count int64
	db.GetConn().Model(&db.Article{}).Count(&count)
	c.JSON(http.StatusOK, gin.H{
		"code":  200,
		"msg":   "success",
		"count": count,
	})
}

func handleArticleIdList(c *gin.Context) {
	// 从请求参数中获取
	spage, berr := c.GetQuery("page")
	page, err := strconv.Atoi(spage)
	if !berr || err != nil || page <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求参数错误" + spage + err.Error(),
		})
		return
	}
	type IDS struct {
		Aid int `gorm:"autoIncrement:true;primaryKey;column:aid;type:int(12);not null;comment:'文章编号'"` // 文章编号
	}
	var ids []IDS
	if err := db.GetConn().Model(&db.Article{}).Select("aid").Limit(10).Offset((page - 1) * 10).Find(&ids).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "服务器内部错误",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "请求成功",
			"data": ids,
		})
	}

}

func handleArticleDetails(c *gin.Context) {
	sid, derr := c.GetQuery("id")
	id, err := strconv.Atoi(sid)
	if !derr || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误" + sid,
		})
		return
	}
	var article db.Article
	if err = db.GetConn().Model(&article).Where("aid = ?", id).First(&article).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误或服务器内部错误" + err.Error(),
		})
		return
	}
	article.Pageviews++
	// 存回数据库
	if err = db.GetConn().Model(&article).Where("aid = ?", id).Updates(&article).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误或服务器内部错误" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "成功",
		"data": article,
	})

}

func handleArticleUpload(c *gin.Context) {
	type Req struct {
		Data struct {
			Type         string `json:"type"`
			Title        string `json:"title"`
			Introduction string `json:"introduction"`
			Coverimage   []struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"coverimage"`
			Contentimage []struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"contentimage"`
		} `json:"data"`
		Content string `json:"content"`
		Author  string `json:"author"`
	}
	req := Req{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "参数错误" + err.Error(),
		})
		return
	}
	mgr := db.ArticleMgr(db.GetConn())
	article := db.Article{
		Coverimg:     req.Data.Coverimage[0].URL,
		Contentimg:   req.Data.Contentimage[0].URL,
		Title:        req.Data.Title,
		Introduction: req.Data.Introduction,
		Text:         req.Content,
		Writetime:    time.Now(),
		Updatetime:   time.Now(),
		Author:       req.Author,
		Pageviews:    0,
		Status:       0,
	}
	err = mgr.Create(&article).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "文章创建失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "文章创建成功",
	})
}

func handleAnounceUpload(c *gin.Context) {
	type Req struct {
		Data struct {
			Type         string      `json:"type"`
			Title        string      `json:"title"`
			Introduction string      `json:"introduction"`
			Coverimage   interface{} `json:"coverimage"`
			Contentimage interface{} `json:"contentimage"`
		} `json:"data"`
		Content string `json:"content"`
		Anchor  string `json:"anchor"`
	}
	var req Req
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	anounce := db.Anounce{
		Writetime:    time.Now(),
		Updatetime:   time.Now(),
		Text:         req.Content,
		Status:       1,
		Title:        req.Data.Title,
		Introduction: req.Data.Introduction,
		Pageviews:    0,
		Author:       req.Anchor,
	}
	mgr := db.AnounceMgr(db.GetConn())
	if err := mgr.Create(&anounce).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "创建公告失败" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "创建公告成功",
	})
}
