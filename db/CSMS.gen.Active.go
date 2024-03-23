package db

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _ActiveMgr struct {
	*_BaseMgr
}

// ActiveMgr open func
func ActiveMgr(db *gorm.DB) *_ActiveMgr {
	if db == nil {
		panic(fmt.Errorf("ActiveMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ActiveMgr{_BaseMgr: &_BaseMgr{DB: db.Table("Active"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_ActiveMgr) Debug() *_ActiveMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ActiveMgr) GetTableName() string {
	return "Active"
}

// Reset 重置gorm会话
func (obj *_ActiveMgr) Reset() *_ActiveMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_ActiveMgr) Get() (result Active, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Active{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ActiveMgr) Gets() (results []*Active, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Active{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_ActiveMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Active{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithAcid acid获取 活动ID
func (obj *_ActiveMgr) WithAcid(acid int) Option {
	return optionFunc(func(o *options) { o.query["acid"] = acid })
}

// WithName name获取 活动名称
func (obj *_ActiveMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithStarttime startTime获取 开始时间
func (obj *_ActiveMgr) WithStarttime(starttime time.Time) Option {
	return optionFunc(func(o *options) { o.query["startTime"] = starttime })
}

// WithOpentime openTime获取 发布时间
func (obj *_ActiveMgr) WithOpentime(opentime time.Time) Option {
	return optionFunc(func(o *options) { o.query["openTime"] = opentime })
}

// WithEndtime endTime获取 结束时间
func (obj *_ActiveMgr) WithEndtime(endtime time.Time) Option {
	return optionFunc(func(o *options) { o.query["endTime"] = endtime })
}

// WithDetail detail获取 活动详情
func (obj *_ActiveMgr) WithDetail(detail string) Option {
	return optionFunc(func(o *options) { o.query["detail"] = detail })
}

// WithText text获取 活动正文
func (obj *_ActiveMgr) WithText(text string) Option {
	return optionFunc(func(o *options) { o.query["text"] = text })
}

// WithViews views获取 浏览量
func (obj *_ActiveMgr) WithViews(views uint32) Option {
	return optionFunc(func(o *options) { o.query["views"] = views })
}

// GetByOption 功能选项模式获取
func (obj *_ActiveMgr) GetByOption(opts ...Option) (result Active, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Active{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ActiveMgr) GetByOptions(opts ...Option) (results []*Active, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Active{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromAcid 通过acid获取内容 活动ID
func (obj *_ActiveMgr) GetFromAcid(acid int) (result Active, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Active{}).Where("`acid` = ?", acid).First(&result).Error

	return
}

// GetBatchFromAcid 批量查找 活动ID
func (obj *_ActiveMgr) GetBatchFromAcid(acids []int) (results []*Active, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Active{}).Where("`acid` IN (?)", acids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 活动名称
func (obj *_ActiveMgr) GetFromName(name string) (results []*Active, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Active{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 活动名称
func (obj *_ActiveMgr) GetBatchFromName(names []string) (results []*Active, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Active{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromStarttime 通过startTime获取内容 开始时间
func (obj *_ActiveMgr) GetFromStarttime(starttime time.Time) (results []*Active, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Active{}).Where("`startTime` = ?", starttime).Find(&results).Error

	return
}

// GetBatchFromStarttime 批量查找 开始时间
func (obj *_ActiveMgr) GetBatchFromStarttime(starttimes []time.Time) (results []*Active, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Active{}).Where("`startTime` IN (?)", starttimes).Find(&results).Error

	return
}

// GetFromOpentime 通过openTime获取内容 发布时间
func (obj *_ActiveMgr) GetFromOpentime(opentime time.Time) (results []*Active, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Active{}).Where("`openTime` = ?", opentime).Find(&results).Error

	return
}

// GetBatchFromOpentime 批量查找 发布时间
func (obj *_ActiveMgr) GetBatchFromOpentime(opentimes []time.Time) (results []*Active, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Active{}).Where("`openTime` IN (?)", opentimes).Find(&results).Error

	return
}

// GetFromEndtime 通过endTime获取内容 结束时间
func (obj *_ActiveMgr) GetFromEndtime(endtime time.Time) (results []*Active, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Active{}).Where("`endTime` = ?", endtime).Find(&results).Error

	return
}

// GetBatchFromEndtime 批量查找 结束时间
func (obj *_ActiveMgr) GetBatchFromEndtime(endtimes []time.Time) (results []*Active, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Active{}).Where("`endTime` IN (?)", endtimes).Find(&results).Error

	return
}

// GetFromDetail 通过detail获取内容 活动详情
func (obj *_ActiveMgr) GetFromDetail(detail string) (results []*Active, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Active{}).Where("`detail` = ?", detail).Find(&results).Error

	return
}

// GetBatchFromDetail 批量查找 活动详情
func (obj *_ActiveMgr) GetBatchFromDetail(details []string) (results []*Active, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Active{}).Where("`detail` IN (?)", details).Find(&results).Error

	return
}

// GetFromText 通过text获取内容 活动正文
func (obj *_ActiveMgr) GetFromText(text string) (results []*Active, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Active{}).Where("`text` = ?", text).Find(&results).Error

	return
}

// GetBatchFromText 批量查找 活动正文
func (obj *_ActiveMgr) GetBatchFromText(texts []string) (results []*Active, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Active{}).Where("`text` IN (?)", texts).Find(&results).Error

	return
}

// GetFromViews 通过views获取内容 浏览量
func (obj *_ActiveMgr) GetFromViews(views uint32) (results []*Active, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Active{}).Where("`views` = ?", views).Find(&results).Error

	return
}

// GetBatchFromViews 批量查找 浏览量
func (obj *_ActiveMgr) GetBatchFromViews(viewss []uint32) (results []*Active, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Active{}).Where("`views` IN (?)", viewss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_ActiveMgr) FetchByPrimaryKey(acid int) (result Active, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Active{}).Where("`acid` = ?", acid).First(&result).Error

	return
}
