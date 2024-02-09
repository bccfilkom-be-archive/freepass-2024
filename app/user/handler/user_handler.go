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
		return
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
		help.FailedResponse(c, http.StatusBadRequest, "failed to bind user request", err)
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

func (h *UserHandler) GetCandidates(c *gin.Context) {
	candidates, errorObject := h.userUsecase.GetCandidates()
	if errorObject != nil {
		errorObject := errorObject.(help.ErrorObject)
		help.FailedResponse(c, errorObject.Code, errorObject.Message, errorObject.Err)
		return
	}

	help.SuccessResponse(c, "success get all candidates information", candidates)
}

func (h *UserHandler) UpdateAccount(c *gin.Context) {
	var userRequest domain.UserRequest
	err := c.ShouldBindJSON(&userRequest)
	if err != nil {
		help.FailedResponse(c, http.StatusBadRequest, "failed to bind user request", err)
		return
	}

	updatedUser, errorObject := h.userUsecase.UpdateAccount(c, userRequest)
	if errorObject != nil {
		errorObject := errorObject.(help.ErrorObject)
		help.FailedResponse(c, errorObject.Code, errorObject.Message, errorObject.Err)
		return
	}

	help.SuccessResponse(c, "success update account", updatedUser)
}

func (h *UserHandler) PromoteUser(c *gin.Context) {
	userIdString := c.Param("userId")
	userId, _ := strconv.Atoi(userIdString)

	promotedUser, errorObject := h.userUsecase.PromoteUser(userId)
	if errorObject != nil {
		errorObject := errorObject.(help.ErrorObject)
		help.FailedResponse(c, errorObject.Code, errorObject.Message, errorObject.Err)
		return
	}

	help.SuccessResponse(c, "success promote user", promotedUser)
}

func (h *UserHandler) DeleteAccount(c *gin.Context) {
	userIdString := c.Param("userId")
	userId, _ := strconv.Atoi(userIdString)

	deletedAccount, errorObject := h.userUsecase.DeleteAccount(c, userId)
	if errorObject != nil {
		errorObject := errorObject.(help.ErrorObject)
		help.FailedResponse(c, errorObject.Code, errorObject.Message, errorObject.Err)
		return
	}

	help.SuccessResponse(c, "success delete account", deletedAccount)
}
