package apiresponse

import (
	"github.com/gin-gonic/gin"
	"github.com/nathakusuma/bcc-be-freepass-2024/util/errortypes"
)

type res struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func Success(ctx *gin.Context, httpCode int, message string, data any) {
	ctx.JSON(httpCode, res{
		Message: message,
		Data:    data,
	})
}

func ApiError(ctx *gin.Context, err *errortypes.ApiError) {
	if !gin.IsDebugging() {
		err.Data = gin.H{}
	}

	ctx.JSON(err.Code, res{
		Message: err.Message,
		Data:    err.Data,
	})
}
