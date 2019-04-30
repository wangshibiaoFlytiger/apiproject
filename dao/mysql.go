package dao

import (
	"apiproject/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

/**
数据库连接初始化
*/
func init() {
	//读取配置
	var err error
	mysqlUrl, err := config.Config.GetValue("mysql", "url")
	if err != nil {
		panic(err)
	}
	maxIdleCount, err := config.Config.Int("mysql", "max.idle.count")
	if err != nil {
		panic(err)
	}
	maxOpenCount, err := config.Config.Int("mysql", "max.open.count")
	if err != nil {
		panic(err)
	}

	Db, err = gorm.Open("mysql", mysqlUrl)
	if err != nil {
		panic(err)
	}

	//配置数据库连接池
	Db.DB().SetMaxIdleConns(maxIdleCount)
	Db.DB().SetMaxOpenConns(maxOpenCount)

	// 启用Logger，显示详细日志
	Db.LogMode(true)
}
