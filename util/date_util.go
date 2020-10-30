package util

import (
	"apiproject/log"
	"fmt"
	"go.uber.org/zap"
	"time"
)

/**
@author 王世彪
	个人博客: https://sofineday.com?from=apiproject
	微信: 645102170
	QQ: 645102170
*/

const (
	timeFormart = "2006-01-02 15:04:05"
)

/**
格式化时间
*/
func FormatTime(time time.Time) (dateTimeFormatted string) {
	dateTimeFormatted = fmt.Sprintf("%s", time.Local().Format("2006-01-02 15:04:05"))
	return
}

/**
解析时间字符串
*/
func ParseTime(timeStr string) time.Time {
	timeParsed, err := time.ParseInLocation(timeFormart, timeStr, time.Local)
	if err != nil {
		log.Logger.Error("解析时间字符串, 失败", zap.Any("timeStr", timeStr), zap.Error(err))
	}

	return timeParsed
}

/**
获取今天零点的时间
*/
func GetBeginTimeToday() time.Time {
	timeStr := time.Now().Format("2006-01-02")
	beginTimeToday, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	return beginTimeToday
}

/**
获取当前的年
*/
func GetCurrentYear() string {
	return time.Now().Format("2006")
}

/**
获取当前的日期字符串, 精确到日,可以指定分隔符
*/
func GetCurrentDateString(split string) string {
	return time.Now().Format("2006" + split + "01" + split + "02")
}
