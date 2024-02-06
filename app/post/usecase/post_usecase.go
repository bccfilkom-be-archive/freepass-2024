package usecase

import (
	"freepass-bcc/app/post/repository"
	// userRepository"freepass-bcc/app/user/repository"
	"freepass-bcc/domain"
	"freepass-bcc/help"

	"github.com/gin-gonic/gin"
)

type IPostUsecase interface {
	CreatePost(c *gin.Context, postRequest domain.PostRequest) (domain.Posts, any)
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
		return domain.Posts{}, err
	}
	
	var post domain.Posts
	post.UserID = loginUser.ID
	post.Post = postRequest.Post

	err = u.postRepository.CreatePost(&post)
	if err != nil {
		return domain.Posts{}, err
	}

	return post, nil
}
