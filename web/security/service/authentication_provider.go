package service

import (
	"github.com/gin-gonic/gin"
	SecurityModel "github.com/litecodex/go-web-framework/web/security/model"
)

type IAuthenticationProvider interface {
	GetAuthMethod() string
	// 认证登录信息信息，传待认证的信息。认证成功后，authentication.Authenticated会设置成true，并authentication.Principal填入User信息
	Authenticate(ctx *gin.Context, authentication *SecurityModel.Authentication) (*SecurityModel.Authentication, error)
}
