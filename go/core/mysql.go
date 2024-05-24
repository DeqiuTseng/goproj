package core

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var LOG = log.New(os.Stdout, "", log.Lshortfile|log.Ldate|log.Ltime)

// InitMysql 初始化mysql
func InitMysql() (*gorm.DB, error) {
	// mysqlConfig := Config.Database.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		// mysqlConfig.UserName,
		"root",
		// mysqlConfig.Password,
		"Z2q.090511",
		// mysqlConfig.Host,
		"localhost",
		// mysqlConfig.Port,
		"3306",
		// mysqlConfig.Database)
		"demo")
	LOG.Println("dsn: ", dsn)
	dbMysql, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		LOG.Println("open db_mysql error ", err)
		return nil, err
	} else {
		LOG.Println("open db_mysql sucess")
	}
	DB = dbMysql

	//迁移表
	// autoMigrateTable()

	// 是否打开日志
	/*
		if mysqlConfig.LogMode {
			dbMysql.Debug()
		}*/

	db, _ := dbMysql.DB()
	//设置连接池的最大闲置连接数
	db.SetMaxIdleConns(10)
	//设置连接池中的最大连接数量
	db.SetMaxOpenConns(100)
	//设置连接的最大复用时间
	db.SetConnMaxLifetime(10 * time.Second)
	return dbMysql, nil
}

/*
// 自动迁移表
func autoMigrateTable() {
	err := DB.AutoMigrate(&do.User{}, &do.OperationLog{})
	if err != nil {
		LOG.Println("迁移表结构失败：", err)
	}
}*/
