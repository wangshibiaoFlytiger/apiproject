package test

import (
	"apiproject/util"
	"testing"
)

/**
@author 王世彪
	个人博客: https://sofineday.com?from=apiproject
	微信: 645102170
	QQ: 645102170
*/

func TestUrl(t *testing.T) {
	url := "https://v2.baidu.com:2000/api?name=zhangsan&sex=男"
	schema, _ := util.ParseSchema(url)
	host, _ := util.ParseHost(url)
	hostname, _ := util.ParseHostname(url)
	port, _ := util.ParseUrlPort(url)
	urlPrefix := util.ParseUrlPrefix(url)
	path := util.ParsePath(url)
	value, _ := util.ParseQueryByName(url, "sex")

	println(schema, host, hostname, port, urlPrefix, path, value)
	println(schema + "://" + host + path)
}
