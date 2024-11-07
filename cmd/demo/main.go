package main

import (
	"github.com/gin-gonic/gin"
	"github.com/litecodex/go-web-framework/web/serverlet"
	"github.com/litecodex/go-web-framework/web/utils/apis"
	"github.com/litecodex/go-web-framework/web/utils/logger"
)

func main() {
	router := serverlet.CreateCommonRouter()
	router.GET("/test", apis.Handler(func(ctx *gin.Context) (interface{}, error) {
		logger.Info(ctx, "testInfo")
		data := map[string]interface{}{
			"key": "jaj",
		}
		return data, nil
	}))
	router.Run(":8080")
}
