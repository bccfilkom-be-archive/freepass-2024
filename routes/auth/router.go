package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

)

func AuthRoutes(route *gin.RouterGroup) {
	route.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Setup Success"})
	})
	route.POST("/signup")
	route.POST("/signin")
}
