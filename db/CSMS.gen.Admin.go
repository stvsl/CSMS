package db

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AdminMgr struct {
	*_BaseMgr
}

// AdminMgr open func
func AdminMgr(db *gorm.DB) *_AdminMgr {
	if db == nil {
		panic(fmt.Errorf("AdminMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AdminMgr{_BaseMgr: &_BaseMgr{DB: db.Table("Admin"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_AdminMgr) Debug() *_AdminMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AdminMgr) GetTableName() string {
	return "Admin"
}

// Reset 重置gorm会话
func (obj *_AdminMgr) Reset() *_AdminMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AdminMgr) Get() (result Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AdminMgr) Gets() (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AdminMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Admin{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_AdminMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取
func (obj *_AdminMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithTel tel获取
func (obj *_AdminMgr) WithTel(tel string) Option {
	return optionFunc(func(o *options) { o.query["tel"] = tel })
}

// WithPasswd passwd获取
func (obj *_AdminMgr) WithPasswd(passwd string) Option {
	return optionFunc(func(o *options) { o.query["passwd"] = passwd })
}

// GetByOption 功能选项模式获取
func (obj *_AdminMgr) GetByOption(opts ...Option) (result Admin, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AdminMgr) GetByOptions(opts ...Option) (results []*Admin, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_AdminMgr) GetFromID(id int) (result Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_AdminMgr) GetBatchFromID(ids []int) (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_AdminMgr) GetFromName(name string) (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_AdminMgr) GetBatchFromName(names []string) (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromTel 通过tel获取内容
func (obj *_AdminMgr) GetFromTel(tel string) (results Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`tel` = ?", tel).First(&results).Error

	return
}

// GetBatchFromTel 批量查找
func (obj *_AdminMgr) GetBatchFromTel(tels []string) (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`tel` IN (?)", tels).Find(&results).Error

	return
}

// GetFromPasswd 通过passwd获取内容
func (obj *_AdminMgr) GetFromPasswd(passwd string) (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`passwd` = ?", passwd).Find(&results).Error

	return
}

// GetBatchFromPasswd 批量查找
func (obj *_AdminMgr) GetBatchFromPasswd(passwds []string) (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`passwd` IN (?)", passwds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AdminMgr) FetchByPrimaryKey(id int) (result Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`id` = ?", id).First(&result).Error

	return
}
