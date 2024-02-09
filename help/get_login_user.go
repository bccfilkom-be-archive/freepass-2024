package help

import (
	"errors"
	"freepass-bcc/domain"

	"github.com/gin-gonic/gin"
)

func GetLoginUser(c *gin.Context) (domain.Users, error){
	user, exist := c.Get("user")
	if !exist {
		return domain.Users{}, errors.New("account not found")
	}

	return user.(domain.Users), nil
}
