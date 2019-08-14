package util

import "encoding/base64"

/**
base64编码
*/
func Base64EncodeString(src string) string {
	return base64.StdEncoding.EncodeToString([]byte(src))
}

/**
base64编码
*/
func Base64EncodeByte(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}
