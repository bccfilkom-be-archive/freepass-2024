package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rafli5131/freepass-2024/handlers"
	"github.com/rafli5131/freepass-2024/middleware"
)

func main() {
	r := gin.Default()

	// Routes
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)
	r.PUT("/edit-profile", middleware.AuthMiddleware(), handlers.EditProfile)
	r.DELETE("/delete-user/:userID", middleware.AdminAuthMiddleware(), handlers.DeleteUser)
	r.GET("/users", middleware.AdminAuthMiddleware(), handlers.GetAllUser)
	r.GET("/view-posts", handlers.ViewPosts)
	r.POST("/create-post", middleware.CandidateAuthMiddleware(), handlers.CreatePost)
	r.PUT("/update-post/:postID", middleware.CandidateAuthMiddleware(), handlers.UpdatePost)
	r.DELETE("/delete-post/:postID", middleware.CandidateAuthMiddleware(), handlers.DeletePost)
	r.POST("/comment/:postID", middleware.AuthMiddleware(), handlers.Comment)
	r.DELETE("/delete-comment/:commentID", middleware.AdminAuthMiddleware(), handlers.DeleteComment)
	r.GET("/view-candidate-info", middleware.AuthMiddleware(), handlers.ViewCandidateInfo)
	r.POST("/vote/:candidateID", middleware.AuthMiddleware(), handlers.Vote)
	r.PUT("/promote/:userID", middleware.AdminAuthMiddleware(), handlers.PromoteToCandidate)
	r.GET("/vote", middleware.AdminAuthMiddleware(), handlers.VoteTotal)
	r.POST("/set-election-dates", middleware.AdminAuthMiddleware(), handlers.SetElectionDates)

	r.Run(":8080")
}