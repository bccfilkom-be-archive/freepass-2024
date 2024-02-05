package usecase

import (
	"errors"
	"freepass-bcc/app/user/repository"
	"freepass-bcc/domain"
	"freepass-bcc/help"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	SignUp(userRequest domain.UserRequest) (domain.Users, any)
}

type UserUsecase struct {
	userRepository repository.IUserRepository
}

func NewUserUsecase(repository repository.IUserRepository) *UserUsecase {
	return &UserUsecase{repository}
}

func (u *UserUsecase) SignUp(userRequest domain.UserRequest) (domain.Users, any) {
	isUserExist := u.userRepository.GetUserByCondition(&domain.Users{}, "email = ?", userRequest.Email)
	if isUserExist == nil {
		return domain.Users{}, help.ErrorObject{
			Code: http.StatusBadRequest,
			Message: "email already used",
			Err: errors.New(""),
		}
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), 10)
	if err != nil {
		return domain.Users{}, help.ErrorObject{
			Code: http.StatusInternalServerError,
			Message: "error occured when hashing password",
			Err: err,
		}
	}

	user := domain.Users{
		Name:     userRequest.Name,
		Email:    userRequest.Email,
		Password: string(hashPassword),
		Role:     userRequest.Role,
	}

	err = u.userRepository.CreateUser(&user)
	if err != nil {
		return domain.Users{}, help.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "error occured when create user",
			Err:     err,
		}
	}

	return user, nil
}
