package service

import (
	"errors"
	"github.com/nathakusuma/bcc-be-freepass-2024/entity"
	"github.com/nathakusuma/bcc-be-freepass-2024/model"
	"github.com/nathakusuma/bcc-be-freepass-2024/repository"
	"github.com/nathakusuma/bcc-be-freepass-2024/util/errortypes"
	"gorm.io/gorm"
	"net/http"
)

type CommentService struct {
	CommentRepo *repository.CommentRepository
	PostRepo    *repository.PostRepository
	UserRepo    *repository.UserRepository
}

func NewCommentService(commRepo *repository.CommentRepository, postRepo *repository.PostRepository, userRepo *repository.UserRepository) *CommentService {
	return &CommentService{commRepo, postRepo, userRepo}
}

func (service *CommentService) GetByPostId(postId uint) ([]entity.Comment, *errortypes.ApiError) {
	comments, err := service.CommentRepo.FindByPostId(postId)
	if err != nil {
		return nil, &errortypes.ApiError{
			Code:    http.StatusInternalServerError,
			Message: "error when finding comments",
			Data:    err,
		}
	}
	return comments, nil
}

func (service *CommentService) Add(request *model.AddCommentRequest, postId uint, userId uint) (*model.AddCommentResponse, *errortypes.ApiError) {
	post, err := service.PostRepo.FindById(postId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errortypes.ApiError{
				Code:    http.StatusNotFound,
				Message: "post not found",
				Data:    err,
			}
		}
		return nil, &errortypes.ApiError{
			Code:    http.StatusInternalServerError,
			Message: "fail to update user data",
			Data:    err,
		}
	}

	user, _ := service.UserRepo.FindById(userId)

	comment := entity.Comment{
		Content: request.Content,
		User:    *user,
		Post:    *post,
	}
	if err := service.CommentRepo.Add(&comment); err != nil {
		return nil, &errortypes.ApiError{
			Code:    http.StatusInternalServerError,
			Message: "fail to save comment",
			Data:    err,
		}
	}

	return &model.AddCommentResponse{ID: comment.ID}, nil
}
