package middleware

import (
	"apiproject/log"
	"apiproject/util"
	"bytes"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer `json:"body"`
}

func (w ResponseWriter) Write(b []byte) (int, error) {
	log.Logger.Info("加密api响应数据中间件", zap.Any("responseData", string(b)))
	return w.ResponseWriter.Write([]byte(util.Base64EncodeByte(b)))
}

/**
加密api响应数据中间件
*/
func EncryptResponseMiddleware(ctx *gin.Context) {
	responseWriter := &ResponseWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
	ctx.Writer = responseWriter

	ctx.Next()
}
