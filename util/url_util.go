package util

import (
	"apiproject/log"
	"go.uber.org/zap"
	"net/url"
)

/**
解析url中的path部分
*/
func ParsePath(httpUrl string) string {
	pUrl, err := url.Parse(httpUrl)
	if err != nil {
		log.Logger.Error("解析url中的path部分, 异常", zap.Any("httpUrl", httpUrl), zap.Error(err))
		return ""
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
	schema, _ := ParseSchema(httpUrl)
	host, _ := ParseHost(httpUrl)

	return schema + "://" + host
}

/**
解析url中的host
返回host或host:port
*/
func ParseHost(httpUrl string) (host string, err error) {
	pUrl, err := url.Parse(httpUrl)
	if err != nil {
		log.Logger.Error("解析url中的host, 异常", zap.Any("httpUrl", httpUrl), zap.Error(err))
		return "", err
	}

	return pUrl.Host, nil
}

/**
解析url中的hostname
不包含port
*/
func ParseHostname(httpUrl string) (host string, err error) {
	pUrl, err := url.Parse(httpUrl)
	if err != nil {
		log.Logger.Error("解析url中的hostname, 异常", zap.Any("httpUrl", httpUrl), zap.Error(err))
		return "", err
	}

	return pUrl.Hostname(), nil
}

/**
解析url中的端口
*/
func ParseUrlPort(httpUrl string) (port string, err error) {
	pUrl, err := url.Parse(httpUrl)
	if err != nil {
		log.Logger.Error("解析url中的端口, 异常", zap.Any("httpUrl", httpUrl), zap.Error(err))
		return "", err
	}

	return pUrl.Port(), nil
}

/**
解析url中指定的查询参数
*/
func ParseQueryByName(httpUrl string, name string) (value string, err error) {
	pUrl, err := url.Parse(httpUrl)
	if err != nil {
		log.Logger.Error("解析url中指定的查询参数, 解析url异常", zap.Any("httpUrl", httpUrl), zap.Error(err))
		return "", err
	}

	query, err := url.ParseQuery(pUrl.RawQuery)
	if err != nil {
		log.Logger.Error("解析url中指定的查询参数, 解析url异常", zap.Any("httpUrl", httpUrl), zap.Error(err))
		return "", err
	}

	return query.Get(name), nil
}

/**
解析url的协议
*/
func ParseSchema(httpUrl string) (schema string, err error) {
	pUrl, err := url.Parse(httpUrl)
	if err != nil {
		log.Logger.Error("解析url的协议, 解析url异常", zap.Any("httpUrl", httpUrl), zap.Error(err))
		return "", err
	}

	return pUrl.Scheme, nil
}
