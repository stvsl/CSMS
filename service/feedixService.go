package service

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"stvsljl.com/CSMS/db"
)

func handleFeedUpload(c *gin.Context) {
	type Req struct {
		Name   string `json:"name"`
		Detail string `json:"detail"`
		Id     uint32 `json:"id"`
	}
	var req Req
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}
	feed := db.Feed{
		Name:     req.Name,
		Detail:   req.Detail,
		UID:      req.Id,
		Feedtime: time.Now(),
		Process:  1,
	}
	if err := db.FeedMgr(db.GetConn()).Create(&feed).Error; err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "创建失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "创建成功",
	})
}

func handleFeedCount(c *gin.Context) {
	var count int64
	if err := db.FeedMgr(db.GetConn()).Where("type = ?", 0).Count(&count).Error; err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "查询失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"data": count,
	})
}

func handleFeedIdList(c *gin.Context) {
	spage := c.Query("page")
	page, err := strconv.Atoi(spage)
	if err != nil || page < 1 {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}
	var feedids []string
	if err := db.FeedMgr(db.GetConn()).Where("type = ?", 0).Select("id").Offset((page - 1) * 10).Limit(10).Find(&feedids).Error; err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "查询失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"data": feedids,
	})
}

func handleFeedDetails(c *gin.Context) {
	sid := c.Query("id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}
	feed, err := db.FeedMgr(db.GetConn()).GetFromID(id)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "查询失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"data": feed,
	})
}

func handleFeedDelete(c *gin.Context) {
	sid := c.Query("id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}
	if err := db.FeedMgr(db.GetConn()).Delete(id).Error; err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "删除失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "删除成功",
	})
}
func handleFeedUpdate(c *gin.Context) {

}

func handleFeedStatus(c *gin.Context) {
	sid := c.Query("id")
	sstatus := c.Query("status")
	id, err := strconv.Atoi(sid)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}
	status, err := strconv.Atoi(sstatus)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}
	if err := db.GetConn().Model(&db.Feed{}).Where("id = ?", id).Update("status", status).Error; err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "更新失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "更新成功",
	})
}

func handleFeedProcess(c *gin.Context) {
	sid := c.Query("id")
	sstatus := c.Query("status")
	id, err := strconv.Atoi(sid)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}
	status, err := strconv.Atoi(sstatus)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}
	if err := db.GetConn().Model(&db.Feed{}).Where("id = ?", id).Update("process", status).Error; err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "更新失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "更新成功",
	})
}

func handleFeedDo(c *gin.Context) {
	sid := c.Query("id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}
	type Req struct {
		ID   string `json:"id"`
		Mode int64  `json:"mode"`
		Name string `json:"name"`
	}
	var req Req
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}
	// 0: id 模式 1： name 模式
	if req.Mode == 0 {
		if err := db.GetConn().Model(&db.Feed{}).Where("id = ?", id).UpdateColumns(
			map[string]interface{}{
				"processor": nil,
				"oid":       req.ID,
				"process":   3,
			}).Error; err != nil {
			c.JSON(400, gin.H{
				"code": 400,
				"msg":  "更新失败",
			})
			return
		}
	} else {
		if err := db.GetConn().Model(&db.Feed{}).Where("id = ?", id).UpdateColumns(
			map[string]interface{}{
				"processor": req.Name,
				"oid":       0,
				"process":   4,
			}).Error; err != nil {
			c.JSON(400, gin.H{
				"code": 400,
				"msg":  "更新失败",
			})
			return
		}
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "更新成功",
	})
}

func handleFeedAdminOverview(c *gin.Context) {
	type Res struct {
		TotalFeedbacks int     `gorm:"column:total_feedbacks;type:int(11);not null;comment:'总反馈数'"`
		TotalRepairs   int     `gorm:"column:total_repairs;type:int(11);not null;comment:'总报修数'"`
		TotalProcessed int     `gorm:"column:total_processed;type:int(11);not null;comment:'已处理反馈数'"`
		TotalPending   int     `gorm:"column:total_pending;type:int(11);not null;comment:'待处理反馈数'"`
		InProgress     int     `gorm:"column:in_progress;type:int(11);not null;comment:'处理中反馈数'"`
		CompletionRate float64 `gorm:"column:completion_rate;type:float(11);not null;comment:'完成率'"`
	}
	var res Res
	if err := db.GetConn().Raw(`
		SELECT 
    		COUNT(*) AS total_feedbacks,
    		SUM(CASE WHEN type = 0 THEN 1 ELSE 0 END) AS total_repairs,
    		COUNT(CASE WHEN type = 0 AND (status = 0 AND process > 1 OR status = 1) THEN 1 END) AS total_processed,
    		(COUNT(*) - COUNT(CASE WHEN type = 0 AND (status = 0 AND process > 1 OR status = 1) THEN 1 END)) AS total_pending,
    		COUNT(CASE WHEN type = 0 AND process > 1 AND process != 4 THEN 1 END) AS in_progress,
    		(1 - ((COUNT(*) - COUNT(CASE WHEN type = 0 AND (status = 0 AND process > 1 OR status = 1) THEN 1 END)) / SUM(CASE WHEN type = 0 THEN 1 ELSE 0 END))) AS completion_rate
		FROM 
    		Feed
		WHERE
    		type = 0;`).Scan(&res).Error; err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "查询失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"data": res,
	})
}

func handleFeedUserOverview(c *gin.Context) {
	suid := c.Query("id")
	uid, err := strconv.Atoi(suid)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}
	type Res struct {
		TotalFeedbacks int     `gorm:"column:total_feedbacks;type:int(11);not null;comment:'总反馈数'"`
		TotalRepairs   int     `gorm:"column:total_repairs;type:int(11);not null;comment:'总报修数'"`
		TotalProcessed int     `gorm:"column:total_processed;type:int(11);not null;comment:'已处理反馈数'"`
		TotalPending   int     `gorm:"column:total_pending;type:int(11);not null;comment:'待处理反馈数'"`
		InProgress     int     `gorm:"column:in_progress;type:int(11);not null;comment:'处理中反馈数'"`
		CompletionRate float64 `gorm:"column:completion_rate;type:float(11);not null;comment:'完成率'"`
	}
	var res Res
	if err := db.GetConn().Raw(`
		SELECT 
    		COUNT(*) AS total_feedbacks,
    		SUM(CASE WHEN type = 0 THEN 1 ELSE 0 END) AS total_repairs,
    		COUNT(CASE WHEN type = 0 AND (status = 0 AND process > 1 OR status = 1) THEN 1 END) AS total_processed,
    		(COUNT(*) - COUNT(CASE WHEN type = 0 AND (status = 0 AND process > 1 OR status = 1) THEN 1 END)) AS total_pending,
    		COUNT(CASE WHEN type = 0 AND process > 1 AND process != 4 THEN 1 END) AS in_progress,
    		(1 - ((COUNT(*) - COUNT(CASE WHEN type = 0 AND (status = 0 AND process > 1 OR status = 1) THEN 1 END)) / SUM(CASE WHEN type = 0 THEN 1 ELSE 0 END))) AS completion_rate
		FROM 
    		Feed
		WHERE
    		type = 0 AND uid = ?;`, uid).Scan(&res).Error; err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "查询失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"data": res,
	})
}

func handleFeedUserCount(c *gin.Context) {
	sid := c.Query("id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}
	var count int64
	if err := db.FeedMgr(db.GetConn()).Where("type = ? AND uid = ?", 0, id).Count(&count).Error; err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "查询失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"data": count,
	})
}

func handleFeedUserIdList(c *gin.Context) {
	spage := c.Query("page")
	sid := c.Query("id")
	page, err := strconv.Atoi(spage)
	if err != nil || page < 1 {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}
	var feedids []string
	if err := db.FeedMgr(db.GetConn()).Where("type = ? AND uid = ?", 0, sid).Select("id").Offset((page - 1) * 5).Limit(5).Find(&feedids).Error; err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "查询失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"data": feedids,
	})
}

func handleFixUpload(c *gin.Context) {
	type Req struct {
		Name   string `json:"name"`
		Detail string `json:"detail"`
		Id     uint32 `json:"id"`
	}
	var req Req
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}
	feed := db.Feed{
		Name:     req.Name,
		Detail:   req.Detail,
		UID:      req.Id,
		Feedtime: time.Now(),
		Type:     1,
		Process:  1,
	}
	if err := db.FeedMgr(db.GetConn()).Create(&feed).Error; err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "创建失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "创建成功",
	})
}

func handleFixCount(c *gin.Context) {
	var count int64
	if err := db.FeedMgr(db.GetConn()).Where("type = ?", 1).Count(&count).Error; err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "查询失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"data": count,
	})
}

func handleFixIdList(c *gin.Context) {
	spage := c.Query("page")
	page, err := strconv.Atoi(spage)
	if err != nil || page < 1 {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}
	var fixids []string
	if err := db.FeedMgr(db.GetConn()).Where("type = ?", 1).Select("id").Offset((page - 1) * 10).Limit(10).Find(&fixids).Error; err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "查询失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"data": fixids,
	})
}

func handleFixDetails(c *gin.Context) {
	sid := c.Query("id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}
	feed, err := db.FeedMgr(db.GetConn()).GetFromID(id)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "查询失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"data": feed,
	})
}

func handleFixDelete(c *gin.Context) {
	sid := c.Query("id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}
	if err := db.FeedMgr(db.GetConn()).Delete(id).Error; err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "删除失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "删除成功",
	})
}

func handleFixUpdate(c *gin.Context) {

}

func handleFixStatus(c *gin.Context) {
	sid := c.Query("id")
	sstatus := c.Query("status")
	id, err := strconv.Atoi(sid)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}
	status, err := strconv.Atoi(sstatus)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}
	if err := db.GetConn().Model(&db.Feed{}).Where("id = ?", id).Update("status", status).Error; err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "更新失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "更新成功",
	})
}

func handleFixProcess(c *gin.Context) {
	sid := c.Query("id")
	sstatus := c.Query("status")
	id, err := strconv.Atoi(sid)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}
	status, err := strconv.Atoi(sstatus)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}
	if err := db.GetConn().Model(&db.Feed{}).Where("id = ?", id).Update("process", status).Error; err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "更新失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "更新成功",
	})
}

func handleFixDo(c *gin.Context) {
	sid := c.Query("id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}
	type Req struct {
		ID   int    `json:"id"`
		Mode int64  `json:"mode"`
		Name string `json:"name"`
	}
	var req Req
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}
	// 0: id 模式 1： name 模式
	if req.Mode == 0 {
		if err := db.GetConn().Model(&db.Feed{}).Where("id = ?", id).UpdateColumns(
			map[string]interface{}{
				"processor": nil,
				"oid":       req.ID,
				"process":   3,
			}).Error; err != nil {
			c.JSON(400, gin.H{
				"code": 400,
				"msg":  "更新失败",
			})
			return
		}
	} else {
		if err := db.GetConn().Model(&db.Feed{}).Where("id = ?", id).UpdateColumns(
			map[string]interface{}{
				"processor": req.Name,
				"oid":       0,
				"process":   4,
			}).Error; err != nil {
			c.JSON(400, gin.H{
				"code": 400,
				"msg":  "更新失败",
			})
			return
		}
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "更新成功",
	})
}

func handleFixAdminOverview(c *gin.Context) {
	type Res struct {
		TotalFeedbacks int     `gorm:"column:total_feedbacks;type:int(11);not null;comment:'总反馈数'"`
		TotalRepairs   int     `gorm:"column:total_repairs;type:int(11);not null;comment:'总报修数'"`
		TotalProcessed int     `gorm:"column:total_processed;type:int(11);not null;comment:'已处理反馈数'"`
		TotalPending   int     `gorm:"column:total_pending;type:int(11);not null;comment:'待处理反馈数'"`
		InProgress     int     `gorm:"column:in_progress;type:int(11);not null;comment:'处理中反馈数'"`
		CompletionRate float64 `gorm:"column:completion_rate;type:float(11);not null;comment:'完成率'"`
	}
	var res Res
	if err := db.GetConn().Raw(`
		SELECT 
    		COUNT(*) AS total_feedbacks,
    		SUM(CASE WHEN type = 1 THEN 1 ELSE 0 END) AS total_repairs,
    		COUNT(CASE WHEN type = 1 AND (status = 0 AND process > 1 OR status = 1) THEN 1 END) AS total_processed,
    		(COUNT(*) - COUNT(CASE WHEN type = 1 AND (status = 0 AND process > 1 OR status = 1) THEN 1 END)) AS total_pending,
    		COUNT(CASE WHEN type = 1 AND process > 1 AND process != 4 THEN 1 END) AS in_progress,
    		(1 - ((COUNT(*) - COUNT(CASE WHEN type = 1 AND (status = 0 AND process > 1 OR status = 1) THEN 1 END)) / SUM(CASE WHEN type = 1 THEN 1 ELSE 0 END))) AS completion_rate
		FROM 
    		Feed
		WHERE
    		type = 1;`).Scan(&res).Error; err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "查询失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"data": res,
	})
}

func handleFixUserOverview(c *gin.Context) {
	suid := c.Query("id")
	uid, err := strconv.Atoi(suid)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}
	type Res struct {
		TotalFeedbacks int     `gorm:"column:total_feedbacks;type:int(11);not null;comment:'总反馈数'"`
		TotalRepairs   int     `gorm:"column:total_repairs;type:int(11);not null;comment:'总报修数'"`
		TotalProcessed int     `gorm:"column:total_processed;type:int(11);not null;comment:'已处理反馈数'"`
		TotalPending   int     `gorm:"column:total_pending;type:int(11);not null;comment:'待处理反馈数'"`
		InProgress     int     `gorm:"column:in_progress;type:int(11);not null;comment:'处理中反馈数'"`
		CompletionRate float64 `gorm:"column:completion_rate;type:float(11);not null;comment:'完成率'"`
	}
	var res Res
	if err := db.GetConn().Raw(`
		SELECT 
    		COUNT(*) AS total_feedbacks,
    		SUM(CASE WHEN type = 1 THEN 1 ELSE 0 END) AS total_repairs,
    		COUNT(CASE WHEN type = 1 AND (status = 0 AND process > 1 OR status = 1) THEN 1 END) AS total_processed,
    		(COUNT(*) - COUNT(CASE WHEN type = 1 AND (status = 0 AND process > 1 OR status = 1) THEN 1 END)) AS total_pending,
    		COUNT(CASE WHEN type = 1 AND process > 1 AND process != 4 THEN 1 END) AS in_progress,
    		(1 - ((COUNT(*) - COUNT(CASE WHEN type = 1 AND (status = 0 AND process > 1 OR status = 1) THEN 1 END)) / SUM(CASE WHEN type = 1 THEN 1 ELSE 0 END))) AS completion_rate
		FROM 
    		Feed
		WHERE
    		type = 1 AND uid = ?;`, uid).Scan(&res).Error; err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "查询失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"data": res,
	})
}

func handleFixUserCount(c *gin.Context) {
	sid := c.Query("id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}
	var count int64
	if err := db.FeedMgr(db.GetConn()).Where("type = ? AND uid = ?", 1, id).Count(&count).Error; err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "查询失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"data": count,
	})
}

func handleFixUserIdList(c *gin.Context) {
	spage := c.Query("page")
	sid := c.Query("id")
	page, err := strconv.Atoi(spage)
	if err != nil || page < 1 {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}
	var fixids []string
	if err := db.FeedMgr(db.GetConn()).Where("type = ? AND uid = ?", 1, sid).Select("id").Offset((page - 1) * 5).Limit(5).Find(&fixids).Error; err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "查询失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"data": fixids,
	})
}
