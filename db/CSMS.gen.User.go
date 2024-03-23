package db

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _UserMgr struct {
	*_BaseMgr
}

// UserMgr open func
func UserMgr(db *gorm.DB) *_UserMgr {
	if db == nil {
		panic(fmt.Errorf("UserMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_UserMgr{_BaseMgr: &_BaseMgr{DB: db.Table("User"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_UserMgr) Debug() *_UserMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UserMgr) GetTableName() string {
	return "User"
}

// Reset 重置gorm会话
func (obj *_UserMgr) Reset() *_UserMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_UserMgr) Get() (result User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_UserMgr) Gets() (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_UserMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(User{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithUID uid获取 用户ID
func (obj *_UserMgr) WithUID(uid int) Option {
	return optionFunc(func(o *options) { o.query["uid"] = uid })
}

// WithName name获取 用户名
func (obj *_UserMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithPasswd passwd获取 用户密码
func (obj *_UserMgr) WithPasswd(passwd string) Option {
	return optionFunc(func(o *options) { o.query["passwd"] = passwd })
}

// WithTel tel获取 联系电话
func (obj *_UserMgr) WithTel(tel int) Option {
	return optionFunc(func(o *options) { o.query["tel"] = tel })
}

// WithAvator avator获取 头像
func (obj *_UserMgr) WithAvator(avator string) Option {
	return optionFunc(func(o *options) { o.query["avator"] = avator })
}

// GetByOption 功能选项模式获取
func (obj *_UserMgr) GetByOption(opts ...Option) (result User, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_UserMgr) GetByOptions(opts ...Option) (results []*User, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromUID 通过uid获取内容 用户ID
func (obj *_UserMgr) GetFromUID(uid int) (result User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`uid` = ?", uid).First(&result).Error

	return
}

// GetBatchFromUID 批量查找 用户ID
func (obj *_UserMgr) GetBatchFromUID(uids []int) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`uid` IN (?)", uids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 用户名
func (obj *_UserMgr) GetFromName(name string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 用户名
func (obj *_UserMgr) GetBatchFromName(names []string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromPasswd 通过passwd获取内容 用户密码
func (obj *_UserMgr) GetFromPasswd(passwd string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`passwd` = ?", passwd).Find(&results).Error

	return
}

// GetBatchFromPasswd 批量查找 用户密码
func (obj *_UserMgr) GetBatchFromPasswd(passwds []string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`passwd` IN (?)", passwds).Find(&results).Error

	return
}

// GetFromTel 通过tel获取内容 联系电话
func (obj *_UserMgr) GetFromTel(tel int) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`tel` = ?", tel).Find(&results).Error

	return
}

// GetBatchFromTel 批量查找 联系电话
func (obj *_UserMgr) GetBatchFromTel(tels []int) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`tel` IN (?)", tels).Find(&results).Error

	return
}

// GetFromAvator 通过avator获取内容 头像
func (obj *_UserMgr) GetFromAvator(avator string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`avator` = ?", avator).Find(&results).Error

	return
}

// GetBatchFromAvator 批量查找 头像
func (obj *_UserMgr) GetBatchFromAvator(avators []string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`avator` IN (?)", avators).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_UserMgr) FetchByPrimaryKey(uid int) (result User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`uid` = ?", uid).First(&result).Error

	return
}
