package test

import (
	"fmt"
	"github.com/objcoding/wxpay"
	"log"
	"testing"
)

/**
测试微信支付
*/
func TestWxpay(t *testing.T) {
	client := wxpay.NewClient(wxpay.NewAccount("wx2ac89d28a6b7f7b3", "1326088401", "579617ce82484c97800e93812f4020ef", false))
	params := make(wxpay.Params)
	params.SetString("body", "标题").
		SetString("out_trade_no", "58867657575757").
		SetInt64("total_fee", 1).
		SetString("spbill_create_ip", "127.0.0.1").
		SetString("notify_url", "http://notify.objcoding.com/notify").
		SetString("trade_type", "MWEB")
	log.Println(client.UnifiedOrder(params))

	paramResponse, err := client.UnifiedOrder(params)
	if err != nil {
		log.Panicln(err)
	}
	log.Println("mweb_url", paramResponse.GetString("mweb_url"))

	// 校验签名
	success := client.ValidSign(paramResponse)
	fmt.Println("验证签名:", success)
}
