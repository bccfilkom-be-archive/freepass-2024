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

	r.GET("/api/posts", middleware.RequireAuth, controllers.FetchPosts)
	r.POST("/api/posts/:id", middleware.RequireAuth, controllers.AddComment)
	r.GET("/api/posts/:id", middleware.RequireAuth, controllers.ViewPost)
	
	r.PUT("/api/vote/:id", middleware.RequireAuth, controllers.CastVote)

	r.PUT("/api/admin/promote/:id", middleware.RequireAuth, controllers.PromoteUser)
	r.GET("/api/admin/posts", middleware.RequireAuth, controllers.FetchPosts)
	r.POST("/api/admin/setElection", middleware.RequireAuth, controllers.SetElection)
	r.DELETE("/api/admin/deleteUser/:id", middleware.RequireAuth, controllers.DeleteUser)
	r.DELETE("/api/admin/deletePost/:id", middleware.RequireAuth, controllers.DeletePostAdmin)
	r.DELETE("/api/admin/deleteComment/:id", middleware.RequireAuth, controllers.DeleteComment)

	r.POST("/api/posts", middleware.RequireAuth, controllers.CreatePost)
	r.PUT("/api/posts/:id", middleware.RequireAuth, controllers.UpdatePost)
	r.DELETE("/api/posts/:id", middleware.RequireAuth, controllers.DeletePost)

	r.Run()
}
