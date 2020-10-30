package c_wxpay

import (
	s_wxpay "apiproject/service/wxpay"
	"apiproject/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
@author 王世彪
	个人博客: https://sofineday.com?from=apiproject
	微信: 645102170
	QQ: 645102170
*/

/**
微信H5支付
*/

/************************start swagger api定义注解 **************/
// @Summary 微信H5支付
// @Description 微信H5支付
// @Tags 支付
// @Accept  json
// @Produce  json
// @Success 200 {object} gin.H
// @Router /api/wxpay/wxH5Pay [post]
/************************end swagger api定义注解 **************/
func WxH5Pay(ctx *gin.Context) {
	mwebUrl, success := s_wxpay.WxpayService.WxH5pay("标题", util.GenUniqueId().String(), 1, ctx.ClientIP())
	if !success {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 0,
			"data": nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": map[string]interface{}{"mwebUrl": mwebUrl},
	})
}

/**
微信H5支付的回调
*/

/************************start swagger api定义注解 **************/
// @Summary 微信H5支付的回调
// @Description 微信H5支付的回调
// @Tags 支付
// @Accept  json
// @Produce  json
// @Success 200 {object} gin.H
// @Router /api/wxpay/wxH5PayCallback [post]
/************************end swagger api定义注解 **************/
func WxH5PayCallback(ctx *gin.Context) {
	ctx.XML(http.StatusOK, s_wxpay.WxpayService.WxH5payCallback())
}
