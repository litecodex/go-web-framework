package service

import (
	"github.com/gin-gonic/gin"
	SecurityModel "github.com/litecodex/go-web-framework/web/security/model"
)

type IAuthenticationManager interface {
	Authenticate(ctx *gin.Context, authentication *SecurityModel.Authentication) (*SecurityModel.Authentication, error)
}
