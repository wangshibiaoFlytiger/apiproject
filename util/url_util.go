package util

import (
	"apiproject/log"
	"go.uber.org/zap"
	"net/url"
	"strings"
)

/**
解析url中的path部分
*/
func ParsePath(httpUrl string) string {
	pUrl, err := url.Parse(httpUrl)
	if err != nil {
		panic(err)
	}

	return pUrl.Path
}

/**
urlencode
*/
func UrlEncode(input string) string {
	return url.QueryEscape(input)
}

/**
urldecode
*/
func UrlDecode(input string) string {
	value, err := url.QueryUnescape(input)
	if err != nil {
		log.Logger.Error("urldecode, 失败", zap.Any("input", input), zap.Error(err))
	}

	return value
}

/**
解析url中path之前的部分
*/
func ParseUrlPrefix(httpUrl string) string {
	path := ParsePath(httpUrl)
	prefix := strings.ReplaceAll(httpUrl, path, "")
	return prefix
}
