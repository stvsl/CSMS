package db

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _OtheruserMgr struct {
	*_BaseMgr
}

// OtheruserMgr open func
func OtheruserMgr(db *gorm.DB) *_OtheruserMgr {
	if db == nil {
		panic(fmt.Errorf("OtheruserMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_OtheruserMgr{_BaseMgr: &_BaseMgr{DB: db.Table("OtherUser"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_OtheruserMgr) Debug() *_OtheruserMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_OtheruserMgr) GetTableName() string {
	return "OtherUser"
}

// Reset 重置gorm会话
func (obj *_OtheruserMgr) Reset() *_OtheruserMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_OtheruserMgr) Get() (result Otheruser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Otheruser{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_OtheruserMgr) Gets() (results []*Otheruser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Otheruser{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_OtheruserMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Otheruser{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithOid oid获取 第三方人员ID
func (obj *_OtheruserMgr) WithOid(oid int) Option {
	return optionFunc(func(o *options) { o.query["oid"] = oid })
}

// WithName name获取 姓名
func (obj *_OtheruserMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithCompany company获取 所属公司
func (obj *_OtheruserMgr) WithCompany(company string) Option {
	return optionFunc(func(o *options) { o.query["company"] = company })
}

// WithTel tel获取 联系电话
func (obj *_OtheruserMgr) WithTel(tel int) Option {
	return optionFunc(func(o *options) { o.query["tel"] = tel })
}

// WithSex sex获取 性别
func (obj *_OtheruserMgr) WithSex(sex int) Option {
	return optionFunc(func(o *options) { o.query["sex"] = sex })
}

// GetByOption 功能选项模式获取
func (obj *_OtheruserMgr) GetByOption(opts ...Option) (result Otheruser, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Otheruser{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_OtheruserMgr) GetByOptions(opts ...Option) (results []*Otheruser, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Otheruser{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromOid 通过oid获取内容 第三方人员ID
func (obj *_OtheruserMgr) GetFromOid(oid int) (result Otheruser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Otheruser{}).Where("`oid` = ?", oid).First(&result).Error

	return
}

// GetBatchFromOid 批量查找 第三方人员ID
func (obj *_OtheruserMgr) GetBatchFromOid(oids []int) (results []*Otheruser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Otheruser{}).Where("`oid` IN (?)", oids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 姓名
func (obj *_OtheruserMgr) GetFromName(name string) (results []*Otheruser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Otheruser{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 姓名
func (obj *_OtheruserMgr) GetBatchFromName(names []string) (results []*Otheruser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Otheruser{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromCompany 通过company获取内容 所属公司
func (obj *_OtheruserMgr) GetFromCompany(company string) (results []*Otheruser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Otheruser{}).Where("`company` = ?", company).Find(&results).Error

	return
}

// GetBatchFromCompany 批量查找 所属公司
func (obj *_OtheruserMgr) GetBatchFromCompany(companys []string) (results []*Otheruser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Otheruser{}).Where("`company` IN (?)", companys).Find(&results).Error

	return
}

// GetFromTel 通过tel获取内容 联系电话
func (obj *_OtheruserMgr) GetFromTel(tel int) (results []*Otheruser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Otheruser{}).Where("`tel` = ?", tel).Find(&results).Error

	return
}

// GetBatchFromTel 批量查找 联系电话
func (obj *_OtheruserMgr) GetBatchFromTel(tels []int) (results []*Otheruser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Otheruser{}).Where("`tel` IN (?)", tels).Find(&results).Error

	return
}

// GetFromSex 通过sex获取内容 性别
func (obj *_OtheruserMgr) GetFromSex(sex int) (results []*Otheruser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Otheruser{}).Where("`sex` = ?", sex).Find(&results).Error

	return
}

// GetBatchFromSex 批量查找 性别
func (obj *_OtheruserMgr) GetBatchFromSex(sexs []int) (results []*Otheruser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Otheruser{}).Where("`sex` IN (?)", sexs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_OtheruserMgr) FetchByPrimaryKey(oid int) (result Otheruser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Otheruser{}).Where("`oid` = ?", oid).First(&result).Error

	return
}
