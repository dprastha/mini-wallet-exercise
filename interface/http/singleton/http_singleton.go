package singleton

import (
	"context"
	"mini-wallet-exercise/constant"

	"github.com/gin-gonic/gin"
)

func GetContextFromGinContext(httpContext *gin.Context) *context.Context {
	context := httpContext.Request.Context()
	return &context
}

func GetHTTPRequest[T any](httpContext *gin.Context) *T {
	return httpContext.MustGet(constant.RequestBodyJSONKey).(*T)
}
