package entity

import (
	"database/sql/driver"
	"fmt"
	"time"
)

/**
Gorm中自定义Time类型的JSON字段格式
将struct成员类型由time.Time改为JsonTime, 则可实现自定义json序列化后的时间格式
*/

// JsonTime format json time field by myself
type JsonTime struct {
	time.Time
}

// MarshalJSON on JsonTime format Time field with %Y-%m-%d %H:%M:%S
func (t JsonTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

// Value insert timestamp into mysql need this function.
func (t JsonTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan valueof time.Time
func (t *JsonTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = JsonTime{Time: value}
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
