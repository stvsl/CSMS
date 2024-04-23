package service

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"stvsljl.com/CSMS/db"
)

func handleAccount3rdCount(c *gin.Context) {
	var count int64
	if err := db.GetConn().Model(&db.Otheruser{}).Count(&count).Error; err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"code": 500,
			"msg":  "fail",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": count,
	})
}

func handleAccount3rdIdList(c *gin.Context) {
	spage := c.Query("page")
	page, err := strconv.Atoi(spage)
	if err != nil || page < 1 {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"code": 500,
			"msg":  "请求参数错误",
		})
		return
	}
	var account3rd []int
	if err := db.GetConn().Model(&db.Otheruser{}).Offset((page-1)*10).Limit(10).Pluck("oid", &account3rd).Error; err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"code": 500,
			"msg":  "fail",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": account3rd,
	})
}

func handleAccount3rdDetail(c *gin.Context) {
	sid := c.Query("id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"code": 500,
			"msg":  "请求参数无效",
		})
		return
	}
	account := db.Otheruser{}
	if err := db.GetConn().Model(&account).Where("oid = ?", id).First(&account).Error; err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"code": 500,
			"msg":  "fail",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": account,
	})
}

func handleAccount3rdDelete(c *gin.Context) {
	sid := c.Query("id")
	id, err := strconv.Atoi(sid)
	if err != nil || id == 0 {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"code": 500,
			"msg":  "请求参数无效",
		})
		return
	}
	if err := db.GetConn().Model(&db.Otheruser{}).Delete("oid", id).Error; err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"code": 500,
			"msg":  "删除失败",
		})
		return
	}
	if err := db.GetConn().Model(&db.User{}).Delete("uid", id).Error; err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"code": 500,
			"msg":  "删除失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "删除成功",
	})
}

func handleAccount3rdRegister(c *gin.Context) {
	sid := c.Query("id")
	company := c.Query("company")
	id, err := strconv.Atoi(sid)
	if err != nil || id == 0 {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"code": 500,
			"msg":  "参数错误",
		})
		return
	}
	user := db.User{}
	if err := db.GetConn().Model(&user).Where("uid = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"code": 500,
			"msg":  "用户不存在",
		})
		return
	}
	other := db.Otheruser{
		Oid:     user.UID,
		Name:    user.Name,
		Tel:     user.Tel,
		Sex:     user.Sex,
		Company: company,
	}
	if err := db.GetConn().Create(&other).Error; err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"code": 500,
			"msg":  "添加失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "添加成功",
	})
}

func handleAccount3rdCancleRight(c *gin.Context) {
	sid := c.Query("id")
	id, err := strconv.Atoi(sid)
	if err != nil || id == 0 {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"code": 500,
			"msg":  "请求参数无效",
		})
		return
	}
	if err := db.GetConn().Model(&db.Otheruser{}).Delete("oid", id).Error; err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"code": 500,
			"msg":  "删除失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "删除成功",
	})
}

func handleAccountFetchList(c *gin.Context) {
	key := c.Query("key")
	var ids []int
	if key != "" {
		if err := db.GetConn().Model(&db.Otheruser{}).Where("oid LIKE ?", "%"+key+"%").Pluck("oid", &ids).Error; err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{
				"code": 500,
				"msg":  "请求参数无效",
			})
			return
		}
	} else {
		if err := db.GetConn().Model(&db.Otheruser{}).Pluck("oid", &ids).Error; err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{
				"code": 500,
				"msg":  "请求参数无效",
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "请求成功",
		"data": ids,
	})
}
