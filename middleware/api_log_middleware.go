package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

/**
api访问日志中间件
*/
func ApiLogMiddleware(ctx *gin.Context) {
	ip := ctx.ClientIP()
	fmt.Printf("接口访问日志: uri[%v], 用户IP[%v]", ctx.Request.RequestURI, ip)

	ctx.Next()
}
