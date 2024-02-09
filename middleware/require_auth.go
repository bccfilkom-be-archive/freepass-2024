package middleware

import (
	"errors"
	"freepass-bcc/domain"
	"freepass-bcc/help"
	"freepass-bcc/infrastucture/database"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func RequireAuth(c *gin.Context) {
	bearerToken := c.Request.Header.Get("Authorization")
	token := strings.Split(bearerToken, " ")[1]

	if token == "" {
		help.UnauthorizedResponse(c, "failed to authentication", errors.New("no token detected"))
		return
	}

	userId, expTime, err := help.ValidateToken(token)
	if err != nil {
		help.UnauthorizedResponse(c, "failed to authentication", err)
		return
	}

	var user domain.Users
	err = database.DB.First(&user, "id = ?", userId).Error
	if err != nil {
		help.UnauthorizedResponse(c, "failed to authentication", err)
		return
	}

	if float64(time.Now().Unix()) > expTime {
		help.UnauthorizedResponse(c, "failed to authentication", errors.New("token expired"))
		return
	}

	c.Set("user", user)
	c.Next()
}
