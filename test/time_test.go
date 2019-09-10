package test

import (
	"apiproject/util"
	"log"
	"testing"
	"time"
)

/**
测试格式化时间
*/
func TestFormatTime(t *testing.T) {
	log.Println(util.FormatTime(time.Now()))
}

/**
测试解析时间字符串
*/
func TestParseTime(t *testing.T) {
	log.Println(util.ParseTime("2019-05-01 12:34:23"))
}
