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
	CreateComment(c *gin.Context, commentRequest domain.CommentRequest, postId int) (domain.Comments, any)
}

type CommentUsecase struct {
	commentRepository repository.ICommentRepository
	postRepository    post_repository.IPostRepository
}

func NewCommentUsecase(commentRepository repository.ICommentRepository, postRepository post_repository.IPostRepository) *CommentUsecase {
	return &CommentUsecase{commentRepository, postRepository}
}

func (u *CommentUsecase) CreateComment(c *gin.Context, commentRequest domain.CommentRequest, postId int) (domain.Comments, any) {
	loginUser, err := help.GetLoginUser(c)
	if err != nil {
		return domain.Comments{}, help.ErrorObject{
			Code:    http.StatusNotFound,
			Message: "user not found",
			Err:     err,
		}
	}

	err = u.postRepository.GetPostByCondition(&domain.Posts{}, "id = ?", postId)
	if err != nil {
		return domain.Comments{}, help.ErrorObject{
			Code: http.StatusNotFound,
			Message: "post not found",
			Err: err,
		}
	}

	var comment domain.Comments
	comment.UserID = loginUser.ID
	comment.PostID = postId
	comment.Comment = commentRequest.Comment

	err = u.commentRepository.CreateComment(&comment)
	if err != nil {
		return domain.Comments{}, help.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "error occured when create comment",
			Err:     err,
		}
	}

	return comment, nil
}
