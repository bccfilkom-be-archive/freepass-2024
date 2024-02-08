package help

import (
	"freepass-bcc/domain"
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

func UnauthorizedResponse(c *gin.Context, message string, err error) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"message": message,
		"error":   err.Error(),
	})
}

type ErrorObject struct {
	Code    int
	Message string
	Err     error
}

func PostResponse(post domain.Posts, name string) domain.PostResponse {
	if name == "" {
		name = post.User.Name
	}
	PostResponse := domain.PostResponse{
		ID: post.ID,
		Candidate: name,
		Post: post.Post,
	}

	return PostResponse
}
