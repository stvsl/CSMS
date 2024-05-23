package service

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"stvsljl.com/CSMS/db"
)

func handleAdminLogin(c *gin.Context) {
	type req struct {
		Id     string `json:"id"`
		Passwd string `json:"passwd"`
	}
	//  从请求体中获取参数
	var r req
	err := c.BindJSON(&r)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 0,
			"msg":  "参数错误",
		})
		return
	}
	id, err := strconv.ParseUint(r.Id, 0, 64)
	if err != nil {
		fmt.Print(err)
		c.JSON(400, gin.H{
			"code": 0,
			"msg":  "参数转换失败",
		})
		return
	}
	var admin db.Admin
	mgr := db.AdminMgr(db.GetConn())
	// id是不是电话号码
	if len(r.Id) == 11 {
		admin, err = mgr.GetFromTel(r.Id)
	} else {
		admin, err = mgr.GetFromID(int(id))
	}
	if err != nil {
		c.JSON(400, gin.H{
			"code": 0,
			"msg":  "用户不存在",
		})
		return
	}
	if admin.Passwd != r.Passwd {
		c.JSON(400, gin.H{
			"code": 0,
			"msg":  "密码错误",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "登录成功",
		"info": admin,
	})
}

func handleAdminRegister(c *gin.Context) {
	// TODO;
}

func handleAdminUpdate(c *gin.Context) {
	id := c.Query("id")
	tel := c.Query("tel")
	name := c.Query("name")
	if err := db.GetConn().Model(&db.Admin{}).Where("id = ?", id).UpdateColumns(map[string]interface{}{
		"tel":  tel,
		"name": name,
	}).Error; err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"code": 400,
			"msg":  "修改失败",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "修改成功",
	})
}

func handleAdminPasswd(c *gin.Context) {
	id := c.Query("id")
	passwd := c.Query("passwd")
	if err := db.GetConn().Model(&db.Admin{}).Where("id = ?", id).UpdateColumn("passwd", passwd).Error; err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"code": 400,
			"msg":  "修改失败",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "修改成功",
	})
}

func handleLastLoginTime(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "请求参数错误",
		})
		return
	}
	if err := db.GetConn().Table("Admin").Where("id = ?", id).Update("lastlogin", time.Now()).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "修改失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "修改成功",
	})
}

func handleUserCount(c *gin.Context) {
	var count int64
	if err := db.GetConn().Table("Admin").Count(&count).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "请求参数错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "请求成功",
		"data": count,
	})
}

func handleUserDetails(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "请求参数错误",
		})
		return
	}
	admin := db.Admin{}
	if err := db.GetConn().Table("Admin").Where("id = ?", id).First(&admin).Error; err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"code": 400,
			"msg":  "请求数据库失败",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "请求成功",
		"data": admin,
	})
}
