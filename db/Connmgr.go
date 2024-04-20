package db

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"stvsljl.com/CSMS/utils"
)

var _db *gorm.DB // 全局连接池

func Connect() {
	if _db == nil {
		dsn := utils.GetSqlConnConfigStr()
		_db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		SetPool()
	} else {
		ReConnect()
	}
}

func ReConnect() {
	Close()
	Connect()
}

func Close() {
	if _db != nil {
		sqldb, _ := _db.DB()
		sqldb.Close()
	}
	_db = nil
}

func SetPool() {
	dbc, err := _db.DB()
	if err != nil {
		fmt.Println("SetPoolerr" + err.Error())
	}
	dbc.SetMaxIdleConns(utils.GetSqlConfig().MaxIdleConns)
	dbc.SetMaxOpenConns(utils.GetSqlConfig().MaxOpenConns)
	dbc.SetConnMaxLifetime(time.Duration(utils.GetSqlConfig().ConnMaxLifetime * int(time.Hour)))
}

func GetConn() *gorm.DB {
	if _db == nil {
		fmt.Println("GetConnErr db is nil")
	}
	return _db
}
