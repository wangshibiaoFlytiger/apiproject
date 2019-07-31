package util

import (
	"strings"
	"time"
)

/**
获取人性化的时间格式
*/
func GetDurationHuman(duration time.Duration) string {
	durationString := duration.String()
	durationString = strings.ReplaceAll(durationString, "h", ":")
	durationString = strings.ReplaceAll(durationString, "m", ":")
	durationString = strings.ReplaceAll(durationString, "s", "")

	return durationString
}
