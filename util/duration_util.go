package util

import (
	"apiproject/log"
	"github.com/davidscholberg/go-durationfmt"
	"go.uber.org/zap"
	"time"
)

/**
@author 王世彪
	个人博客: https://sofineday.com?from=apiproject
	微信: 645102170
	QQ: 645102170
*/

/**
获取人性化的时间格式
*/
func GetDurationHuman(duration time.Duration) string {
	durationFormat, err := durationfmt.Format(duration, "%0h:%0m:%0s")
	if err != nil {
		log.Logger.Error("获取人性化的时间格式, 异常", zap.Any("duration", duration), zap.Error(err))
	}

	return durationFormat
}
