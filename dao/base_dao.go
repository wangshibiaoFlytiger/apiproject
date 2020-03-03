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

/**
所有数据库操作方法都提供了*gorm.DB类型的形参, 该形参提供2个功能
1. 提供事务支持
当不使用事务时, 传入*gorm.DB类型的对象
当使用事务时, 传入*gorm.db.Begin()返回的对象
2. 可以携带where等查询条件
*/

/**
插入对象
*/
func (this *BaseDao) Insert(db *gorm.DB, item interface{}) (err error) {
	if err = db.Create(item).Error; err != nil {
		log.Logger.Error("插入对象, 异常", zap.Any("item", item), zap.Error(err))
		return err
	}

	return nil
}

/**
批量插入数据  values 参数必须为 数组， validColList 为想插入的字段
注意: 该方法使用raw sql实现, 所以不会触发gorm的回调函数
*/
func (this *BaseDao) BulkInsert(db *gorm.DB, values interface{}, validColList []string) (err error) {
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

	err = db.Exec(sqlStr, vals...).Error
	if err != nil {

	}

	return err
}

/**
查询分页数据:
注意: whereBindTable实参需要绑定具体表
*/
func (this *BaseDao) FindPageData(whereBindTable *gorm.DB, page *model.Page) (err error) {
	whereBindTable.Count(&page.ItemCount)
	whereBindTable = whereBindTable.Offset((page.PageNo - 1) * page.PageSize).Limit(page.PageSize)
	whereBindTable.Order("created_at desc").Find(page.ItemList)
	if err := whereBindTable.Order("created_at desc").Find(page.ItemList).Error; err != nil {
		log.Logger.Error("查询分页数据, 异常", zap.Error(err))
		return err
	}

	page.PageCount = page.ItemCount / page.PageSize

	return nil
}

/**
查询列表
*/
func (this *BaseDao) FindList(where *gorm.DB, itemListOut interface{}) (err error) {
	if err := where.Find(itemListOut).Error; err != nil {
		log.Logger.Error("查询列表, 异常", zap.Error(err))
		return err
	}

	return nil
}

/**
查询单个对象
*/
func (this *BaseDao) Get(where *gorm.DB, itemOut interface{}) (err error) {
	if err := where.First(itemOut).Error; err != nil {
		return err
	}

	return nil
}

func (this *BaseDao) Update(db *gorm.DB, item interface{}) (err error) {
	if err := db.Save(item).Error; err != nil {
		return err
	}

	return nil
}

func (this *BaseDao) Delete(db *gorm.DB, table interface{}) (err error) {
	if err := db.Delete(table).Error; err != nil {
		return err
	}

	return nil
}

/**
获取第一个匹配的结果，或创建一个具有给定条件的新纪录:
注意: 仅适用于struct或map条件
*/
func (this *BaseDao) FirstOrCreate(db *gorm.DB, whereStructOrMap interface{}, itemOut interface{}) (err error) {
	if err := db.Where(whereStructOrMap).FirstOrCreate(itemOut).Error; err != nil {
		return err
	}

	return nil
}
