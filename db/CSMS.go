package db

import (
	"time"
)

/******sql******
CREATE TABLE `Active` (
  `acid` int(11) NOT NULL AUTO_INCREMENT COMMENT '活动ID',
  `name` varchar(100) NOT NULL COMMENT '活动名称',
  `startTime` time NOT NULL COMMENT '开始时间',
  `openTime` time NOT NULL COMMENT '发布时间',
  `endTime` time NOT NULL COMMENT '结束时间',
  `detail` varchar(200) NOT NULL COMMENT '活动详情',
  `text` text NOT NULL COMMENT '活动正文',
  `views` int(10) unsigned NOT NULL COMMENT '浏览量',
  PRIMARY KEY (`acid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='活动表'
******sql******/
// Active 活动表
type Active struct {
	Acid      int       `gorm:"autoIncrement:true;primaryKey;column:acid;type:int(11);not null;comment:'活动ID'"` // 活动ID
	Name      string    `gorm:"column:name;type:varchar(100);not null;comment:'活动名称'"`                          // 活动名称
	Starttime time.Time `gorm:"column:startTime;type:time;not null;comment:'开始时间'"`                             // 开始时间
	Opentime  time.Time `gorm:"column:openTime;type:time;not null;comment:'发布时间'"`                              // 发布时间
	Endtime   time.Time `gorm:"column:endTime;type:time;not null;comment:'结束时间'"`                               // 结束时间
	Detail    string    `gorm:"column:detail;type:varchar(200);not null;comment:'活动详情'"`                        // 活动详情
	Text      string    `gorm:"column:text;type:text;not null;comment:'活动正文'"`                                  // 活动正文
	Views     uint32    `gorm:"column:views;type:int(10) unsigned;not null;comment:'浏览量'"`                      // 浏览量
}

// TableName get sql table name.获取数据库表名
func (m *Active) TableName() string {
	return "Active"
}

/******sql******
CREATE TABLE `ActivityParticipation` (
  `uid` int(10) unsigned NOT NULL COMMENT '参与用户ID',
  `acid` varchar(100) NOT NULL COMMENT '活动ID',
  `status` int(1) unsigned NOT NULL COMMENT '参与状态'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='活动参与状态表'
******sql******/
// Activityparticipation 活动参与状态表
type Activityparticipation struct {
	UID    uint32 `gorm:"column:uid;type:int(10) unsigned;not null;comment:'参与用户ID'"` // 参与用户ID
	Acid   string `gorm:"column:acid;type:varchar(100);not null;comment:'活动ID'"`      // 活动ID
	Status uint32 `gorm:"column:status;type:int(1) unsigned;not null;comment:'参与状态'"` // 参与状态
}

// TableName get sql table name.获取数据库表名
func (m *Activityparticipation) TableName() string {
	return "ActivityParticipation"
}

/******sql******
CREATE TABLE `Admin` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `tel` int(10) unsigned NOT NULL,
  `passwd` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='管理员表'
******sql******/
// Admin 管理员表
type Admin struct {
	ID     int    `gorm:"autoIncrement:true;primaryKey;column:id;type:int(11);not null"`
	Name   string `gorm:"column:name;type:varchar(50);not null"`
	Tel    uint32 `gorm:"column:tel;type:int(10) unsigned;not null"`
	Passwd string `gorm:"column:passwd;type:varchar(100);not null"`
}

// TableName get sql table name.获取数据库表名
func (m *Admin) TableName() string {
	return "Admin"
}

/******sql******
CREATE TABLE `Article` (
  `aid` int(12) NOT NULL AUTO_INCREMENT COMMENT '文章编号',
  `coverimg` varchar(150) NOT NULL COMMENT '封面图片',
  `contentimg` varchar(150) NOT NULL COMMENT '内容大图',
  `title` varchar(50) NOT NULL COMMENT '标题',
  `introduction` varchar(200) NOT NULL COMMENT '简介',
  `text` longtext NOT NULL COMMENT '正文',
  `writetime` datetime NOT NULL COMMENT '发表日期',
  `updatetime` datetime NOT NULL COMMENT '更新日期',
  `author` varchar(10) NOT NULL COMMENT '作者',
  `pageviews` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '浏览量',
  `status` int(1) NOT NULL COMMENT '文章状态',
  PRIMARY KEY (`aid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='网站文章相关数据'
******sql******/
// Article 网站文章相关数据
type Article struct {
	Aid          int       `gorm:"autoIncrement:true;primaryKey;column:aid;type:int(12);not null;comment:'文章编号'"` // 文章编号
	Coverimg     string    `gorm:"column:coverimg;type:varchar(150);not null;comment:'封面图片'"`                     // 封面图片
	Contentimg   string    `gorm:"column:contentimg;type:varchar(150);not null;comment:'内容大图'"`                   // 内容大图
	Title        string    `gorm:"column:title;type:varchar(50);not null;comment:'标题'"`                           // 标题
	Introduction string    `gorm:"column:introduction;type:varchar(200);not null;comment:'简介'"`                   // 简介
	Text         string    `gorm:"column:text;type:longtext;not null;comment:'正文'"`                               // 正文
	Writetime    time.Time `gorm:"column:writetime;type:datetime;not null;comment:'发表日期'"`                        // 发表日期
	Updatetime   time.Time `gorm:"column:updatetime;type:datetime;not null;comment:'更新日期'"`                       // 更新日期
	Author       string    `gorm:"column:author;type:varchar(10);not null;comment:'作者'"`                          // 作者
	Pageviews    uint64    `gorm:"column:pageviews;type:bigint(20) unsigned;not null;default:0;comment:'浏览量'"`    // 浏览量
	Status       int       `gorm:"column:status;type:int(1);not null;comment:'文章状态'"`                             // 文章状态
}

// TableName get sql table name.获取数据库表名
func (m *Article) TableName() string {
	return "Article"
}

/******sql******
CREATE TABLE `Feed` (
  `uid` int(10) unsigned NOT NULL COMMENT '反馈用户',
  `name` varchar(100) NOT NULL COMMENT '反馈名称',
  `feedtime` varchar(100) NOT NULL COMMENT '反馈时间',
  `detail` varchar(200) NOT NULL COMMENT '反馈内容',
  `status` int(11) NOT NULL COMMENT '反馈状态',
  `oid` int(10) unsigned NOT NULL COMMENT '委派人',
  `processor` int(10) unsigned DEFAULT NULL COMMENT '处理人'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='反馈表'
******sql******/
// Feed 反馈表
type Feed struct {
	UID       uint32 `gorm:"column:uid;type:int(10) unsigned;not null;comment:'反馈用户'"`          // 反馈用户
	Name      string `gorm:"column:name;type:varchar(100);not null;comment:'反馈名称'"`             // 反馈名称
	Feedtime  string `gorm:"column:feedtime;type:varchar(100);not null;comment:'反馈时间'"`         // 反馈时间
	Detail    string `gorm:"column:detail;type:varchar(200);not null;comment:'反馈内容'"`           // 反馈内容
	Status    int    `gorm:"column:status;type:int(11);not null;comment:'反馈状态'"`                // 反馈状态
	Oid       uint32 `gorm:"column:oid;type:int(10) unsigned;not null;comment:'委派人'"`           // 委派人
	Processor uint32 `gorm:"column:processor;type:int(10) unsigned;default:null;comment:'处理人'"` // 处理人
}

// TableName get sql table name.获取数据库表名
func (m *Feed) TableName() string {
	return "Feed"
}

/******sql******
CREATE TABLE `OtherUser` (
  `oid` int(11) NOT NULL AUTO_INCREMENT COMMENT '第三方人员ID',
  `name` varchar(100) NOT NULL COMMENT '姓名',
  `company` varchar(100) NOT NULL COMMENT '所属公司',
  `tel` int(11) NOT NULL COMMENT '联系电话',
  `sex` int(1) NOT NULL COMMENT '性别',
  PRIMARY KEY (`oid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='第三方用户表'
******sql******/
// Otheruser 第三方用户表
type Otheruser struct {
	Oid     int    `gorm:"autoIncrement:true;primaryKey;column:oid;type:int(11);not null;comment:'第三方人员ID'"` // 第三方人员ID
	Name    string `gorm:"column:name;type:varchar(100);not null;comment:'姓名'"`                              // 姓名
	Company string `gorm:"column:company;type:varchar(100);not null;comment:'所属公司'"`                         // 所属公司
	Tel     int    `gorm:"column:tel;type:int(11);not null;comment:'联系电话'"`                                  // 联系电话
	Sex     int    `gorm:"column:sex;type:int(1);not null;comment:'性别'"`                                     // 性别
}

// TableName get sql table name.获取数据库表名
func (m *Otheruser) TableName() string {
	return "OtherUser"
}

/******sql******
CREATE TABLE `User` (
  `uid` int(11) NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `name` varchar(50) NOT NULL COMMENT '用户名',
  `passwd` varchar(100) NOT NULL COMMENT '用户密码',
  `tel` int(11) NOT NULL COMMENT '联系电话',
  `avator` varchar(100) NOT NULL COMMENT '头像',
  PRIMARY KEY (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表'
******sql******/
// User 用户表
type User struct {
	UID    int    `gorm:"autoIncrement:true;primaryKey;column:uid;type:int(11);not null;comment:'用户ID'"` // 用户ID
	Name   string `gorm:"column:name;type:varchar(50);not null;comment:'用户名'"`                           // 用户名
	Passwd string `gorm:"column:passwd;type:varchar(100);not null;comment:'用户密码'"`                       // 用户密码
	Tel    int    `gorm:"column:tel;type:int(11);not null;comment:'联系电话'"`                               // 联系电话
	Avator string `gorm:"column:avator;type:varchar(100);not null;comment:'头像'"`                         // 头像
}

// TableName get sql table name.获取数据库表名
func (m *User) TableName() string {
	return "User"
}
