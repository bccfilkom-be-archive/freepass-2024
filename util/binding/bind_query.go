package binding

import (
	"github.com/gin-gonic/gin"
	"github.com/nathakusuma/bcc-be-freepass-2024/util/errortypes"
	"net/http"
	"strconv"
)

func ShouldQueryExist(ctx *gin.Context, query string) (string, *errortypes.ApiError) {
	value, isExist := ctx.GetQuery(query)
	if !isExist {
		return "", &errortypes.ApiError{
			Code:    http.StatusBadRequest,
			Message: "missing query: " + query,
			Data:    gin.H{},
		}
	}
	return value, nil
}

func ShouldUintQuery(ctx *gin.Context, query string) (uint, *errortypes.ApiError) {
	valueStr, err := ShouldQueryExist(ctx, query)
	if err != nil {
		return 0, err
	}

	value, err2 := strconv.ParseUint(valueStr, 10, 64)
	if err2 != nil {
		return 0, &errortypes.ApiError{
			Code:    http.StatusBadRequest,
			Message: "invalid query: " + query,
			Data:    gin.H{},
		}
	}

	return uint(value), nil
}
