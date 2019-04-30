package middleware

import (
	"apiproject/log"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

/**
api访问日志中间件
*/
func ApiLogMiddleware(ctx *gin.Context) {
	ip := ctx.ClientIP()
	log.Logger.Info("接口访问日志", zap.String("uri", ctx.Request.RequestURI), zap.String("ip", ip))

	ctx.Next()
}
