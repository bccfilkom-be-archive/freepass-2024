package usecase

import (
	"errors"
	comment_repository "freepass-bcc/app/comment/repository"
	post_repository "freepass-bcc/app/post/repository"
	"net/http"

	"freepass-bcc/domain"
	"freepass-bcc/help"

	"github.com/gin-gonic/gin"
)

type IPostUsecase interface {
	GetAllPost() ([]domain.PostResponse, any)
	GetPost(postId int) (domain.PostResponse, []domain.CommentResponse, any)
	CreatePost(c *gin.Context, postRequest domain.PostRequest) (domain.PostResponse, any)
	UpdatePost(c *gin.Context, postRequest domain.PostRequest, postId int) (domain.PostResponse, any)
	DeletePost(c *gin.Context, postId int) (domain.PostResponse, any)
}

type PostUsecase struct {
	postRepository    post_repository.IPostRepository
	commentRepository comment_repository.ICommentRepository
}

func NewPostUsecase(postRepository post_repository.IPostRepository, commentRepository comment_repository.ICommentRepository) *PostUsecase {
	return &PostUsecase{postRepository, commentRepository}
}

func (u *PostUsecase) GetAllPost() ([]domain.PostResponse, any) {
	var posts []domain.Posts
	err := u.postRepository.GetAllPost(&posts)
	if err != nil {
		return []domain.PostResponse{}, help.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "error occured when get all posts",
			Err:     err,
		}
	}

	var postResponses []domain.PostResponse
	for _, p := range posts {
		postResponse := help.PostResponse(p, "")

		postResponses = append(postResponses, postResponse)
	}

	return postResponses, nil
}

func (u *PostUsecase) GetPost(postId int) (domain.PostResponse, []domain.CommentResponse, any) {
	var post domain.Posts
	err := u.postRepository.GetPostByCondition(&post, "id = ?", postId)
	if err != nil {
		return domain.PostResponse{}, []domain.CommentResponse{}, help.ErrorObject{
			Code:    http.StatusNotFound,
			Message: "post not found",
			Err:     err,
		}
	}

	postResponse := help.PostResponse(post, "")

	var comments []domain.Comments
	err = u.commentRepository.GetAllCommentByPostID(&comments, postId)
	if err != nil {
		return domain.PostResponse{}, []domain.CommentResponse{}, help.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "failed to get all comments",
			Err:     err,
		}
	}

	var commentResponses []domain.CommentResponse

	for _, c := range comments {
		commentResponse := help.CommentResponse(c)

		commentResponses = append(commentResponses, commentResponse)
	}

	return postResponse, commentResponses, nil
}

func (u *PostUsecase) CreatePost(c *gin.Context, postRequest domain.PostRequest) (domain.PostResponse, any) {
	loginUser, err := help.GetLoginUser(c)
	if err != nil {
		return domain.PostResponse{}, help.ErrorObject{
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
		return domain.PostResponse{}, help.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "error occured when create post",
			Err:     err,
		}
	}

	postResponse := help.PostResponse(post, loginUser.Name)

	return postResponse, nil
}

func (u *PostUsecase) UpdatePost(c *gin.Context, postRequest domain.PostRequest, postId int) (domain.PostResponse, any) {
	loginUser, err := help.GetLoginUser(c)
	if err != nil {
		return domain.PostResponse{}, help.ErrorObject{
			Code:    http.StatusNotFound,
			Message: "account not found",
			Err:     err,
		}
	}

	var post domain.Posts
	err = u.postRepository.GetPostByCondition(&post, "id = ?", postId)
	if err != nil {
		return domain.PostResponse{}, help.ErrorObject{
			Code:    http.StatusNotFound,
			Message: "post not found",
			Err:     err,
		}
	}

	if loginUser.ID != post.UserID {
		return domain.PostResponse{}, help.ErrorObject{
			Code:    http.StatusBadRequest,
			Message: "can't edit other candidate post",
			Err:     errors.New("access denied"),
		}
	}

	post.Post = postRequest.Post
	err = u.postRepository.UpdatePost(&post)
	if err != nil {
		return domain.PostResponse{}, help.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "error occured when update post",
			Err:     err,
		}
	}

	postResponse := help.PostResponse(post, "")

	return postResponse, nil
}

func (u *PostUsecase) DeletePost(c *gin.Context, postId int) (domain.PostResponse, any) {
	loginUser, err := help.GetLoginUser(c)
	if err != nil {
		return domain.PostResponse{}, help.ErrorObject{
			Code:    http.StatusNotFound,
			Message: "account not found",
			Err:     err,
		}
	}

	var post domain.Posts
	err = u.postRepository.GetPostByCondition(&post, "id = ?", postId)
	if err != nil {
		return domain.PostResponse{}, help.ErrorObject{
			Code:    http.StatusNotFound,
			Message: "post not found",
			Err:     err,
		}
	}

	var comments []domain.Comments
	err = u.commentRepository.GetAllCommentByPostID(&comments, postId)
	if err != nil {
		return domain.PostResponse{}, help.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "failed to get all comments",
			Err:     err,
		}
	}

	if loginUser.Role != "ADMIN" && loginUser.ID != post.UserID {
		return domain.PostResponse{}, help.ErrorObject{
			Code:    http.StatusBadRequest,
			Message: "can't delete other candidate post",
			Err:     errors.New("access denied"),
		}
	}

	if len(comments) != 0 {
		err = u.commentRepository.DeleteAllComment(&comments)
		if err != nil {
			return domain.PostResponse{}, help.ErrorObject{
				Code:    http.StatusInternalServerError,
				Message: "error occured when delete all comments",
				Err:     err,
			}
		}
	}

	err = u.postRepository.DeletePost(&post)
	if err != nil {
		return domain.PostResponse{}, help.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "error occured when delete post",
			Err:     err,
		}
	}

	postResponse := help.PostResponse(post, "")

	return postResponse, nil
}
