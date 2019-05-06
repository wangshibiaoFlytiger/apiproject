package dao

import (
	"apiproject/config"
	"fmt"
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

	//为插入,更新,删除操作替换默认回调
	//Db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	Db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	Db.Callback().Delete().Replace("gorm:delete", deleteCallback)

	// 启用Logger，显示详细日志
	Db.LogMode(true)
}

/**
updateTimeStampForCreateCallback will set `CreatedAt`, `UpdatedAt` when creating
该插入回调没有验证成功, 提示错误:using unaddressable value
*/
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		now := gorm.NowFunc()

		if createdAtField, ok := scope.FieldByName("CreatedAt"); ok {
			if createdAtField.IsBlank {
				err := createdAtField.Set(now)
				fmt.Println(err)
			}
		}

		if updatedAtField, ok := scope.FieldByName("UpdatedAt"); ok {
			if updatedAtField.IsBlank {
				err := updatedAtField.Set(now)
				fmt.Println(err)
			}
		}
	}
}

// updateTimeStampForUpdateCallback will set `UpdatedAt` when updating
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("UpdatedAt", gorm.NowFunc())
	}
}

// deleteCallback used to delete data from database or set deleted_at to current time (when using with soft delete)
func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deletedAtField, hasDeletedAtField := scope.FieldByName("DeletedAt")

		if !scope.Search.Unscoped && hasDeletedAtField {
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedAtField.DBName),
				scope.AddToVars(gorm.NowFunc()),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}

// addExtraSpaceIfExist adds a separator
func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
