package main

import (
	post_handler "freepass-bcc/app/post/handler"
	post_repository "freepass-bcc/app/post/repository"
	post_usecase "freepass-bcc/app/post/usecase"
	user_handler "freepass-bcc/app/user/handler"
	user_repository "freepass-bcc/app/user/repository"
	user_usecase "freepass-bcc/app/user/usecase"
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

	userUsecase := user_usecase.NewUserUsecase(userRepository)
	postUsecase := post_usecase.NewPostUsecase(postRepository)

	userHandler := user_handler.NewUserHandler(userUsecase)
	postHandler := post_handler.NewPostHandler(postUsecase)

	router := rest.NewRest(gin.Default())

	router.RouteUser(userHandler)
	router.RoutePost(postHandler)

	router.Run()
}
