package response

import (
	"github.com/gin-gonic/gin"
	"github.com/litecodex/go-web-framework/web/exceptions"
	I18nContext "github.com/litecodex/go-web-framework/web/utils/i18n"
	RequestUtil "github.com/litecodex/go-web-framework/web/utils/request"
	"net/http"
)

type Result struct {
	ICode    int         `json:"status"`
	IMessage string      `json:"message"`
	IData    interface{} `json:"data"`
}

func NewResult() *Result {
	return &Result{
		ICode: 1,
	}
}

func NewErrResult(ctx *gin.Context, customErr *exceptions.CustomError) *Result {
	return NewResult().Code(customErr.Code).I18nMsg(ctx, customErr.Message)
}

func (result *Result) Code(code int) *Result {
	result.ICode = code
	return result
}

func (result *Result) Message(msg string) *Result {
	result.IMessage = msg
	return result
}

// 对消息进行国际化翻译
func (result *Result) I18nMsg(ctx *gin.Context, msg string) *Result {
	result.IMessage = I18nContext.GetI18n().TranslateSimple(msg, RequestUtil.GetLanguage(ctx))
	return result
}

func (result *Result) Data(data interface{}) *Result {
	result.IData = data
	return result
}

func HandleRsp(ctx *gin.Context, data interface{}, err error) {
	// 走异常错误处理器
	if err != nil {
		ctx.Error(err)
		return
	}

	// 正常返回数据
	if data == nil {
		ctx.JSON(http.StatusOK, NewResult().Data(make(map[string]string)))
	} else {
		ctx.JSON(http.StatusOK, NewResult().Data(data))
	}
}
