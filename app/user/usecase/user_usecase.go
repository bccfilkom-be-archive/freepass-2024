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
	LoginUser(userLogin domain.UserLogin) (domain.Users, interface{}, any)
	PromoteUser(userId int) (domain.Users, any)
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
			Code:    http.StatusBadRequest,
			Message: "email already used",
			Err:     errors.New(""),
		}
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), 10)
	if err != nil {
		return domain.Users{}, help.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "error occured when hashing password",
			Err:     err,
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

func (u *UserUsecase) LoginUser(userLogin domain.UserLogin) (domain.Users, interface{}, any) {
	var user domain.Users
	err := u.userRepository.GetUserByCondition(&user, "email = ?", userLogin.Email)
	if err != nil {
		return domain.Users{}, "", help.ErrorObject{
			Code:    http.StatusNotFound,
			Message: "invalid username or password",
			Err:     err,
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLogin.Password))
	if err != nil {
		return domain.Users{}, "", help.ErrorObject{
			Code:    http.StatusNotFound,
			Message: "invalid username or password",
			Err:     err,
		}
	}

	tokenString, err := help.GenerateToken(user)
	if err != nil {
		return domain.Users{}, "", help.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "error occured when making jwt token",
			Err:     err,
		}
	}

	apiResponse := struct {
		Token string `json:"token"`
	}{
		tokenString,
	}

	return user, apiResponse, nil
}

func (u *UserUsecase) PromoteUser(userId int) (domain.Users, any){
	var user domain.Users
	err := u.userRepository.GetUserByCondition(&user, "id = ?", userId)
	if err != nil {
		return domain.Users{}, help.ErrorObject{
			Code: http.StatusNotFound,
			Message: "user not found",
			Err: err,
		}
	}

	if user.Role == "ADMIN" {
		return domain.Users{}, help.ErrorObject{
			Code: http.StatusBadRequest,
			Message: "admin can't be candidate",
			Err: errors.New("role not user"),
		}
	}

	if user.Role == "CANDIDATE" {
		return domain.Users{}, help.ErrorObject {
			Code: http.StatusBadRequest,
			Message: "already candidate",
			Err: errors.New("role not user"),
		}
	}
	
	user.Role = "CANDIDATE"

	err = u.userRepository.UpdateUser(&user)
	if err != nil{
		return domain.Users{}, help.ErrorObject{
			Code: http.StatusInternalServerError,
			Message: "error occured when update user",
			Err: err,
		}
	}

	return user, nil
}
