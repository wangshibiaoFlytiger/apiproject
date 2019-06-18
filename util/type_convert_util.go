package util

import "strconv"

func Int64ToInt(int64Value int64) int {
	strValue := strconv.FormatInt(int64Value, 10)
	intValue, _ := strconv.Atoi(strValue)

	return intValue
}

func IntToInt64(intValue int) int64 {
	strValue := strconv.Itoa(intValue)
	int64Value, _ := strconv.ParseInt(strValue, 10, 64)

	return int64Value
}
