package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

/**
数据库连接初始化
*/
func init() {
	var err error
	Db, err = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/spider?charset=utf8&parseTime=True&loc=Local&timeout=10ms")

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	if Db.Error != nil {
		fmt.Printf("database error %v", Db.Error)
	}

	//配置数据库连接池
	Db.DB().SetMaxIdleConns(10)
	Db.DB().SetMaxOpenConns(100)

	// 启用Logger，显示详细日志
	Db.LogMode(true)
}
