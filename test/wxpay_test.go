package test

import (
	"fmt"
	"github.com/iGoogle-ink/gopay"
	"testing"
)

/**
测试微信支付:github.com/iGoogle-ink/gopay
*/
func TestWxpay(t *testing.T) {
	//初始化微信客户端
	//    appId：应用ID
	//    MchID：商户ID
	//    apiKey：API秘钥值
	//    isProd：是否是正式环境
	client := gopay.NewWeChatClient("wx2ac89d28a6b7f7b3", "1326088401", "579617ce82484c97800e93812f4020ef", true)

	//初始化参数Map
	body := make(gopay.BodyMap)
	body.Set("nonce_str", gopay.GetRandomString(32))
	body.Set("body", "测试支付")
	number := gopay.GetRandomString(32)
	fmt.Println("out_trade_no:", number)
	body.Set("out_trade_no", number)
	body.Set("total_fee", 1)
	body.Set("spbill_create_ip", "127.0.0.1")
	body.Set("notify_url", "http://www.gopay.ink")
	body.Set("trade_type", gopay.TradeType_H5)
	//body.Set("sign_type", gopay.SignType_MD5)

	//请求支付下单，成功后得到结果
	wxRsp, err := client.UnifiedOrder(body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("wxRsp:", *wxRsp)
}
