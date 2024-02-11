package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	middleware "github.com/AkbarFikri/freepassBCC-2024/middlewares"
	RegisteredRoute "github.com/AkbarFikri/freepassBCC-2024/routes/registeredRoute"
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
	RegisteredRoute.AuthRoute(auth)

	post := router.Group("/post", middleware.Auth())
	RegisteredRoute.PostRoute(post)

	comment := router.Group("/comment", middleware.Auth())
	RegisteredRoute.CommentRoute(comment)

	election := router.Group("/election", middleware.Auth())
	RegisteredRoute.ElectionRoute(election)

	admin := router.Group("/admin", middleware.AuthAdmin())
	RegisteredRoute.AdminRoute(admin)

	user := router.Group("/user", middleware.Auth())
	RegisteredRoute.UserRoute(user)

	return router
}
