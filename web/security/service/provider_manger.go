package service

import (
	"github.com/gin-gonic/gin"
	"github.com/litecodex/go-web-framework/web/exceptions"
	SecurityModel "github.com/litecodex/go-web-framework/web/security/model"
)

type AuthProviderManager struct {
	// key是authMethod
	authenticationProviderMap map[string]IAuthenticationProvider
}

func NewAuthProviderManager() *AuthProviderManager {
	return &AuthProviderManager{
		authenticationProviderMap: make(map[string]IAuthenticationProvider),
	}
}

func (thiz *AuthProviderManager) AddProvider(provider IAuthenticationProvider) {
	authMethod := provider.GetAuthMethod()
	_, exists := thiz.authenticationProviderMap[authMethod]
	if exists {
		panic("authMethod: " + authMethod + " already bind provider!")
	}
	thiz.authenticationProviderMap[authMethod] = provider
}

func (thiz AuthProviderManager) Authenticate(ctx *gin.Context,
	authentication *SecurityModel.Authentication) (*SecurityModel.Authentication, error) {
	// 支持多种认证方式
	provider, exists := thiz.authenticationProviderMap[authentication.AuthMethod]
	if !exists {
		return nil, exceptions.OfMessage("unknown login type")
	}
	return provider.Authenticate(ctx, authentication)
}
