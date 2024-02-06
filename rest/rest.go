package rest

import (
	post_handler "freepass-bcc/app/post/handler"
	user_handler "freepass-bcc/app/user/handler"
	"freepass-bcc/middleware"

	"github.com/gin-gonic/gin"
)

type Rest struct {
	gin *gin.Engine
}

func NewRest(gin *gin.Engine) Rest {
	return Rest{
		gin: gin,
	}
}

func (r *Rest) RouteUser(userHandler *user_handler.UserHandler) {
	validate := middleware.RequireAuth
	adminOnly := middleware.CheckAdmin

	r.gin.GET("/users/candidates-information", validate, userHandler.GetCandidates)
	r.gin.POST("/signup", userHandler.SignUp)
	r.gin.POST("/login", userHandler.LoginUser)
	r.gin.PUT("/users/:userId", validate, adminOnly, userHandler.PromoteUser)
	r.gin.PUT("/users", validate, userHandler.UpdateAccount)
	r.gin.DELETE("/users/:userId", validate, adminOnly, userHandler.DeleteAccount)
}

func (r *Rest) RoutePost(postHandler *post_handler.PostHandler) {
	validate := middleware.RequireAuth
	CandidateOnly := middleware.CheckCandidate

	r.gin.POST("/posts", validate, CandidateOnly, postHandler.CreatePost)
	r.gin.PUT("/posts/:postId", validate, CandidateOnly, postHandler.UpdatePost)
}

func (r *Rest) Run() {
	r.gin.Run()
}
