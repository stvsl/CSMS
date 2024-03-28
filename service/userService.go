package service

import (
	"github.com/gin-gonic/gin"
	"stvsljl.com/CSMS/db"
)

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
