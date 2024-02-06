package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	controller "github.com/AkbarFikri/freepassBCC-2024/controllers/Auth"
)

func AuthRoutes(route *gin.RouterGroup) {
	route.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Setup Success"})
	})
	route.POST("/register", controller.RegisterController)
	route.POST("/login", controller.LoginController)
}
