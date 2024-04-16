package db

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AnounceMgr struct {
	*_BaseMgr
}

// AnounceMgr open func
func AnounceMgr(db *gorm.DB) *_AnounceMgr {
	if db == nil {
		panic(fmt.Errorf("AnounceMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AnounceMgr{_BaseMgr: &_BaseMgr{DB: db.Table("Anounce"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_AnounceMgr) Debug() *_AnounceMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AnounceMgr) GetTableName() string {
	return "Anounce"
}

// Reset 重置gorm会话
func (obj *_AnounceMgr) Reset() *_AnounceMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AnounceMgr) Get() (result Anounce, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Anounce{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AnounceMgr) Gets() (results []*Anounce, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Anounce{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AnounceMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Anounce{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithAid aid获取 公告编号
func (obj *_AnounceMgr) WithAid(aid int) Option {
	return optionFunc(func(o *options) { o.query["aid"] = aid })
}

// WithTitle title获取 标题
func (obj *_AnounceMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithIntroduction introduction获取 简介
func (obj *_AnounceMgr) WithIntroduction(introduction string) Option {
	return optionFunc(func(o *options) { o.query["introduction"] = introduction })
}

// WithText text获取 正文
func (obj *_AnounceMgr) WithText(text string) Option {
	return optionFunc(func(o *options) { o.query["text"] = text })
}

// WithWritetime writetime获取 发表日期
func (obj *_AnounceMgr) WithWritetime(writetime time.Time) Option {
	return optionFunc(func(o *options) { o.query["writetime"] = writetime })
}

// WithUpdatetime updatetime获取 更新日期
func (obj *_AnounceMgr) WithUpdatetime(updatetime time.Time) Option {
	return optionFunc(func(o *options) { o.query["updatetime"] = updatetime })
}

// WithAuthor author获取 作者
func (obj *_AnounceMgr) WithAuthor(author string) Option {
	return optionFunc(func(o *options) { o.query["author"] = author })
}

// WithPageviews pageviews获取 浏览量
func (obj *_AnounceMgr) WithPageviews(pageviews uint64) Option {
	return optionFunc(func(o *options) { o.query["pageviews"] = pageviews })
}

// WithStatus status获取 状态
func (obj *_AnounceMgr) WithStatus(status int) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// GetByOption 功能选项模式获取
func (obj *_AnounceMgr) GetByOption(opts ...Option) (result Anounce, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Anounce{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AnounceMgr) GetByOptions(opts ...Option) (results []*Anounce, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Anounce{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromAid 通过aid获取内容 公告编号
func (obj *_AnounceMgr) GetFromAid(aid int) (result Anounce, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Anounce{}).Where("`aid` = ?", aid).First(&result).Error

	return
}

// GetBatchFromAid 批量查找 公告编号
func (obj *_AnounceMgr) GetBatchFromAid(aids []int) (results []*Anounce, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Anounce{}).Where("`aid` IN (?)", aids).Find(&results).Error

	return
}

// GetFromTitle 通过title获取内容 标题
func (obj *_AnounceMgr) GetFromTitle(title string) (results []*Anounce, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Anounce{}).Where("`title` = ?", title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量查找 标题
func (obj *_AnounceMgr) GetBatchFromTitle(titles []string) (results []*Anounce, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Anounce{}).Where("`title` IN (?)", titles).Find(&results).Error

	return
}

// GetFromIntroduction 通过introduction获取内容 简介
func (obj *_AnounceMgr) GetFromIntroduction(introduction string) (results []*Anounce, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Anounce{}).Where("`introduction` = ?", introduction).Find(&results).Error

	return
}

// GetBatchFromIntroduction 批量查找 简介
func (obj *_AnounceMgr) GetBatchFromIntroduction(introductions []string) (results []*Anounce, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Anounce{}).Where("`introduction` IN (?)", introductions).Find(&results).Error

	return
}

// GetFromText 通过text获取内容 正文
func (obj *_AnounceMgr) GetFromText(text string) (results []*Anounce, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Anounce{}).Where("`text` = ?", text).Find(&results).Error

	return
}

// GetBatchFromText 批量查找 正文
func (obj *_AnounceMgr) GetBatchFromText(texts []string) (results []*Anounce, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Anounce{}).Where("`text` IN (?)", texts).Find(&results).Error

	return
}

// GetFromWritetime 通过writetime获取内容 发表日期
func (obj *_AnounceMgr) GetFromWritetime(writetime time.Time) (results []*Anounce, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Anounce{}).Where("`writetime` = ?", writetime).Find(&results).Error

	return
}

// GetBatchFromWritetime 批量查找 发表日期
func (obj *_AnounceMgr) GetBatchFromWritetime(writetimes []time.Time) (results []*Anounce, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Anounce{}).Where("`writetime` IN (?)", writetimes).Find(&results).Error

	return
}

// GetFromUpdatetime 通过updatetime获取内容 更新日期
func (obj *_AnounceMgr) GetFromUpdatetime(updatetime time.Time) (results []*Anounce, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Anounce{}).Where("`updatetime` = ?", updatetime).Find(&results).Error

	return
}

// GetBatchFromUpdatetime 批量查找 更新日期
func (obj *_AnounceMgr) GetBatchFromUpdatetime(updatetimes []time.Time) (results []*Anounce, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Anounce{}).Where("`updatetime` IN (?)", updatetimes).Find(&results).Error

	return
}

// GetFromAuthor 通过author获取内容 作者
func (obj *_AnounceMgr) GetFromAuthor(author string) (results []*Anounce, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Anounce{}).Where("`author` = ?", author).Find(&results).Error

	return
}

// GetBatchFromAuthor 批量查找 作者
func (obj *_AnounceMgr) GetBatchFromAuthor(authors []string) (results []*Anounce, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Anounce{}).Where("`author` IN (?)", authors).Find(&results).Error

	return
}

// GetFromPageviews 通过pageviews获取内容 浏览量
func (obj *_AnounceMgr) GetFromPageviews(pageviews uint64) (results []*Anounce, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Anounce{}).Where("`pageviews` = ?", pageviews).Find(&results).Error

	return
}

// GetBatchFromPageviews 批量查找 浏览量
func (obj *_AnounceMgr) GetBatchFromPageviews(pageviewss []uint64) (results []*Anounce, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Anounce{}).Where("`pageviews` IN (?)", pageviewss).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态
func (obj *_AnounceMgr) GetFromStatus(status int) (results []*Anounce, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Anounce{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 状态
func (obj *_AnounceMgr) GetBatchFromStatus(statuss []int) (results []*Anounce, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Anounce{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AnounceMgr) FetchByPrimaryKey(aid int) (result Anounce, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Anounce{}).Where("`aid` = ?", aid).First(&result).Error

	return
}
