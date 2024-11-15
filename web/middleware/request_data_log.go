package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	LoggerContext "github.com/litecodex/go-web-framework/web/logger"
	"go.uber.org/zap"
	"io"
	"strings"
)

// 定义一个结构体用于保存响应数据
type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w responseBodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// 打印请求、响应数据
func RequestResponseLogger() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		url := ctx.Request.Method + " " + ctx.Request.URL.String()
		headers := ctx.Request.Header
		param := ctx.Request.URL.Query()

		reqBody := ""
		if !isBinary(ctx.GetHeader("Content-Type")) {
			// 读取请求体
			bodyBytes, err := io.ReadAll(ctx.Request.Body)
			if err != nil {
				return
			}
			// 恢复请求体数据
			ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			reqBody = string(bodyBytes)
		} else {
			// 二进制数据不打印！
			reqBody = "[binary data]"
		}

		// 保存原始响应写入器
		writer := &responseBodyWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
		ctx.Writer = writer

		// 处理请求
		ctx.Next()

		// 获取响应的Content-Type,  判断是否为二进制内容类型
		rspBody := ""
		if !isBinary(ctx.Writer.Header().Get("Content-Type")) {
			rspBody = writer.body.String()
		} else {
			rspBody = "[binary data]"
		}

		// 打印响应参数
		LoggerContext.GetLogger().Info("ReqLog: ",
			LoggerContext.TraceId(ctx),
			zap.String("url", url),
			zap.Any("reqParam", param),
			zap.Any("reqBody", reqBody),
			zap.Any("rspBody", rspBody),
			zap.Any("headers", headers),
		)
	}
}

func isBinary(contentType string) bool {
	isBinary := strings.Contains(contentType, "image") ||
		strings.Contains(contentType, "audio") ||
		strings.Contains(contentType, "video") ||
		strings.Contains(contentType, "application/octet-stream")
	return isBinary
}
