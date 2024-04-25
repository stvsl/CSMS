package service

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"stvsljl.com/CSMS/db"
)

func handleActivityUpload(c *gin.Context) {
	type Req struct {
		Data struct {
			Name      string `json:"name"`
			StartDate string `json:"startDate"`
			StartTime string `json:"startTime"`
			StopDate  string `json:"stopDate"`
			StopTime  string `json:"stopTime"`
			Maxcount  int    `json:"maxcount"`
			Position  string `json:"position"`
			Detail    string `json:"detail"`
		} `json:"data"`
		Content string `json:"content"`
	}
	var req Req
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}
	starttime, _ := time.Parse("2006-01-02 15:04:05", req.Data.StartDate+" "+req.Data.StartTime)
	endtime, _ := time.Parse("2006-01-02 15:04:05", req.Data.StopDate+" "+req.Data.StopTime)
	activity := db.Active{
		Name:      req.Data.Name,
		Starttime: starttime,
		Endtime:   endtime,
		Detail:    req.Data.Detail,
		Views:     0,
		Opentime:  time.Now(),
		Text:      req.Content,
		Maxcount:  req.Data.Maxcount,
		Position:  req.Data.Position,
	}
	mgr := db.ActiveMgr(db.GetConn())
	// 插入数据库
	if err := mgr.Create(&activity).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "服务器错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "上传成功",
	})
}

func handleActivityCount(c *gin.Context) {
	var count int64
	if err := db.GetConn().Model(&db.Active{}).Count(&count).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "服务器错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "查询成功",
		"data":    count,
	})
}

func handleActivityIdList(c *gin.Context) {
	spage := c.Query("page")
	page, err := strconv.Atoi(spage)
	if page == 0 || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "请求参数错误" + err.Error(),
		})
		return
	}
	type Res struct {
		Acid int `gorm:"autoIncrement:true;primaryKey;column:acid;type:int(11);not null;comment:'活动ID'"` // 活动ID
	}
	var res []Res
	if err := db.GetConn().Model(&db.Active{}).Select("acid").Offset((page - 1) * 10).Limit(10).Find(&res).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "服务器错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "查询成功",
		"data":    res,
	})
}

func handleActivityDetails(c *gin.Context) {
	sid := c.Query("id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "请求参数错误",
		})
		return
	}
	active := db.Active{}
	if err := db.GetConn().Model(&db.Active{}).Where("acid = ?", id).First(&active).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "服务器错误",
		})
		return
	}
	db.GetConn().Model(&db.Active{}).Where("acid = ?", id).Update("views", active.Views+1)
	// 查找Activityparticipation表中acid相同的总数
	var count int64
	if err := db.GetConn().Model(&db.Activityparticipation{}).Where("acid = ?", id).Count(&count).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "服务器错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "查询成功",
		"data":    active,
		"count":   count,
	})
}

func handleActivityStatus(c *gin.Context) {
	acid := c.Query("id")
	type ActivityParticipation struct {
		UID    int    `gorm:"autoIncrement:true;primaryKey;column:uid;type:int(11);not null;comment:'用户ID'"` // 用户ID
		Name   string `gorm:"column:name;type:varchar(50);not null;comment:'用户名'"`                           // 用户名
		Tel    string `gorm:"unique;column:tel;type:char(11);not null;comment:'联系电话'"`                       // 联系电话
		Status uint32 `gorm:"column:status;type:int(1) unsigned;not null;comment:'参与状态'"`                    // 参与状态
		Avator string `gorm:"column:avator;type:varchar(100);not null;comment:'头像'"`                         // 头像
	}
	var users []ActivityParticipation
	if err := db.GetConn().Table("ActivityParticipation").Select("ActivityParticipation.uid,ActivityParticipation.status,User.name,User.tel,User.avator").Where("acid = ?", acid).Joins("left join User on User.uid = ActivityParticipation.uid").Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "服务器错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "查询成功",
		"data":    users,
	})
}

func handleActivityDelete(c *gin.Context) {
	sacid := c.Query("id")
	acid, err := strconv.Atoi(sacid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "请求参数错误",
		})
		return
	}
	mgr := db.ActiveMgr(db.GetConn())
	if err := mgr.Delete(&db.Active{Acid: acid}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "服务器错误",
		})
		return
	}
	if err := db.GetConn().Where("acid = ?", sacid).Delete(&db.Activityparticipation{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "在子事务处理阶段服务器错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "删除成功",
	})
}

func handleActivityPeopleStatus(c *gin.Context) {
	type Req struct {
		ID     int `json:"id"`
		Status int `json:"status"`
	}
	var req Req
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "请求参数错误",
		})
		return
	}
	if err := db.GetConn().Model(&db.Activityparticipation{}).Where("uid = ?", req.ID).Update("status", req.Status).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "服务器错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "更新成功",
	})
}

func handleActivityUpdate(c *gin.Context) {
	type Req struct {
		Data struct {
			Name      string `json:"name"`
			StartDate string `json:"startDate"`
			StartTime string `json:"startTime"`
			StopDate  string `json:"stopDate"`
			StopTime  string `json:"stopTime"`
			Maxcount  int    `json:"maxcount"`
			Position  string `json:"position"`
			Detail    string `json:"detail"`
		} `json:"data"`
		Content string `json:"content"`
		Acid    int    `json:"acid"`
	}
	var req Req
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	starttime, _ := time.Parse("2006-01-02 15:04:05", req.Data.StartDate+" "+req.Data.StartTime)
	endtime, _ := time.Parse("2006-01-02 15:04:05", req.Data.StopDate+" "+req.Data.StopTime)
	active := db.Active{
		Acid:      req.Acid,
		Name:      req.Data.Name,
		Starttime: starttime,
		Endtime:   endtime,
		Detail:    req.Data.Detail,
		Text:      req.Content,
		Maxcount:  req.Data.Maxcount,
		Position:  req.Data.Position,
	}
	if err := db.GetConn().Model(&db.Active{}).Where("acid = ?", req.Acid).UpdateColumns(
		map[string]interface{}{
			"name":      active.Name,
			"starttime": active.Starttime,
			"endtime":   active.Endtime,
			"detail":    active.Detail,
			"text":      active.Text,
			"maxcount":  active.Maxcount,
			"position":  active.Position,
		},
	).Error; err != nil {
		c.JSON(http.StatusNotModified, gin.H{
			"code": 500,
			"msg":  "更新失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "更新成功",
		})
	}
}

func handleActivityRecentIdList(c *gin.Context) {
	// 查找endTime大于当前时间的活动
	var ids []int
	if err := db.GetConn().Model(&db.Active{}).Where("endTime > ?", time.Now()).Pluck("acid", &ids).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "服务器错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "查询成功",
		"data":    ids,
	})

}

func handleActivityHasjion(c *gin.Context) {
	sid := c.Query("id")
	sacid := c.Query("acid")
	if sid == "" || sacid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "参数错误",
		})
		return
	}
	var count int64
	db.GetConn().Model(&db.Activityparticipation{}).Where("uid = ? and acid = ?", sid, sacid).Count(&count)
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": count,
	})
}

func handleActivityUserStatus(c *gin.Context) {
	sid := c.Query("id")
	sacid := c.Query("acid")
	if sid == "" || sacid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "参数错误",
		})
		return
	}
	var status int
	if err := db.GetConn().Model(&db.Activityparticipation{}).Where("uid = ? and acid = ?", sid, sacid).Pluck("status", &status).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "服务器错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "查询成功",
		"data":    status,
	})
}

func handleActivityJoin(c *gin.Context) {
	sid := c.Query("id")
	sacid := c.Query("acid")
	id, _ := strconv.Atoi(sid)
	ap := db.Activityparticipation{
		UID:    uint32(id),
		Acid:   sacid,
		Status: 0,
	}
	if err := db.GetConn().Create(&ap).Error; err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"code":    500,
			"message": "服务器错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "报名成功",
	})
}

func handleActivityExit(c *gin.Context) {
	sid := c.Query("id")
	sacid := c.Query("acid")
	if err := db.GetConn().Where("acid = ?", sacid).Where("uid = ?", sid).Delete(&db.Activityparticipation{}).Error; err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"code":    500,
			"message": "服务器错误",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "退出成功",
	})
}

func handleActivityUserCount(c *gin.Context) {
	sid := c.Query("id")
	var count int64
	if err := db.GetConn().Where("uid = ?", sid).Find(&db.Activityparticipation{}).Count(&count).Error; err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"code":    500,
			"message": "服务器错误",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "查询成功",
		"data":    count,
	})
}

func handleActivityUserIdList(c *gin.Context) {
	sid := c.Query("id")
	spage := c.Query("page")
	if spage == "" {
		spage = "1"
	}
	page, _ := strconv.Atoi(spage)
	if sid == "" {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"code":    500,
			"message": "参数错误",
		})
		return
	}
	var ids []string
	if err := db.GetConn().Model(&db.Activityparticipation{}).Offset((page-1)*10).Where("uid = ?", sid).Pluck("acid", &ids).Error; err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"code":    500,
			"message": "数据库错误",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    ids,
	})
}
