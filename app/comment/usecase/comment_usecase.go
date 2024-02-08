package usecase

import (
	"freepass-bcc/app/comment/repository"
	post_repository "freepass-bcc/app/post/repository"
	"freepass-bcc/domain"
	"freepass-bcc/help"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ICommentUsecase interface {
	CreateComment(c *gin.Context, commentRequest domain.CommentRequest, postId int) (domain.PostResponse, []domain.CommentResponse, any)
	DeleteComment(postId int, commentId int) (domain.PostResponse, []domain.CommentResponse, any)
}

type CommentUsecase struct {
	commentRepository repository.ICommentRepository
	postRepository    post_repository.IPostRepository
}

func NewCommentUsecase(commentRepository repository.ICommentRepository, postRepository post_repository.IPostRepository) *CommentUsecase {
	return &CommentUsecase{commentRepository, postRepository}
}

func (u *CommentUsecase) CreateComment(c *gin.Context, commentRequest domain.CommentRequest, postId int) (domain.PostResponse, []domain.CommentResponse, any) {
	loginUser, err := help.GetLoginUser(c)
	if err != nil {
		return domain.PostResponse{}, []domain.CommentResponse{}, help.ErrorObject{
			Code:    http.StatusNotFound,
			Message: "user not found",
			Err:     err,
		}
	}

	var post domain.Posts
	err = u.postRepository.GetPostByCondition(&post, "id = ?", postId)
	if err != nil {
		return domain.PostResponse{}, []domain.CommentResponse{}, help.ErrorObject{
			Code:    http.StatusNotFound,
			Message: "post not found",
			Err:     err,
		}
	}

	var comment domain.Comments
	comment.UserID = loginUser.ID
	comment.PostID = postId
	comment.Comment = commentRequest.Comment

	err = u.commentRepository.CreateComment(&comment)
	if err != nil {
		return domain.PostResponse{}, []domain.CommentResponse{}, help.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "error occured when create comment",
			Err:     err,
		}
	}

	postResponse := help.PostResponse(post, "")

	var comments []domain.Comments
	err = u.commentRepository.GetAllComment(&comments, postId)
	if err != nil {
		return domain.PostResponse{}, []domain.CommentResponse{}, help.ErrorObject{
			Code:    http.StatusNotFound,
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

func (u *CommentUsecase) DeleteComment(postId int, commentId int) (domain.PostResponse, []domain.CommentResponse, any) {
	var post domain.Posts
	err := u.postRepository.GetPostByCondition(&post, "id = ?", postId)
	if err != nil {
		return domain.PostResponse{}, []domain.CommentResponse{}, help.ErrorObject{
			Code:    http.StatusNotFound,
			Message: "post not found",
			Err:     err,
		}
	}

	var comment domain.Comments
	err = u.commentRepository.GetCommentByCondition(&comment, "id = ?", commentId)
	if err != nil {
		return domain.PostResponse{}, []domain.CommentResponse{}, help.ErrorObject{
			Code:    http.StatusNotFound,
			Message: "comment not found",
			Err:     err,
		}
	}

	err = u.commentRepository.DeleteComment(&comment)
	if err != nil {
		return domain.PostResponse{}, []domain.CommentResponse{}, help.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "error occured when delete comment",
			Err:     err,
		}
	}

	postResponse := help.PostResponse(post, "")

	var comments []domain.Comments
	err = u.commentRepository.GetAllComment(&comments, postId)
	if err != nil {
		return domain.PostResponse{}, []domain.CommentResponse{}, help.ErrorObject{
			Code:    http.StatusNotFound,
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
