package rest

import (
	user_handler "freepass-bcc/app/user/handler"

	"github.com/gin-gonic/gin"
)

type Rest struct {
	gin *gin.Engine
}

func NewRest(gin *gin.Engine) Rest {
	return Rest{
		gin: gin,
	}
}

func (r *Rest) RouteUser(userHandler *user_handler.UserHandler) {
	r.gin.POST("/signup", userHandler.SignUp)
	r.gin.POST("/login", userHandler.LoginUser)
}

func (r *Rest) Run() {
	r.gin.Run()
}
