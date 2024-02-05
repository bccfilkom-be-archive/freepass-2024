package main

import (
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

	userUsecase := user_usecase.NewUserUsecase(userRepository)

	userHandler := user_handler.NewUserHandler(userUsecase)

	router := rest.NewRest(gin.Default())

	router.RouteUser(userHandler)

	router.Run()
}
