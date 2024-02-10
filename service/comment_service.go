package service

import (
	"github.com/nathakusuma/bcc-be-freepass-2024/entity"
	"github.com/nathakusuma/bcc-be-freepass-2024/repository"
	"github.com/nathakusuma/bcc-be-freepass-2024/util/errortypes"
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
