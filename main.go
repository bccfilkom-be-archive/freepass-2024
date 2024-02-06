package main

import (
	post_handler "freepass-bcc/app/post/handler"
	post_repository "freepass-bcc/app/post/repository"
	post_usecase "freepass-bcc/app/post/usecase"
	user_handler "freepass-bcc/app/user/handler"
	user_repository "freepass-bcc/app/user/repository"
	user_usecase "freepass-bcc/app/user/usecase"
	comment_handler "freepass-bcc/app/comment/handler"
	comment_repository "freepass-bcc/app/comment/repository"
	comment_usecase "freepass-bcc/app/comment/usecase"
	"freepass-bcc/infrastucture"
	"freepass-bcc/infrastucture/database"
	"freepass-bcc/rest"

	"github.com/gin-gonic/gin"
)

func main() {
	infrastucture.LoadEnv()
	database.ConnectToDB()
	database.Migrate()

	userRepository := user_repository.NewUserRepository(database.DB)
	postRepository := post_repository.NewPostRepository(database.DB)
	commentRepository := comment_repository.NewCommentRepository(database.DB)

	userUsecase := user_usecase.NewUserUsecase(userRepository)
	postUsecase := post_usecase.NewPostUsecase(postRepository)
	commentUsecase := comment_usecase.NewCommentUsecase(commentRepository, postRepository)

	userHandler := user_handler.NewUserHandler(userUsecase)
	postHandler := post_handler.NewPostHandler(postUsecase)
	commentHandler := comment_handler.NewCommentHandler(commentUsecase)

	router := rest.NewRest(gin.Default())

	router.RouteUser(userHandler)
	router.RoutePost(postHandler)
	router.RouteComment(commentHandler)

	router.Run()
}
