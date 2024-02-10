package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nathakusuma/bcc-be-freepass-2024/database"
	"github.com/nathakusuma/bcc-be-freepass-2024/handler"
	"github.com/nathakusuma/bcc-be-freepass-2024/middleware"
	"github.com/nathakusuma/bcc-be-freepass-2024/repository"
	"github.com/nathakusuma/bcc-be-freepass-2024/service"
	"github.com/nathakusuma/bcc-be-freepass-2024/util/roles"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error when loading .env file: " + err.Error())
	}

	db, err := database.MakeConnection(
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	if err != nil {
		log.Fatalln("Error when connecting to database: " + err.Error())
	}

	userRepo := repository.NewUserRepository(db)
	candidateRepo := repository.NewCandidateRepository(db, userRepo)
	periodRepo := repository.NewElectionPeriodRepository(db)
	commentRepo := repository.NewCommentRepository(db)
	postRepo := repository.NewPostRepository(db)

	userService := service.NewUserService(userRepo)
	candidateService := service.NewCandidateService(userRepo, candidateRepo)
	periodService := service.NewElectionPeriodService(periodRepo)
	voteService := service.NewVoteService(userRepo, candidateRepo, periodService)
	commentService := service.NewCommentService(commentRepo, postRepo, userRepo)
	postService := service.NewPostService(postRepo, candidateRepo, commentService)

	roleMid := middleware.NewRoleMiddleware(userRepo)

	userHandler := handler.NewUserHandler(userService)
	candidateHandler := handler.NewCandidateHandler(candidateService)
	periodHandler := handler.NewElectionPeriodHandler(periodService)
	voteHandler := handler.NewVoteHandler(voteService)
	postHandler := handler.NewPostHandler(postService)

	gin.SetMode(os.Getenv("GIN_MODE"))

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.POST("/register", userHandler.Register)
	v1.POST("/login", userHandler.Login)

	v1.GET("/users", middleware.Auth, userHandler.Get)
	v1.PATCH("/users", middleware.Auth, userHandler.Update)
	v1.DELETE("/users", middleware.Auth, roleMid.RequireRole(roles.Admin), userHandler.Delete)

	v1.GET("/candidates", middleware.Auth, candidateHandler.Get)
	v1.POST("/candidates", middleware.Auth, roleMid.RequireRole(roles.Admin), candidateHandler.Promote)

	v1.GET("/electionPeriod", middleware.Auth, periodHandler.Get)
	v1.POST("/electionPeriod", middleware.Auth, roleMid.RequireRole(roles.Admin), periodHandler.Set)

	v1.PUT("/votes", middleware.Auth, voteHandler.AddVote)

	v1.GET("/posts", middleware.Auth, postHandler.Get)
	v1.POST("/posts", middleware.Auth, roleMid.RequireRole(roles.Candidate), postHandler.Create)
	v1.PATCH("/posts", middleware.Auth, roleMid.RequireRole(roles.Candidate), postHandler.Update)
	v1.DELETE("/posts", middleware.Auth, roleMid.RequireRole(roles.Admin, roles.Candidate), postHandler.Delete)

	if err := router.Run(":" + os.Getenv("PORT")); err != nil {
		log.Fatalln(err)
	}
}
