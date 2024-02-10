package handler

import (
	"bcc-be-freepass-2024/model"
	"bcc-be-freepass-2024/service"
	"bcc-be-freepass-2024/util/apiresponse"
	"bcc-be-freepass-2024/util/binding"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService}
}

func (handler *UserHandler) Register(ctx *gin.Context) {
	var request model.RegisterUserRequest
	if err := binding.ShouldBindJSON(ctx, &request); err != nil {
		apiresponse.ApiError(ctx, err)
		return
	}

	response, err := handler.UserService.Register(&request)
	if err != nil {
		apiresponse.ApiError(ctx, err)
		return
	}

	apiresponse.Success(ctx, http.StatusCreated, "user successfully registered", response)
}
