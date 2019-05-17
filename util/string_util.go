package util

/**
判断字符串是否包含于数组
*/
func Contains(arr []string, elem string) bool {
	for index := range arr {
		if arr[index] == elem {
			return true
		}
	}

	return false
}
