package handler

import (
	"bcc-be-freepass-2024/model"
	"bcc-be-freepass-2024/service"
	"bcc-be-freepass-2024/util/apiresponse"
	"bcc-be-freepass-2024/util/binding"
	"bcc-be-freepass-2024/util/errortypes"
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

func (handler *UserHandler) Login(ctx *gin.Context) {
	var request model.LoginUserRequest
	if err := binding.ShouldBindJSON(ctx, &request); err != nil {
		apiresponse.ApiError(ctx, err)
		return
	}

	response, err := handler.UserService.Login(&request)
	if err != nil {
		apiresponse.ApiError(ctx, err)
		return
	}

	apiresponse.Success(ctx, http.StatusOK, "successfully logged in", response)
}

func (handler *UserHandler) Get(ctx *gin.Context) {
	var response *model.GetUserResponse
	var apiErr *errortypes.ApiError
	if userId, err := binding.ShouldUintQuery(ctx, "userId"); err == nil {
		response, apiErr = handler.UserService.GetById(userId)
	} else if username, err := binding.ShouldQueryExist(ctx, "username"); err == nil {
		response, apiErr = handler.UserService.GetByUsername(username)
	} else {
		apiresponse.ApiError(ctx, &errortypes.ApiError{
			Code:    http.StatusBadRequest,
			Message: "invalid/missing query",
			Data:    gin.H{},
		})
		return
	}

	if apiErr != nil {
		apiresponse.ApiError(ctx, apiErr)
		return
	}

	apiresponse.Success(ctx, http.StatusOK, "successfully retrieved user data", response)
}
