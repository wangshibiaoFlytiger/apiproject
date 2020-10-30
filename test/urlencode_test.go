package test

import (
	"apiproject/util"
	"fmt"
	"testing"
)

/**
@author 王世彪
	个人博客: https://sofineday.com?from=apiproject
	微信: 645102170
	QQ: 645102170
*/

/**
测试urlencode
*/
func TestUrlEncode(t *testing.T) {
	fmt.Println(util.UrlEncode("https://www.wechatpay.com.cn"))
	fmt.Println(util.UrlDecode("https%3A%2F%2Fwww.wechatpay.com.cn"))
}
