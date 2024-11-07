package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	RequestUtil "github.com/litecodex/go-web-framework/web/utils/request"
)

func CreateTraceId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// traceId用于跟踪请求
		traceId := uuid.NewString()
		RequestUtil.SetTraceId(ctx, traceId)
		ctx.Next()
	}
}
