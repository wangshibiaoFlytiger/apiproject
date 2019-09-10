package util

import (
	"apiproject/log"
	"fmt"
	"go.uber.org/zap"
	"time"
)

const (
	timeFormart = "2006-01-02 15:04:05"
)

/**
格式化时间
*/
func FormatTime(time time.Time) (dateTimeFormatted string) {
	dateTimeFormatted = fmt.Sprintf("\"%s\"", time.Format("2006-01-02 15:04:05"))
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
