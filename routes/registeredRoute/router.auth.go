package RegisteredRoute

import (
	"net/http"

	"github.com/gin-gonic/gin"

	authController "github.com/AkbarFikri/freepassBCC-2024/controllers/auth"
)

func AuthRoute(route *gin.RouterGroup) {
	route.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Setup Success"})
	})
	route.POST("/admin/login", authController.LoginAdmin)
	route.POST("/admin/register", authController.RegisterAdmin)
	route.POST("/register", authController.RegisterUser)
	route.POST("/login", authController.LoginUser)
}
