package usecase

import (
	"errors"
	"freepass-bcc/app/post/repository"
	"net/http"

	// userRepository"freepass-bcc/app/user/repository"
	"freepass-bcc/domain"
	"freepass-bcc/help"

	"github.com/gin-gonic/gin"
)

type IPostUsecase interface {
	CreatePost(c *gin.Context, postRequest domain.PostRequest) (domain.Posts, any)
	UpdatePost(c *gin.Context, postRequest domain.PostRequest, postId int) (domain.Posts, any)
}

type PostUsecase struct {
	postRepository repository.IPostRepository
	// userRepository userRepository.IUserRepository
}

func NewPostUsecase(repository repository.IPostRepository) *PostUsecase {
	return &PostUsecase{repository}
}

func (u *PostUsecase) CreatePost(c *gin.Context, postRequest domain.PostRequest) (domain.Posts, any) {
	loginUser, err := help.GetLoginUser(c)
	if err != nil {
		return domain.Posts{}, help.ErrorObject{
			Code:    http.StatusNotFound,
			Message: "account not found",
			Err:     err,
		}
	}
	
	var post domain.Posts
	post.UserID = loginUser.ID
	post.Post = postRequest.Post

	err = u.postRepository.CreatePost(&post)
	if err != nil {
		return domain.Posts{}, help.ErrorObject{
			Code: http.StatusInternalServerError,
			Message: "error occured when create post",
			Err: err,
		}
	}

	return post, nil
}

func (u *PostUsecase) UpdatePost(c *gin.Context, postRequest domain.PostRequest, postId int) (domain.Posts, any) {
	loginUser, err := help.GetLoginUser(c)
		if err != nil{
			return domain.Posts{}, help.ErrorObject{
				Code:    http.StatusNotFound,
				Message: "account not found",
				Err:     err,
			}
		}
	

	var post domain.Posts
	err = u.postRepository.GetPostByCondition(&post, "id = ?", postId)
	if err != nil {
		return domain.Posts{}, help.ErrorObject{
			Code: http.StatusInternalServerError,
			Message: "error occured when get post id",
			Err: err,
		}
	}

	if loginUser.ID != post.UserID {
		return domain.Posts{}, help.ErrorObject{
			Code: http.StatusBadRequest,
			Message: "can't edit other candidate post",
			Err: errors.New("access denied"),
		}
	}

	post.Post = postRequest.Post
	err = u.postRepository.UpdatePost(&post)
	if err != nil {
		return domain.Posts{}, help.ErrorObject{
			Code: http.StatusInternalServerError,
			Message: "error occured when update post",
			Err: err,
		}
	}

	return post, nil
}
