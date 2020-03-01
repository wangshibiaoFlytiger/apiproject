package dao

import (
	"apiproject/entity"
	"apiproject/log"
	"apiproject/model"
	"errors"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"reflect"
	"strings"
	"time"
)

type BaseDao struct {
}

// 批量插入数据  values 参数必须为 数组， validColList 为想插入的字段
func (this *BaseDao) BulkInsert(values interface{}, validColList []string) error {

	t1 := time.Now()
	defer func() {
		elapsed := time.Since(t1)
		log.Logger.Info("批量插入数据", zap.Any("count", reflect.ValueOf(values).Len()), zap.Any("duration", elapsed))
	}()

	dataType := reflect.TypeOf(values)
	if dataType.Kind() != reflect.Slice {
		return errors.New("批量插入数据, values参数必须为slice类型")
	}

	val := reflect.ValueOf(values)
	if val.Len() <= 0 {
		return nil
	}

	scope := Db.NewScope(val.Index(0).Interface())
	var realColList []string
	if len(validColList) == 0 {
		for _, field := range scope.Fields() {
			realColList = append(realColList, field.DBName)
		}
	} else {
		for _, colName := range validColList {
			realColList = append(realColList, colName)
		}
	}

	var args []string
	for i := 0; i < len(realColList); i++ {
		args = append(args, "?")
	}

	rowSQL := "(" + strings.Join(args, ", ") + ")"

	sqlStr := "REPLACE INTO " + scope.TableName() + "(" + strings.Join(realColList, ",") + ") VALUES "

	var vals []interface{}

	var inserts []string

	for sliceIndex := 0; sliceIndex < val.Len(); sliceIndex++ {
		data := val.Index(sliceIndex).Interface()

		inserts = append(inserts, rowSQL)
		//vals = append(vals, elem.Prop1, elem.Prop2, elem.Prop3)
		elemScope := Db.NewScope(data)
		for _, validCol := range realColList {
			field, ok := elemScope.FieldByName(validCol)
			if !ok {
				return errors.New("can not find col(" + validCol + ")")
			}

			var val interface{}
			value, ok := field.Field.Interface().(entity.JsonTime)
			//对jsonTime自定义时间类型字段特殊处理
			if ok {
				val = value.Time
			} else {
				val = field.Field.Interface()
			}

			vals = append(vals, val)
		}
	}

	sqlStr = sqlStr + strings.Join(inserts, ",")

	err := Db.Exec(sqlStr, vals...).Error
	if err != nil {

	}

	return err
}

/**
查询分页数据:
注意: db实参需要绑定具体表
*/
func (this *BaseDao) FindPageData(db *gorm.DB, page *model.Page) error {
	db.Count(&page.ItemCount)
	db = db.Offset((page.PageNo - 1) * page.PageSize).Limit(page.PageSize)
	db.Order("created_at desc").Find(page.ItemList)

	page.PageCount = page.ItemCount / page.PageSize

	return nil
}
