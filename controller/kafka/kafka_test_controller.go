package c_kafka

import (
	s_kafka "apiproject/service/kafka"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
发送消息
*/
func SendMessage(ctx *gin.Context) {
	s_kafka.KafkaService.SendKafkaMessage("test", "232323")
	ctx.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": nil,
	})
}
