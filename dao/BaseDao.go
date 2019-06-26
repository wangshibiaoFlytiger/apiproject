package dao

import (
	"apiproject/log"
	"errors"
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
		log.Logger.Sugar().Info("App elapsed: ", elapsed)
	}()

	dataType := reflect.TypeOf(values)
	if dataType.Kind() != reflect.Slice {
		return errors.New("values must be a slice!")
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

			vals = append(vals, field.Field.Interface())
		}
	}

	sqlStr = sqlStr + strings.Join(inserts, ",")

	err := Db.Exec(sqlStr, vals...).Error
	if err != nil {

	}

	return err
}
