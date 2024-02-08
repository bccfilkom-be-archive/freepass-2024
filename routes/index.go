package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	middleware "github.com/AkbarFikri/freepassBCC-2024/middlewares"
	authRoute "github.com/AkbarFikri/freepassBCC-2024/routes/auth"
	commentRoute "github.com/AkbarFikri/freepassBCC-2024/routes/comment"
	postRoute "github.com/AkbarFikri/freepassBCC-2024/routes/post"
	userRoute "github.com/AkbarFikri/freepassBCC-2024/routes/user"
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
	authRoute.RegisterRoute(auth)

	post := router.Group("/post", middleware.Auth())
	postRoute.RegisterRoute(post)

	comment := router.Group("/comment", middleware.Auth())
	commentRoute.RegisterRoute(comment)

	user := router.Group("/user", middleware.Auth())
	userRoute.RegisterRoute(user)
	return router
}
