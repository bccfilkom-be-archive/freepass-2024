package service

import (
	"bcc-be-freepass-2024/entity"
	"bcc-be-freepass-2024/model"
	"bcc-be-freepass-2024/repository"
	"bcc-be-freepass-2024/util/crypto"
	"bcc-be-freepass-2024/util/errortypes"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

type UserService struct {
	UserRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo}
}

func validateUsername(username string) error {
	validUsernamePattern := `^[a-z0-9]+$`

	regex, err := regexp.Compile(validUsernamePattern)
	if err != nil {
		return fmt.Errorf("regex compilation error: %v", err)
	}

	if !regex.MatchString(username) {
		return fmt.Errorf("invalid username: %s. Username must contain only lowercase letters and numbers", username)
	}

	return nil
}

func (service *UserService) Register(request *model.RegisterUserRequest) (*model.RegisterUserResponse, *errortypes.ApiError) {
	if err := validateUsername(request.Username); err != nil {
		return nil, &errortypes.ApiError{
			Code:    http.StatusBadRequest,
			Message: "invalid username",
			Data:    err,
		}
	}

	if request.Name == "" {
		request.Name = request.Username
	}

	hashedPassword, err := crypto.Hash(request.Password)
	if err != nil {
		return nil, &errortypes.ApiError{
			Code:    http.StatusInternalServerError,
			Message: "fail to hash password",
			Data:    err,
		}
	}

	user := entity.User{
		Name:     request.Name,
		Username: request.Username,
		Password: hashedPassword,
		Bio:      request.Bio,
	}

	if isUsernameExist, err := service.UserRepo.ExistsUsername(user.Username); isUsernameExist {
		if err != nil {
			return nil, &errortypes.ApiError{
				Code:    http.StatusInternalServerError,
				Message: "fail to check username data",
				Data:    err,
			}
		}
		return nil, &errortypes.ApiError{
			Code:    http.StatusConflict,
			Message: "username already exists",
			Data:    gin.H{},
		}
	}

	if err := service.UserRepo.Create(&user); err != nil {
		return nil, &errortypes.ApiError{
			Code:    http.StatusInternalServerError,
			Message: "fail to save user data",
			Data:    err,
		}
	}
	return &model.RegisterUserResponse{
		ID: user.ID,
	}, nil
}
