package entity

import (
	"database/sql/driver"
	"fmt"
	"time"
)

/**
@author 王世彪
	个人博客: https://sofineday.com?from=apiproject
	微信: 645102170
	QQ: 645102170
*/

/**
Gorm中自定义Time类型的JSON字段格式
将struct成员类型由time.Time改为JsonTime, 则可实现自定义json序列化后的时间格式
*/

// JsonTime format json time field by myself
type JsonTime struct {
	time.Time
}

// MarshalJSON on JsonTime format Time field with %Y-%m-%d %H:%M:%S
func (this *JsonTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", this.Local().Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

// Value insert timestamp into mysql need this function.
func (this *JsonTime) Value() (driver.Value, error) {
	//若当前时间字段没有取值
	if this == nil {
		return nil, nil
	}

	var zeroTime time.Time
	if this.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}

	return this.Time, nil
}

// Scan valueof time.Time
func (this *JsonTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*this = JsonTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

const (
	timeFormart = "2006-01-02 15:04:05"
)

/**
json反序列化
*/
func (t *JsonTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormart+`"`, string(data), time.Local)
	*t = JsonTime{now}
	return
}
