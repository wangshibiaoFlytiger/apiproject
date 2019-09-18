package s_wxpay

import (
	"apiproject/config"
	"apiproject/log"
	"github.com/iGoogle-ink/gopay"
	"go.uber.org/zap"
)

type WxpayService struct {
}

/**
微信H5支付
*/
func (this *WxpayService) WxH5pay(title string, orderNo string, fee int64, userIp string) (mwebUrl string, success bool) {
	//初始化微信客户端
	//    appId：应用ID
	//    MchID：商户ID
	//    apiKey：API秘钥值
	//    isProd：是否是正式环境
	client := gopay.NewWeChatClient(config.GlobalConfig.WxpayH5Appid, config.GlobalConfig.WxpayH5Mchid, config.GlobalConfig.WxpayH5Apikey, true)

	//初始化参数Map
	body := make(gopay.BodyMap)
	body.Set("nonce_str", gopay.GetRandomString(32))
	body.Set("body", title)
	body.Set("out_trade_no", orderNo)
	body.Set("total_fee", fee)
	body.Set("spbill_create_ip", userIp)
	body.Set("notify_url", config.GlobalConfig.WxpayH5Notifyurl)
	body.Set("trade_type", gopay.TradeType_H5)

	//请求支付下单，成功后得到结果
	wxRsp, err := client.UnifiedOrder(body)
	if err != nil {
		log.Logger.Error("微信H5支付, 请求api异常", zap.Error(err))
		return "", false
	}

	success = false
	if wxRsp.ResultCode == "SUCCESS" && wxRsp.ReturnCode == "SUCCESS" {
		success = true
	}

	log.Logger.Info("微信H5支付完成", zap.Any("requestPara", body), zap.Any("responseData", wxRsp))
	return wxRsp.MwebUrl, success
}

/**
查询微信订单是否支付完成
*/
func (this *WxpayService) IsOrderSuccess(orderNo string) (success bool) {
	//初始化微信客户端
	//    appId：应用ID
	//    MchID：商户ID
	//    apiKey：API秘钥值
	//    isProd：是否是正式环境
	client := gopay.NewWeChatClient(config.GlobalConfig.WxpayH5Appid, config.GlobalConfig.WxpayH5Mchid, config.GlobalConfig.WxpayH5Apikey, true)

	//初始化参数结构体
	body := make(gopay.BodyMap)
	body.Set("out_trade_no", orderNo)
	body.Set("nonce_str", gopay.GetRandomString(32))
	body.Set("sign_type", gopay.SignType_MD5)

	//请求微信订单查询接口
	wxRsp, err := client.QueryOrder(body)
	if err != nil {
		log.Logger.Error("查询微信订单是否支付完成, 异常", zap.Error(err))
		return false
	}

	if wxRsp.ResultCode == "SUCCESS" && wxRsp.ReturnCode == "SUCCESS" && wxRsp.ReturnMsg == "OK" && wxRsp.TradeState == "SUCCESS" {
		log.Logger.Info("查询微信订单是否支付完成, 已支付成功", zap.Any("requestBody", body), zap.Any("responseBody", wxRsp))
		return true
	} else {
		log.Logger.Info("查询微信订单是否支付完成, 非支付成功的订单", zap.Any("requestBody", body), zap.Any("responseBody", wxRsp))
		return false
	}
}

/**
微信H5支付的回调
*/
func (this *WxpayService) WxH5payCallback() interface{} {
	type xml struct {
		ReturnCode string `xml:"return_code"`
		ReturnMsg  string `xml:"return_msg"`
	}

	result := xml{"SUCCESS", "OK"}
	return result
}
