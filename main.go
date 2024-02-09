package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rafli5131/freepass-2024/handlers"
	"github.com/rafli5131/freepass-2024/middleware"
	"github.com/rafli5131/freepass-2024/models"
)

func main() {
	r := gin.Default()

	// Routes
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)
	r.PUT("/edit-profile", middleware.AuthMiddleware(), handlers.EditProfile)
	r.DELETE("/delete-user/:userID", middleware.AuthMiddleware(), handlers.DeleteUser)
	r.GET("/users", middleware.AuthMiddleware(), handlers.GetAllUser)
	r.GET("/view-posts", handlers.ViewPosts)
	r.POST("/create-post", middleware.AuthMiddleware(), handlers.CreatePost)
	r.PUT("/update-post/:postID", middleware.AuthMiddleware(), handlers.UpdatePost)
	r.DELETE("/delete-post/:postID", middleware.AuthMiddleware(), handlers.DeletePost)
	r.POST("/comment/:postID", middleware.AuthMiddleware(), handlers.Comment)
	r.DELETE("/delete-comment/:commentID", middleware.AuthMiddleware(), handlers.DeleteComment)
	r.GET("/view-candidate-info", middleware.AuthMiddleware(), handlers.ViewCandidateInfo)
	r.POST("/vote/:candidateID", middleware.AuthMiddleware(), handlers.Vote)
	r.PUT("/promote/:userID", middleware.AuthMiddleware(), handlers.PromoteToCandidate)
	r.GET("/vote", middleware.AuthMiddleware(), handlers.VoteTotal)
	r.POST("/set-election-dates", middleware.AuthMiddleware(), handlers
