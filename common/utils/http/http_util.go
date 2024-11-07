package http

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/goccy/go-json"
	JSON "github.com/litecodex/go-web-framework/common/utils/json"
	LoggerContext "github.com/litecodex/go-web-framework/web/utils/logger"
	"time"
)

func NewRestyClientWithCtx() *resty.Client {
	client := resty.New().
		SetTimeout(30 * time.Second).
		SetJSONMarshaler(json.Marshal). // 使用高性能的go-json序列化工具
		SetJSONUnmarshaler(json.Unmarshal)

	logger := LoggerContext.GetLogger()
	client.SetLogger(logger.Sugar())

	// 请求之前打印数据
	//client.OnBeforeRequest(func(c *resty.Client, req *resty.Request) error {
	//	logger.Info(fmt.Sprintf("Request: %s %s", req.Method, req.URL))
	//	return nil
	//})

	// 请求之后打印数据
	client.OnAfterResponse(func(c *resty.Client, resp *resty.Response) error {
		requestParam, _ := JSON.Stringify(resp.Request.QueryParam)
		requestBody, _ := JSON.Stringify(resp.Request.Body)
		logger.Info(fmt.Sprintf("Response: %s %s === ReqParams: %s === ReqBody:  %s === RspData: %s",
			resp.Request.Method,
			resp.Request.URL,
			requestParam,
			requestBody,
			resp.String()))
		return nil
	})

	return client
}
