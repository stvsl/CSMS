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
	if len(req.Avator) > 0 {
		avatarURL = req.Avator[0].Response.URL
		if avatarURL == "" {
			avatarURL = req.Avator[0].URL
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
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "error": "数据上传格式错误"})
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

func handleAccountIdsGetByKey(c *gin.Context) {
	key := c.Query("key")
	value := c.Query("value")
	if key == "" || value == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}
	if key == "名称" {
		key = "name"
	}
	if key == "联系电话" {
		key = "tel"
	}
	if key == "UID" {
		key = "uid"
	}
	var ids []int
	if err := db.GetConn().Model(&db.User{}).Where(key+" like ?", "%"+value+"%").Pluck("uid", &ids).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "服务器错误" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": ids,
	})
}

func handleAccountOverview(c *gin.Context) {
	type Res struct {
		HasRegisterd   int64 `json:"hasRegisterd"`
		OtherUserCount int64 `json:"otherUserCount"`
	}
	var res Res
	db.GetConn().Table("User").Count(&res.HasRegisterd)
	db.GetConn().Table("OtherUser").Count(&res.OtherUserCount)
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "请求成功",
		"data": res,
	})
}

func handleAccountAdd(c *gin.Context) {
	type Req struct {
		Name   string `json:"name"`
		Tel    string `json:"tel"`
		Idcard string `json:"idcard"`
		Passwd string `json:"passwd"`
	}
	req := Req{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "error": err.Error()})
		return
	}
	if req.Name == "" || req.Tel == "" || req.Idcard == "" || req.Passwd == "" {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "请求参数错误"})
		return
	}
	user := db.User{
		Name:   req.Name,
		Tel:    req.Tel,
		IDcard: req.Idcard,
		Passwd: req.Passwd,
	}
	if err := db.GetConn().Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "服务器错误"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "添加成功",
	})
}
func handleUserSearch(c *gin.Context) {

	key := c.Query("key")
	if key == "" {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "参数错误"})
		return
	}
	var arctileids []int
	if err := db.GetConn().Table("Article").Select("aid").Where("title LIKE ? OR introduction LIKE ? OR text LIKE ?", "%"+key+"%", "%"+key+"%", "%"+key+"%").Pluck("aid", &arctileids).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "搜索失败",
		})
		return
	}
	var anounceids []int
	if err := db.GetConn().Table("Anounce").Select("aid").Where("title LIKE ? OR introduction LIKE ? OR text LIKE ?", "%"+key+"%", "%"+key+"%", "%"+key+"%").Pluck("aid", &anounceids).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "搜索失败",
		})
		return
	}
	var activeids []int
	if err := db.GetConn().Table("Active").Select("acid").Where("name LIKE ? OR text LIKE ?", "%"+key+"%", "%"+key+"%").Pluck("acid", &activeids).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "搜索失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "搜索成功",
		"data": gin.H{
			"article": arctileids,
			"anounce": anounceids,
			"active":  activeids,
		},
	})
}
