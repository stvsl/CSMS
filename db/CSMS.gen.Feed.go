package db

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _FeedMgr struct {
	*_BaseMgr
}

// FeedMgr open func
func FeedMgr(db *gorm.DB) *_FeedMgr {
	if db == nil {
		panic(fmt.Errorf("FeedMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_FeedMgr{_BaseMgr: &_BaseMgr{DB: db.Table("Feed"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_FeedMgr) Debug() *_FeedMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_FeedMgr) GetTableName() string {
	return "Feed"
}

// Reset 重置gorm会话
func (obj *_FeedMgr) Reset() *_FeedMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_FeedMgr) Get() (result Feed, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Feed{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_FeedMgr) Gets() (results []*Feed, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Feed{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_FeedMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Feed{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 ID
func (obj *_FeedMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUID uid获取 反馈用户
func (obj *_FeedMgr) WithUID(uid uint32) Option {
	return optionFunc(func(o *options) { o.query["uid"] = uid })
}

// WithType type获取 类型（0反馈1报修）
func (obj *_FeedMgr) WithType(_type uint32) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithName name获取 反馈名称
func (obj *_FeedMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithFeedtime feedtime获取 反馈时间
func (obj *_FeedMgr) WithFeedtime(feedtime string) Option {
	return optionFunc(func(o *options) { o.query["feedtime"] = feedtime })
}

// WithDetail detail获取 反馈内容
func (obj *_FeedMgr) WithDetail(detail string) Option {
	return optionFunc(func(o *options) { o.query["detail"] = detail })
}

// WithProcess process获取 反馈进度
func (obj *_FeedMgr) WithProcess(process uint32) Option {
	return optionFunc(func(o *options) { o.query["process"] = process })
}

// WithStatus status获取 反馈状态
func (obj *_FeedMgr) WithStatus(status int) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithOid oid获取 委派人
func (obj *_FeedMgr) WithOid(oid uint32) Option {
	return optionFunc(func(o *options) { o.query["oid"] = oid })
}

// WithProcessor processor获取 处理人
func (obj *_FeedMgr) WithProcessor(processor uint32) Option {
	return optionFunc(func(o *options) { o.query["processor"] = processor })
}

// WithUpdatetime updateTime获取 更新时间
func (obj *_FeedMgr) WithUpdatetime(updatetime time.Time) Option {
	return optionFunc(func(o *options) { o.query["updateTime"] = updatetime })
}

// WithRecord record获取 数据记录
func (obj *_FeedMgr) WithRecord(record string) Option {
	return optionFunc(func(o *options) { o.query["record"] = record })
}

// GetByOption 功能选项模式获取
func (obj *_FeedMgr) GetByOption(opts ...Option) (result Feed, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Feed{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_FeedMgr) GetByOptions(opts ...Option) (results []*Feed, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Feed{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 ID
func (obj *_FeedMgr) GetFromID(id int) (result Feed, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Feed{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 ID
func (obj *_FeedMgr) GetBatchFromID(ids []int) (results []*Feed, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Feed{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromUID 通过uid获取内容 反馈用户
func (obj *_FeedMgr) GetFromUID(uid uint32) (results []*Feed, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Feed{}).Where("`uid` = ?", uid).Find(&results).Error

	return
}

// GetBatchFromUID 批量查找 反馈用户
func (obj *_FeedMgr) GetBatchFromUID(uids []uint32) (results []*Feed, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Feed{}).Where("`uid` IN (?)", uids).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 类型（0反馈1报修）
func (obj *_FeedMgr) GetFromType(_type uint32) (results []*Feed, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Feed{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找 类型（0反馈1报修）
func (obj *_FeedMgr) GetBatchFromType(_types []uint32) (results []*Feed, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Feed{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 反馈名称
func (obj *_FeedMgr) GetFromName(name string) (results []*Feed, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Feed{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 反馈名称
func (obj *_FeedMgr) GetBatchFromName(names []string) (results []*Feed, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Feed{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromFeedtime 通过feedtime获取内容 反馈时间
func (obj *_FeedMgr) GetFromFeedtime(feedtime string) (results []*Feed, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Feed{}).Where("`feedtime` = ?", feedtime).Find(&results).Error

	return
}

// GetBatchFromFeedtime 批量查找 反馈时间
func (obj *_FeedMgr) GetBatchFromFeedtime(feedtimes []string) (results []*Feed, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Feed{}).Where("`feedtime` IN (?)", feedtimes).Find(&results).Error

	return
}

// GetFromDetail 通过detail获取内容 反馈内容
func (obj *_FeedMgr) GetFromDetail(detail string) (results []*Feed, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Feed{}).Where("`detail` = ?", detail).Find(&results).Error

	return
}

// GetBatchFromDetail 批量查找 反馈内容
func (obj *_FeedMgr) GetBatchFromDetail(details []string) (results []*Feed, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Feed{}).Where("`detail` IN (?)", details).Find(&results).Error

	return
}

// GetFromProcess 通过process获取内容 反馈进度
func (obj *_FeedMgr) GetFromProcess(process uint32) (results []*Feed, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Feed{}).Where("`process` = ?", process).Find(&results).Error

	return
}

// GetBatchFromProcess 批量查找 反馈进度
func (obj *_FeedMgr) GetBatchFromProcess(processs []uint32) (results []*Feed, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Feed{}).Where("`process` IN (?)", processs).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 反馈状态
func (obj *_FeedMgr) GetFromStatus(status int) (results []*Feed, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Feed{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 反馈状态
func (obj *_FeedMgr) GetBatchFromStatus(statuss []int) (results []*Feed, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Feed{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromOid 通过oid获取内容 委派人
func (obj *_FeedMgr) GetFromOid(oid uint32) (results []*Feed, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Feed{}).Where("`oid` = ?", oid).Find(&results).Error

	return
}

// GetBatchFromOid 批量查找 委派人
func (obj *_FeedMgr) GetBatchFromOid(oids []uint32) (results []*Feed, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Feed{}).Where("`oid` IN (?)", oids).Find(&results).Error

	return
}

// GetFromProcessor 通过processor获取内容 处理人
func (obj *_FeedMgr) GetFromProcessor(processor uint32) (results []*Feed, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Feed{}).Where("`processor` = ?", processor).Find(&results).Error

	return
}

// GetBatchFromProcessor 批量查找 处理人
func (obj *_FeedMgr) GetBatchFromProcessor(processors []uint32) (results []*Feed, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Feed{}).Where("`processor` IN (?)", processors).Find(&results).Error

	return
}

// GetFromUpdatetime 通过updateTime获取内容 更新时间
func (obj *_FeedMgr) GetFromUpdatetime(updatetime time.Time) (results []*Feed, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Feed{}).Where("`updateTime` = ?", updatetime).Find(&results).Error

	return
}

// GetBatchFromUpdatetime 批量查找 更新时间
func (obj *_FeedMgr) GetBatchFromUpdatetime(updatetimes []time.Time) (results []*Feed, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Feed{}).Where("`updateTime` IN (?)", updatetimes).Find(&results).Error

	return
}

// GetFromRecord 通过record获取内容 数据记录
func (obj *_FeedMgr) GetFromRecord(record string) (results []*Feed, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Feed{}).Where("`record` = ?", record).Find(&results).Error

	return
}

// GetBatchFromRecord 批量查找 数据记录
func (obj *_FeedMgr) GetBatchFromRecord(records []string) (results []*Feed, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Feed{}).Where("`record` IN (?)", records).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_FeedMgr) FetchByPrimaryKey(id int) (result Feed, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Feed{}).Where("`id` = ?", id).First(&result).Error

	return
}
