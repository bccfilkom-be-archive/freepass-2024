package binding

import (
	"github.com/gin-gonic/gin"
	"github.com/nathakusuma/bcc-be-freepass-2024/util/errortypes"
	"net/http"
)

func ShouldBindJSON(ctx *gin.Context, request any) *errortypes.ApiError {
	if err := ctx.ShouldBindJSON(&request); err != nil {
		return &errortypes.ApiError{
			Code:    http.StatusBadRequest,
			Message: "invalid request body",
			Data:    err,
		}
	}
	return nil
}
