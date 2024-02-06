package middleware

import (
	"errors"
	"freepass-bcc/domain"
	"freepass-bcc/help"

	"github.com/gin-gonic/gin"
)

func CheckAdmin(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		help.UnauthorizedResponse(c, "user not found", errors.New(""))
	}

	if user.(domain.Users).Role != "ADMIN" {
		help.UnauthorizedResponse(c, "admin only", errors.New("access denied"))
	}
}

func CheckCandidate(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		help.UnauthorizedResponse(c, "user not found", errors.New(""))
	}

	if user.(domain.Users).Role != "CANDIDATE" {
		help.UnauthorizedResponse(c, "candidate only", errors.New("access denied"))
	}
}

func CheckNotUser(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		help.UnauthorizedResponse(c, "user not found", errors.New(""))
	}

	if user.(domain.Users).Role == "USER" {
		help.UnauthorizedResponse(c, "candidate or admin only", errors.New("access denied"))
	}
}
