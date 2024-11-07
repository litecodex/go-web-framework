package request

import (
	"github.com/gin-gonic/gin"
	"strings"
)

// 从header中提取语言
func GetLanguage(c *gin.Context) string {
	language := "en" // 默认语言
	acceptLanguage := c.GetHeader("Accept-Language")
	if acceptLanguage != "" {
		languages := strings.Split(acceptLanguage, ",")
		if len(languages) > 0 {
			language = strings.TrimSpace(languages[0])
		}
	}
	return language
}

func GetHeaderMap(c *gin.Context) map[string]string {
	headers := c.Request.Header
	headerMap := make(map[string]string)

	for key, values := range headers {
		// 由于一个头部可能有多个值，我们只取第一个值
		if len(values) > 0 {
			headerMap[key] = values[0]
		}
	}

	return headerMap
}

const keyTraceId = "_traceId"

func SetTraceId(ctx *gin.Context, traceId string) {
	ctx.Set(keyTraceId, traceId)
}

func GetTraceId(ctx *gin.Context) string {
	traceId, exists := ctx.Get(keyTraceId)
	if exists {
		return traceId.(string)
	} else {
		return ""
	}
}
