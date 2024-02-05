package middleware

import (
	"errors"
	"freepass-bcc/domain"
	"freepass-bcc/help"

	"github.com/gin-gonic/gin"
)

func OnlyAdmin(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		help.UnauthorizedResponse(c, "user not found", errors.New(""))
	}

	if user.(domain.Users).Role != "ADMIN" {
		help.UnauthorizedResponse(c, "admin only", errors.New("access denied"))
	}
}
