package s_wxpay

import (
	"apiproject/config"
	"apiproject/log"
	"github.com/objcoding/wxpay"
	"go.uber.org/zap"
)

type WxpayService struct {
}

/**
微信H5支付
*/
func (this *WxpayService) WxH5pay(title string, orderNo string, fee int64, userIp string) (mwebUrl string, success bool) {
	client := wxpay.NewClient(wxpay.NewAccount(config.GlobalConfig.WxpayH5Appid, config.GlobalConfig.WxpayH5Mchid, config.GlobalConfig.WxpayH5Apikey, false))
	params := make(wxpay.Params)
	params.SetString("body", title).
		SetString("out_trade_no", orderNo).
		SetInt64("total_fee", fee).
		SetString("spbill_create_ip", userIp).
		SetString("notify_url", config.GlobalConfig.WxpayH5Notifyurl).
		SetString("trade_type", "MWEB")

	paramResponse, err := client.UnifiedOrder(params)
	if err != nil {
		log.Logger.Error("微信H5支付, 请求api异常", zap.Error(err))
		return "", false
	}

	resultCode := paramResponse.GetString("result_code")
	returnCode := paramResponse.GetString("return_code")

	mwebUrl = paramResponse.GetString("mweb_url")
	// 校验签名
	signSuccess := client.ValidSign(paramResponse)

	success = false
	if resultCode == "SUCCESS" && returnCode == "SUCCESS" || signSuccess {
		success = true
	}

	log.Logger.Info("微信H5支付完成", zap.Any("requestPara", params), zap.Any("responsePara", paramResponse), zap.Any("resultCode", resultCode), zap.Any("returnCode", returnCode), zap.Any("signSuccess", signSuccess))
	return "", success
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
