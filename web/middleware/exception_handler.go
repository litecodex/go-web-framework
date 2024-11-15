package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/litecodex/go-web-framework/web/exceptions"
	"github.com/litecodex/go-web-framework/web/logger"
	WebModel "github.com/litecodex/go-web-framework/web/model/response"
	"go.uber.org/zap"
	"net/http"
)

var unKnowErr = exceptions.NewErrorCode(500, "Internal System Error")

// 捕捉全局异常并返回统一的JSON响应
func HandleGlobalException() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 处理panic抛出来的异常（这种方式，go一般不推荐）
		defer func() {
			if r := recover(); r != nil {
				var errResp *WebModel.Result
				switch e := r.(type) {
				case *exceptions.CustomError:
					errResp = WebModel.NewErrResult(c, e)
					logger.GetLogger().Error("ErrorLog: ", logger.TraceId(c),
						zap.String("error", e.Error()))
				default:
					errResp = WebModel.NewErrResult(c, exceptions.OfCode(unKnowErr))
					logger.GetLogger().Error("ErrorLog: ",
						logger.TraceId(c),
						zap.String("error", fmt.Sprintf("Recovered from panic: %v", r)))
				}
				c.JSON(http.StatusOK, errResp)
				c.Abort()
			}
		}()

		c.Next()

		// 获取正常情况下的错误信息（go推荐这种在方法主动返回的error）
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			var errResp *WebModel.Result
			var customError *exceptions.CustomError
			switch {
			case errors.As(err, &customError):
				errResp = WebModel.NewResult().
					Code(customError.Code).
					I18nMsg(c, customError.GetI18nMsgTemplate())
				logger.GetLogger().Warn("ErrorLog: ", logger.TraceId(c),
					zap.String("errMsg", customError.Message))
			default:
				errResp = WebModel.NewErrResult(c, exceptions.OfCode(unKnowErr))
				logger.GetLogger().Warn("ErrorLog: ", logger.TraceId(c), zap.String("errMsg", err.Error()))
			}
			c.JSON(http.StatusOK, errResp)
			c.Abort()
		}
	}
}
