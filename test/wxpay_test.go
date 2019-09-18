package test

import (
	"fmt"
	"github.com/iGoogle-ink/gopay"
	"testing"
)

/**
测试微信支付的统一下单接口:github.com/iGoogle-ink/gopay
*/
func TestWxpayUnifiedOrder(t *testing.T) {
	//初始化微信客户端
	//    appId：应用ID
	//    MchID：商户ID
	//    apiKey：API秘钥值
	//    isProd：是否是正式环境
	client := gopay.NewWeChatClient("appId", "商户ID", "apikey", true)

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

/**
测试微信支付查询订单接口
*/
func TestWxpayQueryOrder(t *testing.T) {
	//初始化微信客户端
	//    appId：应用ID
	//    MchID：商户ID
	//    apiKey：API秘钥值
	//    isProd：是否是正式环境
	client := gopay.NewWeChatClient("appId", "商户ID", "apikey", true)

	//初始化参数结构体
	body := make(gopay.BodyMap)
	body.Set("out_trade_no", "1174152428149096448")
	body.Set("nonce_str", gopay.GetRandomString(32))
	body.Set("sign_type", gopay.SignType_MD5)

	//请求订单查询，成功后得到结果
	wxRsp, err := client.QueryOrder(body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if wxRsp.ResultCode == "SUCCESS" && wxRsp.ReturnCode == "SUCCESS" && wxRsp.ReturnMsg == "OK" && wxRsp.TradeState == "SUCCESS" {
		fmt.Println("支付成功")
	} else {
		fmt.Println("非支付成功的订单")
	}
	fmt.Println("wxRsp：", *wxRsp)
}
