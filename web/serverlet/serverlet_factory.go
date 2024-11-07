package serverlet

import (
	"github.com/gin-gonic/gin"
	CustomMiddleware "github.com/litecodex/go-web-framework/web/middleware"
	WebModel "github.com/litecodex/go-web-framework/web/model/response"
	LoggerContext "github.com/litecodex/go-web-framework/web/utils/logger"
	"go.uber.org/zap"
	"net/http"
)

func CreateCommonRouter() *gin.Engine {
	router := gin.Default()
	// 单个上传文件的前 50MB 会暂时存储在内存中，超过部分会写入磁盘的临时文件。
	router.MaxMultipartMemory = 50 << 20 // 50MB
	// 设置NoRoute处理函数，404也返回json结构，修改默认响应数据
	router.NoRoute(func(c *gin.Context) {
		errResp := WebModel.NewResult().Code(404).Message("API not found")
		c.JSON(http.StatusNotFound, errResp)
	})

	// 将 zap 作为 Gin 的日志记录器
	logger := LoggerContext.GetLogger()
	gin.DefaultWriter = zapWriter{logger: logger}
	gin.DefaultErrorWriter = zapWriter{logger: logger}

	// 加载全局Filter
	router.Use(CustomMiddleware.CreateTraceId())         // 生成一个唯一id，用于跟踪请求，注意这个要放第一位
	router.Use(CustomMiddleware.HandleGlobalException()) // 捕捉全局异常
	return router
}

// zapWriter 用于将 Gin 的日志输出重定向到 zap
type zapWriter struct {
	logger *zap.Logger
}

func (z zapWriter) Write(p []byte) (n int, err error) {
	// 输出日志到 zap
	z.logger.Info(string(p))
	return len(p), nil
}
