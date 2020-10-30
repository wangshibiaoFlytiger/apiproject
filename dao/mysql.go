package dao

import (
	"apiproject/config"
	"apiproject/util"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

/**
@author 王世彪
	个人博客: https://sofineday.com?from=apiproject
	微信: 645102170
	QQ: 645102170
*/

var Db *gorm.DB

/**
初始化数据库连接
*/
func Init() {
	if !config.GlobalConfig.MysqlSwitch {
		return
	}

	var err error
	Db, err = gorm.Open("mysql", config.GlobalConfig.MysqlUrl)
	if err != nil {
		log.Panicln("初始化数据库连接, 异常", err)
	}

	//配置数据库连接池
	Db.DB().SetMaxIdleConns(config.GlobalConfig.MysqlMaxIdleCount)
	Db.DB().SetMaxOpenConns(config.GlobalConfig.MysqlMaxOpenCount)

	//为插入,更新,删除操作替换默认回调
	//在默认回调gorm:before_create之前注册自定义回调,用于自动生成主键ID
	Db.Callback().Create().Before("gorm:before_create").Register("generate_id", generateIdForCreateCallback)

	// 启用Logger，显示详细日志
	Db.LogMode(true)

	log.Println("初始化数据库连接, 完成")
}

/**
模型创建之前自动生成主键ID
*/
func generateIdForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		if idField, ok := scope.FieldByName("ID"); ok {
			if idField.IsBlank {
				err := idField.Set(util.GenUniqueId())
				fmt.Println(err)
			}
		}
	}
}
