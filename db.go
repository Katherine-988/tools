package tools

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

type DBConfig struct {
	DSN                    string //user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	MaxIdleConns           int
	MaxOpenConns           int
	ConnMaxLifetimeSeconds int
	ConnMaxIdleTimeSeconds int
}

type DBMgrType struct {
	DB     *gorm.DB
	Config *DBConfig
}

var DBMgr DBMgrType

func (m *DBMgrType) Init() {
	db, err := gorm.Open(mysql.Open(DBMgr.Config.DSN), &gorm.Config{})
	if err != nil {
		log.Fatalln("init db err:", err.Error())
	}
	sqlDb, _ := db.DB()
	// 设置空闲连接池中链接的最大数量
	sqlDb.SetMaxIdleConns(DBMgr.Config.MaxIdleConns)
	// 设置打开数据库链接的最大数量
	sqlDb.SetMaxOpenConns(DBMgr.Config.MaxOpenConns)
	// 设置连接池里面的连接最大存活时长
	sqlDb.SetConnMaxLifetime(time.Duration(DBMgr.Config.ConnMaxIdleTimeSeconds) * time.Second)
	// 设置连接池里面的连接最大空闲时长
	sqlDb.SetConnMaxIdleTime(time.Duration(DBMgr.Config.ConnMaxIdleTimeSeconds) * time.Second)
	DBMgr.DB = db
}
