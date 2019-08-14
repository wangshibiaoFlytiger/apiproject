package util

import "encoding/base64"

/**
base64编码
*/
func Base64Encode(src string) string {
	return base64.StdEncoding.EncodeToString([]byte(src))
}
