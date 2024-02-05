package handler

import (
	"freepass-bcc/app/user/usecase"
	"freepass-bcc/domain"
	"freepass-bcc/help"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUsecase usecase.IUserUsecase
}

func NewUserHandler(userUsecase usecase.IUserUsecase) *UserHandler {
	return &UserHandler{userUsecase}
}

func (h *UserHandler) SignUp(c *gin.Context) {
	var userRequest domain.UserRequest

	err := c.ShouldBindJSON(&userRequest)
	if err != nil {
		help.FailedResponse(c, http.StatusBadRequest, "failed to bind user request", err)
	}

	newUser, errorObject := h.userUsecase.SignUp(userRequest)

	if errorObject != nil {
		errorObject := errorObject.(help.ErrorObject)
		help.FailedResponse(c, errorObject.Code, errorObject.Message, errorObject.Err)
		return
	}

	help.SuccessResponse(c, "succes to create account", newUser)
}

func (h *UserHandler) LoginUser(c *gin.Context) {
	var userLogin domain.UserLogin

	err := c.ShouldBindJSON(&userLogin)
	if err != nil {
		return
	}

	user, apiResponse, errorObject := h.userUsecase.LoginUser(userLogin)
	if errorObject != nil {
		errorObject := errorObject.(help.ErrorObject)
		help.FailedResponse(c, errorObject.Code, errorObject.Message, errorObject.Err)
		return
	}

	help.SuccessResponse(c, "Welcome "+user.Name, apiResponse)
}

func (h *UserHandler) PromoteUser(c *gin.Context) {
	userIdString := c.Param("userId")
	userId, _ := strconv.Atoi(userIdString)

	promotedUser, errorObject := h.userUsecase.PromoteUser(userId)
	if errorObject != nil{
		errorObject := errorObject.(help.ErrorObject)
		help.FailedResponse(c, errorObject.Code, errorObject.Message, errorObject.Err)
		return
	}

	help.SuccessResponse(c, "succes promote user", promotedUser)
}
