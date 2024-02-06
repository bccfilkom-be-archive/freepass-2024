package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	middleware "github.com/AkbarFikri/freepassBCC-2024/middlewares"
	routers "github.com/AkbarFikri/freepassBCC-2024/routes/auth"
)

func SetupRoute() *gin.Engine {
	router := gin.New()
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Route Not Found"})
	})
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.CORSMiddleware())

	auth := router.Group("/auth")
	routers.AuthRoutes(auth)

	return router
}
