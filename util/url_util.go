package util

import "net/url"

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
