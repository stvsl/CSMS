package service

import (
	"fmt"
	"net/http"
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

func handleAccountCount(c *gin.Context) {
	var count int64
	if err := db.GetConn().Model(&db.User{}).Count(&count).Error; err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"code": 200,
			"msg":  "failed",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "succeed",
		"info": count,
	})
}

func handleAccountIdList(c *gin.Context) {
	spage := c.Query("page")
	page, err := strconv.Atoi(spage)
	if spage == "" || err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"code": 400,
			"msg":  "请求格式错误",
		})
		return
	}
	var ids []string
	if err := db.GetConn().Model(&db.User{}).Offset((page-1)*10).Pluck("uid", &ids).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "服务器错误",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "请求成功",
		"data": ids,
	})
}

func handleAccountDetail(c *gin.Context) {
	var id = c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求参数错误",
		})
		return
	}
	user := &db.User{}
	if err := db.GetConn().Model(user).Where("uid=?", id).Find(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "服务器错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "请求成功",
		"data": user,
	})
}

func handleAccountUpdatePasswd(c *gin.Context) {
	type Req struct {
		ID     string `json:"id"`
		Passwd string `json:"passwd"`
	}
	var req Req
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求参数错误",
		})
		return
	}
	if req.ID == "" || req.Passwd == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求参数异常",
		})
		return
	}
	if err := db.GetConn().Model(&db.User{}).Where("uid = ?", req.ID).UpdateColumn("passwd", req.Passwd).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "用户不存在",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "修改成功",
	})

}

func handleAccountUpdateInfo(c *gin.Context) {
	type Req struct {
		Idcard   string `json:"idcard"`
		Location string `json:"location"`
		Name     string `json:"name"`
		Sex      int    `json:"sex"`
		Tel      string `json:"tel"`
		UID      int    `json:"uid"`
		Avator   []struct {
			Name       string `json:"name"`
			Percentage int    `json:"percentage"`
			Status     string `json:"status"`
			Size       int    `json:"size"`
			Raw        struct {
				UID int64 `json:"uid"`
			} `json:"raw"`
			UID      int64  `json:"uid"`
			URL      string `json:"url"`
			Response struct {
				Code int    `json:"code"`
				Msg  string `json:"msg"`
				URL  string `json:"url"`
			} `json:"response"`
		} `json:"avator"`
	}
	req := Req{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	avatarURL := ""
	if req.Avator != nil {
		avatarURL = req.Avator[0].Response.URL
	}
	if err := db.GetConn().Model(&db.User{}).Where("uid = ?", req.UID).UpdateColumns(
		map[string]interface{}{
			"name":     req.Name,
			"sex":      req.Sex,
			"tel":      req.Tel,
			"idcard":   req.Idcard,
			"location": req.Location,
			"avator":   avatarURL,
		},
	).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "更新成功"})
}

func handleAccountDelete(c *gin.Context) {
	sid := c.Query("id")
	if err := db.GetConn().Where("uid = ?", sid).Delete(&db.User{}).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "用户不存在",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "删除成功",
	})
}
