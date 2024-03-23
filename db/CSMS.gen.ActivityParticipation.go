package db

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ActivityparticipationMgr struct {
	*_BaseMgr
}

// ActivityparticipationMgr open func
func ActivityparticipationMgr(db *gorm.DB) *_ActivityparticipationMgr {
	if db == nil {
		panic(fmt.Errorf("ActivityparticipationMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ActivityparticipationMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ActivityParticipation"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_ActivityparticipationMgr) Debug() *_ActivityparticipationMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ActivityparticipationMgr) GetTableName() string {
	return "ActivityParticipation"
}

// Reset 重置gorm会话
func (obj *_ActivityparticipationMgr) Reset() *_ActivityparticipationMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_ActivityparticipationMgr) Get() (result Activityparticipation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Activityparticipation{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ActivityparticipationMgr) Gets() (results []*Activityparticipation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Activityparticipation{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_ActivityparticipationMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Activityparticipation{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithUID uid获取 参与用户ID
func (obj *_ActivityparticipationMgr) WithUID(uid uint32) Option {
	return optionFunc(func(o *options) { o.query["uid"] = uid })
}

// WithAcid acid获取 活动ID
func (obj *_ActivityparticipationMgr) WithAcid(acid string) Option {
	return optionFunc(func(o *options) { o.query["acid"] = acid })
}

// WithStatus status获取 参与状态
func (obj *_ActivityparticipationMgr) WithStatus(status uint32) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// GetByOption 功能选项模式获取
func (obj *_ActivityparticipationMgr) GetByOption(opts ...Option) (result Activityparticipation, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Activityparticipation{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ActivityparticipationMgr) GetByOptions(opts ...Option) (results []*Activityparticipation, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Activityparticipation{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromUID 通过uid获取内容 参与用户ID
func (obj *_ActivityparticipationMgr) GetFromUID(uid uint32) (results []*Activityparticipation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Activityparticipation{}).Where("`uid` = ?", uid).Find(&results).Error

	return
}

// GetBatchFromUID 批量查找 参与用户ID
func (obj *_ActivityparticipationMgr) GetBatchFromUID(uids []uint32) (results []*Activityparticipation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Activityparticipation{}).Where("`uid` IN (?)", uids).Find(&results).Error

	return
}

// GetFromAcid 通过acid获取内容 活动ID
func (obj *_ActivityparticipationMgr) GetFromAcid(acid string) (results []*Activityparticipation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Activityparticipation{}).Where("`acid` = ?", acid).Find(&results).Error

	return
}

// GetBatchFromAcid 批量查找 活动ID
func (obj *_ActivityparticipationMgr) GetBatchFromAcid(acids []string) (results []*Activityparticipation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Activityparticipation{}).Where("`acid` IN (?)", acids).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 参与状态
func (obj *_ActivityparticipationMgr) GetFromStatus(status uint32) (results []*Activityparticipation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Activityparticipation{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 参与状态
func (obj *_ActivityparticipationMgr) GetBatchFromStatus(statuss []uint32) (results []*Activityparticipation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Activityparticipation{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
