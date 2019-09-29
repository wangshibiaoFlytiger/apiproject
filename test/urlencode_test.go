package test

import (
	"apiproject/util"
	"fmt"
	"testing"
)

/**
测试urlencode
*/
func TestUrlEncode(t *testing.T) {
	fmt.Println(util.UrlEncode("https://www.wechatpay.com.cn"))
	fmt.Println(util.UrlDecode("https%3A%2F%2Fwww.wechatpay.com.cn"))
}
