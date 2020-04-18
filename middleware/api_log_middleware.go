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

	//解析query参数和form参数
	if err := ctx.Request.ParseForm(); err != nil {
		log.Logger.Error("api访问日志中间件, 解析参数, 异常", zap.Error(err))
	}

	ctx.Next()

	//请求耗时
	duration := time.Since(startTime)
	log.Logger.Info("用户访问日志", zap.String("uri", ctx.Request.URL.Path), zap.Any("method", ctx.Request.Method), zap.Any("queryParaList", ctx.Request.URL.Query()), zap.Any("formParaList", ctx.Request.PostForm), zap.Any("header", ctx.Request.Header), zap.Any("userAgent", ctx.Request.UserAgent()), zap.Any("cookies", ctx.Request.Cookies()), zap.String("ip", ctx.ClientIP()), zap.Any("ipLocation", ip_location.GetIpLocationString(ctx.ClientIP())), zap.Any("duration", duration), zap.Any("durationNanosecond", int64(duration)))
}
