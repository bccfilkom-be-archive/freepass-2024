package usecase

import (
	"errors"
	comment_repository "freepass-bcc/app/comment/repository"
	post_repository "freepass-bcc/app/post/repository"
	user_repository "freepass-bcc/app/user/repository"
	vote_repository "freepass-bcc/app/vote/repository"
	"freepass-bcc/domain"
	"freepass-bcc/help"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	SignUp(userRequest domain.UserRequest) (domain.Users, any)
	LoginUser(userLogin domain.UserLogin) (domain.Users, interface{}, any)
	GetCandidates() ([]domain.Users, any)
	UpdateAccount(c *gin.Context, userRequest domain.UserRequest) (domain.Users, any)
	PromoteUser(userId int) (domain.Users, any)
	DeleteAccount(c *gin.Context, userId int) (domain.Users, any)
}

type UserUsecase struct {
	userRepository    user_repository.IUserRepository
	postRepository    post_repository.IPostRepository
	commentRepository comment_repository.ICommentRepository
	voteRepository    vote_repository.IVoteRepository
}

func NewUserUsecase(userRepository user_repository.IUserRepository, postRepository post_repository.IPostRepository, commentRepository comment_repository.ICommentRepository, voteRepository vote_repository.IVoteRepository) *UserUsecase {
	return &UserUsecase{userRepository, postRepository, commentRepository, voteRepository}
}

func (u *UserUsecase) SignUp(userRequest domain.UserRequest) (domain.Users, any) {
	isUserExist := u.userRepository.GetUserByCondition(&domain.Users{}, "email = ?", userRequest.Email)
	if isUserExist == nil {
		return domain.Users{}, help.ErrorObject{
			Code:    http.StatusConflict,
			Message: "can't use same email",
			Err:     errors.New("email already user"),
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
			Message: "error occured when create account",
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

func (u *UserUsecase) GetCandidates() ([]domain.Users, any) {
	var candidates []domain.Users
	err := u.userRepository.GetCandidates(&candidates)
	if err != nil {
		return []domain.Users{}, help.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "error occured when get all candidates information",
			Err:     err,
		}
	}

	return candidates, nil
}

func (u *UserUsecase) UpdateAccount(c *gin.Context, userRequest domain.UserRequest) (domain.Users, any) {
	loginUser, err := help.GetLoginUser(c)
	if err != nil {
		return domain.Users{}, help.ErrorObject{
			Code:    http.StatusNotFound,
			Message: "account not found",
			Err:     err,
		}
	}

	var user domain.Users
	err = u.userRepository.GetUserByCondition(&user, "id = ?", loginUser.ID)
	if err != nil {
		return domain.Users{}, help.ErrorObject{
			Code:    http.StatusNotFound,
			Message: "account not found",
			Err:     err,
		}
	}

	user.Name = userRequest.Name

	err = u.userRepository.UpdateAccount(&user)
	if err != nil {
		return domain.Users{}, help.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "error occured when update user",
			Err:     err,
		}
	}

	return user, nil
}

func (u *UserUsecase) PromoteUser(userId int) (domain.Users, any) {
	var user domain.Users
	err := u.userRepository.GetUserByCondition(&user, "id = ?", userId)
	if err != nil {
		return domain.Users{}, help.ErrorObject{
			Code:    http.StatusNotFound,
			Message: "account not found",
			Err:     err,
		}
	}

	if user.Role == "ADMIN" {
		return domain.Users{}, help.ErrorObject{
			Code:    http.StatusBadRequest,
			Message: "admin can't be candidate",
			Err:     errors.New("role not user"),
		}
	}

	if user.Role == "CANDIDATE" {
		return domain.Users{}, help.ErrorObject{
			Code:    http.StatusBadRequest,
			Message: "already candidate",
			Err:     errors.New("role not user"),
		}
	}

	user.Role = "CANDIDATE"

	err = u.userRepository.UpdateAccount(&user)
	if err != nil {
		return domain.Users{}, help.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "error occured when update user",
			Err:     err,
		}
	}

	return user, nil
}

func (u *UserUsecase) DeleteAccount(c *gin.Context, userId int) (domain.Users, any) {
	loginUser, err := help.GetLoginUser(c)
	if err != nil {
		return domain.Users{}, help.ErrorObject{
			Code:    http.StatusNotFound,
			Message: "account not found",
			Err:     err,
		}
	}

	var user domain.Users
	err = u.userRepository.GetUserByCondition(&user, "id = ?", userId)
	if err != nil {
		return domain.Users{}, help.ErrorObject{
			Code:    http.StatusNotFound,
			Message: "account not found",
			Err:     err,
		}
	}

	if user.Role == "ADMIN" && loginUser.ID != userId {
		return domain.Users{}, help.ErrorObject{
			Code:    http.StatusBadRequest,
			Message: "admin can't delete other admin account",
			Err:     errors.New("acces denied"),
		}
	}

	if user.Role == "CANDIDATE" {
		var posts []domain.Posts
		err := u.postRepository.GetAllPostByUserId(&posts, userId)
		if err != nil {
			return domain.Users{}, help.ErrorObject{
				Code:    http.StatusInternalServerError,
				Message: "error occured when get all post by user id",
				Err:     err,
			}
		}

		if len(posts) != 0 {
			for _, p := range posts {
				var comments []domain.Comments
				err := u.commentRepository.GetAllCommentByCondition(&comments, "post_id = ?", p.ID)
				if err != nil {
					return domain.Users{}, help.ErrorObject{
						Code:    http.StatusInternalServerError,
						Message: "error occured when get all comment by post id",
						Err:     err,
					}
				}

				if len(comments) != 0 {
					err = u.commentRepository.DeleteAllComment(&comments)
					if err != nil {
						return domain.Users{}, help.ErrorObject{
							Code:    http.StatusInternalServerError,
							Message: "error occured when delete all comments",
							Err:     err,
						}
					}
				}
			}

			err := u.postRepository.DeleteAllPost(&posts)
			if err != nil {
				return domain.Users{}, help.ErrorObject{
					Code:    http.StatusInternalServerError,
					Message: "error occured when deleteing all user posts",
					Err:     err,
				}
			}
		}
	}

	if user.Role == "USER" {
		var voted domain.Votes
		err = u.voteRepository.GetVoteByCondition(&voted, "user_id = ?", userId)
		if err == nil {
			err = u.voteRepository.DeleteVote(&voted)
			if err != nil {
				return domain.Users{}, help.ErrorObject{
					Code:    http.StatusInternalServerError,
					Message: "error occured when delete vote",
					Err:     err,
				}
			}
		}

		var comments []domain.Comments
		err = u.commentRepository.GetAllCommentByCondition(&comments, "user_id = ?", userId)
		if err != nil {
			return domain.Users{}, help.ErrorObject{
				Code:    http.StatusInternalServerError,
				Message: "error occured when get all user comments",
				Err:     err,
			}
		}

		if len(comments) != 0 {
			err = u.commentRepository.DeleteAllComment(&comments)
			if err != nil {
				return domain.Users{}, help.ErrorObject{
					Code:    http.StatusInternalServerError,
					Message: "error occured when get all user comments",
					Err:     err,
				}
			}
		}
	}
	err = u.userRepository.DeleteAccount(&user)
	if err != nil {
		return domain.Users{}, help.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "error occured when deleting account",
			Err:     err,
		}
	}

	return user, nil
}
