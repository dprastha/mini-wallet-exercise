package handler

import (
	"log"
	"mini-wallet-exercise/config"
	"mini-wallet-exercise/interface/http/exception"
	"mini-wallet-exercise/pkg/utils"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func RecoverPanic() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		defer handlePanic(httpContext)

		httpContext.Next()
	}
}

func handlePanic(httoContext *gin.Context) {
	if err := recover(); err != nil {
		panicException := createPanicException(err)
		stackTrace := getStackTrace()

		log.Println(stackTrace)

		errorResponse := utils.ErrorResponse(panicException.ErrorMessage)
		httoContext.JSON(panicException.StatusCode, errorResponse)
		httoContext.Abort()
	}
}

func createPanicException(err interface{}) exception.Exception {
	if ex, ok := err.(exception.Exception); ok {
		return ex
	}

	return exception.Exception{
		ErrorMessage: "Internal Server Error",
		StatusCode:   http.StatusInternalServerError,
	}
}

func getStackTrace() string {
	if config.AppMode != "PROD" {
		return string(debug.Stack())
	}

	return ""
}
