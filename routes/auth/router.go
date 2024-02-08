package authRoute

import (
	"net/http"

	"github.com/gin-gonic/gin"

	authController "github.com/AkbarFikri/freepassBCC-2024/controllers/Auth"

)

func RegisterRoute(route *gin.RouterGroup) {
	route.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Setup Success"})
	})
	route.POST("/register", authController.RegisterUser)
	route.POST("/login", authController.LoginUser)
}
