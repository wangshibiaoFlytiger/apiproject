package middleware

import (
	"apiproject/ip_location"
	"apiproject/log"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

/**
api访问日志中间件
*/
func ApiLogMiddleware(ctx *gin.Context) {
	startTime := time.Now()

	ctx.Next()

	//请求耗时
	duration := time.Since(startTime)
	log.Logger.Info("用户访问日志", zap.String("uri", ctx.Request.RequestURI), zap.String("ip", ctx.ClientIP()), zap.Any("ipLocation", ip_location.GetIpLocationString(ctx.ClientIP())), zap.Any("duration", duration))
}
