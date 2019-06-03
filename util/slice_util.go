package util

/**
反转切片
*/
func ReverseSlice(slice []interface{}) []interface{} {
	for i := len(slice)/2 - 1; i >= 0; i-- {
		opp := len(slice) - 1 - i
		slice[i], slice[opp] = slice[opp], slice[i]
	}

	return slice
}
