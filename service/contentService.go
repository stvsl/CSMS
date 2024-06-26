package service

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

func handleArticleIdCount(c *gin.Context) {
	var count int64
	db.GetConn().Model(&db.Article{}).Where("status = ?", 1).Count(&count)
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

func handleArticleDelete(c *gin.Context) {
	sid, derr := c.GetQuery("id")
	id, err := strconv.Atoi(sid)
	if !derr || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误" + sid,
		})
		return
	}
	if err := db.GetConn().Model(&db.Article{}).Where("aid = ?", id).Delete(&db.Article{}).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误或服务器内部错误" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "成功",
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

func handleAncouceCount(c *gin.Context) {
	var count int64
	if err := db.GetConn().Model(&db.Anounce{}).Count(&count).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "服务器内部错误" + err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "请求成功",
		"count":   count,
	})
}

func handleAncouceIdCount(c *gin.Context) {
	var count int64
	if err := db.GetConn().Model(&db.Anounce{}).Where("status = ?", 1).Count(&count).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "服务器内部错误" + err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "请求成功",
		"count":   count,
	})
}

func handleAnounceIdList(c *gin.Context) {
	spage := c.Query("page")
	page, err := strconv.Atoi(spage)
	if page == 0 || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "请求参数错误" + err.Error(),
		})
		return
	}
	type IDS struct {
		Aid int `gorm:"autoIncrement:true;primaryKey;column:aid;type:int(12);not null;comment:'文章编号'"` // 文章编号
	}
	var ids []IDS
	if err := db.GetConn().Model(&db.Anounce{}).Select("aid").Limit(10).Offset((page - 1) * 10).Find(&ids).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "服务器内部错误" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "请求成功",
		"data":    ids,
	})
}

func handleAnounceDetails(c *gin.Context) {
	sid, derr := c.GetQuery("id")
	if !derr {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "请求参数错误",
		})
		return
	}
	anounce := db.Anounce{}
	if err := db.GetConn().Model(&db.Anounce{}).Where("aid = ?", sid).First(&anounce).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "参数错误" + err.Error(),
		})
		return
	}
	db.GetConn().Model(&anounce).Where("aid = ?", sid).UpdateColumn("pageviews", gorm.Expr("pageviews + ?", 1))
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "请求成功",
		"data":    anounce,
	})
}

func handleArticleUpdate(c *gin.Context) {
	type Req struct {
		Aid  string `json:"aid"`
		Data struct {
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
		Anchor  string `json:"anchor"`
	}
	var req Req
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	mgr := db.ArticleMgr(db.GetConn())
	if err := mgr.Where("aid = ?", req.Aid).Updates(&db.Article{
		Title:        req.Data.Title,
		Introduction: req.Data.Introduction,
		Text:         req.Content,
		Coverimg:     req.Data.Coverimage[0].URL,
		Contentimg:   req.Data.Contentimage[0].URL,
		Updatetime:   time.Now(),
		Author:       req.Anchor,
	}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "更新文章失败" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "更新文章成功",
	})
}

func handleArticleVisible(c *gin.Context) {
	type Req struct {
		Aid    string `json:"aid"`
		Status int    `json:"status"`
	}
	var req Req
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.GetConn().Model(&db.Article{}).Where("aid = ?", req.Aid).UpdateColumn("status", req.Status).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "更新文章状态失败" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "更新文章状态成功",
	})
}

func handleAnounceDelete(c *gin.Context) {
	sid := c.Query("id")
	if err := db.GetConn().Model(&db.Anounce{}).Where("aid = ?", sid).Delete(&db.Anounce{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "删除公告失败" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "删除公告成功",
	})
}

func handleAnounceUpdate(c *gin.Context) {
	type Req struct {
		Aid  string `json:"aid"`
		Data struct {
			Title        string `json:"title"`
			Introduction string `json:"introduction"`
		} `json:"data"`
		Content string `json:"content"`
		Anchor  string `json:"anchor"`
	}
	var req Req
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.GetConn().Model(&db.Anounce{}).Where("aid = ?", req.Aid).Updates(&db.Anounce{
		Title:        req.Data.Title,
		Introduction: req.Data.Introduction,
		Text:         req.Content,
		Updatetime:   time.Now(),
		Author:       req.Anchor,
	}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "更新公告失败" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "更新公告成功",
	})
}

func handleAnounceVisible(c *gin.Context) {
	type Req struct {
		Aid    string `json:"aid"`
		Status int    `json:"status"`
	}
	var req Req
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.GetConn().Model(&db.Anounce{}).Where("aid = ?", req.Aid).UpdateColumn("status", req.Status).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "更新公告状态失败" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "更新公告状态成功",
	})
}

func handleArticleRecent(c *gin.Context) {
	var resid []db.Article
	if err := db.GetConn().Table("Article").Order("writetime desc").Limit(5).Find(&resid).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "获取最新文章失败" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "获取最新文章成功",
		"data":    resid,
	})
}

func handleArticleAll(c *gin.Context) {
	spage := c.Query("page")
	page, err := strconv.Atoi(spage)
	if err != nil || page < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "页码错误",
		})
		return
	}
	var resid []int
	if err := db.GetConn().Model(&db.Article{}).Where("status = ?", 1).Offset((page-1)*10).Limit(10).Pluck("aid", &resid).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "获取文章失败" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "获取文章成功",
		"data":    resid,
	})
}

func handleAnounceRecent(c *gin.Context) {
	var res []int
	if err := db.GetConn().Table("Anounce").Select("aid").Order("writetime desc").Limit(5).Find(&res).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "获取公告失败" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "获取公告成功",
		"data":    res,
	})
}

func handleAncouceAll(c *gin.Context) {
	spage := c.Query("page")
	page, err := strconv.Atoi(spage)
	if err != nil || page < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "页码错误",
		})
		return
	}
	var res []int
	if err := db.GetConn().Table("Anounce").Select("aid").Offset((page - 1) * 10).Limit(10).Find(&res).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "获取公告失败" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "获取公告成功",
		"data":    res,
	})
}

func handleArticleOverview(c *gin.Context) {
	type Res struct {
		ArticleCount      int64 `json:"articleCount"`
		ArticleHasPublish int64 `json:"articlePublish"`
		AnounceCount      int64 `json:"anounceCount"`
		AnounceHasPublish int64 `json:"anouncePublish"`
	}
	var res Res
	db.GetConn().Table("Article").Count(&res.ArticleCount)
	db.GetConn().Table("Article").Where("status = ?", 1).Count(&res.ArticleHasPublish)
	db.GetConn().Table("Anounce").Count(&res.AnounceCount)
	db.GetConn().Table("Anounce").Where("status = ?", 1).Count(&res.AnounceHasPublish)
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "获取文章概览成功",
		"data":    res,
	})
}
