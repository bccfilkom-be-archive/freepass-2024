package main

import (
	"freepass-2024/controllers"
	"freepass-2024/initializers"
	"freepass-2024/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.POST("/api/register", controllers.Register)
	r.POST("/api/login", controllers.Login)
	r.PUT("/api/edit", middleware.RequireAuth, controllers.Edit)

	r.GET("/api/posts", middleware.RequireAuth, controllers.PostsIndex)
	r.GET("/api/posts/:id", middleware.RequireAuth, controllers.PostsShow)

	r.POST("/api/posts", middleware.RequireAuth, controllers.CreatePost)
	r.PUT("/api/posts/:id", middleware.RequireAuth, controllers.UpdatePost)
	r.DELETE("/api/posts/:id", middleware.RequireAuth, controllers.DeletePost)

	r.Run()
}
