package apis

import (
	"github.com/gin-gonic/gin"
	ResponseModel "github.com/litecodex/go-web-framework/web/model/response"
)

type ControllerInterface func(ctx *gin.Context) (interface{}, error)

// controller层入口方法，统一封装响应数据
func Handler(controllerFunc ControllerInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data, err := controllerFunc(ctx) // 执行业务代码
		ResponseModel.HandleRsp(ctx, data, err)
	}
}
