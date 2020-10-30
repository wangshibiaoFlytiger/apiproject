package test

import (
	"fmt"
	"github.com/imroc/req"
	"github.com/tidwall/gjson"
	"testing"
)

/**
@author 王世彪
	个人博客: https://sofineday.com?from=apiproject
	微信: 645102170
	QQ: 645102170
*/

/**
测试http请求库github.com/imroc/req
*/
func TestHttpRequest(t *testing.T) {
	params := req.Param{
		"page": 1,
		"rows": 20,
	}
	resp, err := req.Get("http://spideradmin.oupeng.com/galleries/list?page=1&rows=10&status=1", params)
	if err != nil {
		panic(err)
	}

	galleryList := gjson.Get(resp.String(), "data").Array()
	for _, gallery := range galleryList {
		fmt.Println(gallery.Get("verticalUrl"))
	}
}
