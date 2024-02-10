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

type PostService struct {
	PostRepo       *repository.PostRepository
	CandidateRepo  *repository.CandidateRepository
	CommentService *CommentService
}

func NewPostService(postRepo *repository.PostRepository, candRepo *repository.CandidateRepository, commServ *CommentService) *PostService {
	return &PostService{postRepo, candRepo, commServ}
}

func (service *PostService) GetById(postId uint) (*model.GetPostResponse, *errortypes.ApiError) {
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
			Message: "error when finding post",
			Data:    err,
		}
	}

	comments, err2 := service.CommentService.GetByPostId(post.ID)
	if err2 != nil {
		return nil, &errortypes.ApiError{
			Code:    http.StatusInternalServerError,
			Message: "error when finding comments",
			Data:    err2,
		}
	}

	var commentResponses []model.GetCommentResponse
	for _, comment := range comments {
		commentResponses = append(commentResponses, model.GetCommentResponse{
			ID:        comment.ID,
			Content:   comment.Content,
			UserID:    comment.UserID,
			CreatedAt: comment.CreatedAt,
		})
	}

	return &model.GetPostResponse{
		ID:          post.ID,
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
		Title:       post.Title,
		Content:     post.Content,
		CandidateID: post.CandidateID,
		Comments:    commentResponses,
	}, nil
}

func (service *PostService) GetByCandidateId(candidateId uint) ([]model.GetPostResponse, *errortypes.ApiError) {
	posts, err := service.PostRepo.FindByCandidateId(candidateId)
	if err != nil {
		return nil, &errortypes.ApiError{
			Code:    http.StatusInternalServerError,
			Message: "error when finding post",
			Data:    err,
		}
	}

	var postResponses []model.GetPostResponse
	for _, post := range posts {
		comments, err := service.CommentService.GetByPostId(post.ID)
		if err != nil {
			return nil, &errortypes.ApiError{
				Code:    http.StatusInternalServerError,
				Message: "error when finding comments",
				Data:    err,
			}
		}

		var commentResponses []model.GetCommentResponse
		for _, comment := range comments {
			commentResponses = append(commentResponses, model.GetCommentResponse{
				ID:        comment.ID,
				Content:   comment.Content,
				UserID:    comment.UserID,
				CreatedAt: comment.CreatedAt,
			})
		}

		postResponses = append(postResponses, model.GetPostResponse{
			ID:          post.ID,
			CreatedAt:   post.CreatedAt,
			UpdatedAt:   post.UpdatedAt,
			Title:       post.Title,
			Content:     post.Content,
			CandidateID: post.CandidateID,
			Comments:    commentResponses,
		})
	}

	if postResponses == nil {
		postResponses = make([]model.GetPostResponse, 0)
	}

	return postResponses, nil
}

func (service *PostService) Create(request *model.CreatePostRequest, candidateId uint) (uint, *errortypes.ApiError) {
	candidate, _ := service.CandidateRepo.FindById(candidateId)
	post := entity.Post{
		Title:       request.Title,
		Content:     request.Content,
		CandidateID: candidateId,
		Candidate:   *candidate,
	}
	err := service.PostRepo.Create(&post)
	if err != nil {
		return 0, &errortypes.ApiError{
			Code:    http.StatusInternalServerError,
			Message: "fail to save post",
			Data:    err,
		}
	}

	return post.ID, nil
}
