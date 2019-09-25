package util

import (
	"apiproject/log"
	"encoding/base64"
	"go.uber.org/zap"
)

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

/**
base64解码
*/
func Base64Decode(src string) (result string) {
	bytes, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		log.Logger.Error("base64解码, 异常", zap.Any("string", src), zap.Error(err))
		return
	}

	result = string(bytes)
	return
}
