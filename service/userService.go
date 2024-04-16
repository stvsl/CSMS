package service

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"stvsljl.com/CSMS/db"
)

func handleUserLogin(c *gin.Context) {
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
	var user db.User
	mgr := db.UserMgr(db.GetConn())
	// id是不是电话号码
	if len(r.Id) == 11 {
		user, err = mgr.GetFromTel(r.Id)
	} else {
		user, err = mgr.GetFromUID(int(id))
	}
	if err != nil {
		c.JSON(400, gin.H{
			"code": 0,
			"msg":  "用户不存在",
		})
		return
	}
	if user.Passwd != r.Passwd {
		c.JSON(400, gin.H{
			"code": 0,
			"msg":  "密码错误",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "登录成功",
		"info": user,
	})
}
func handleUserRegister(c *gin.Context) {
	type Req struct {
		Name   string `json:"name"`
		Tel    string `json:"tel"`
		Sex    int    `json:"sex"`
		Passwd string `json:"passwd"`
	}
	var req Req
	c.BindJSON(&req)
	user := db.User{
		Name:   req.Name,
		Tel:    req.Tel,
		Sex:    req.Sex,
		Passwd: req.Passwd,
	}
	conn := db.GetConn()
	err := conn.Create(&user).Error
	if err != nil {
		c.JSON(400, gin.H{
			"code": 0,
			"msg":  "failed",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "succeed",
		"info": user,
	})
}
