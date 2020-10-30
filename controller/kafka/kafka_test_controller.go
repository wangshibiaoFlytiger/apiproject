package c_kafka

import (
	s_kafka "apiproject/service/kafka"
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
发送消息
*/

/************************start swagger api定义注解 **************/
// @Summary 发送消息
// @Description 发送消息
// @Tags kafka
// @ID SendMessage
// @Accept  json
// @Produce  json
// @Success 200 {object} gin.H
// @Router /api/kafka/sendMessage [post]
/************************end swagger api定义注解 **************/
func SendMessage(ctx *gin.Context) {
	s_kafka.KafkaService.SendKafkaMessage("test", "232323")
	ctx.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": nil,
	})
}
