package c_kafka

import (
	"apiproject/kafka"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
发送消息
*/
func SendMessage(ctx *gin.Context) {
	kafka.SendKafkaMessage("test", "232323")
	ctx.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": nil,
	})
}
