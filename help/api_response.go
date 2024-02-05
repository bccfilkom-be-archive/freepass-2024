package help

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func FailedResponse(c *gin.Context, code int, message string, err error) {
	c.JSON(code, gin.H{
		"message": message,
		"error":   err.Error(),
	})
}

func SuccessResponse(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"message": message,
		"data":    data,
	})
}

type ErrorObject struct {
	Code    int
	Message string
	Err     error
}
