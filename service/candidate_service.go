package service

import (
	"github.com/nathakusuma/bcc-be-freepass-2024/model"
	"github.com/nathakusuma/bcc-be-freepass-2024/repository"
	"github.com/nathakusuma/bcc-be-freepass-2024/util/errortypes"
	"net/http"
)

type CandidateService struct {
	UserRepo      *repository.UserRepository
	CandidateRepo *repository.CandidateRepository
}

func NewCandidateService(userRepo *repository.UserRepository, candidateRepo *repository.CandidateRepository) *CandidateService {
	return &CandidateService{userRepo, candidateRepo}
}

func (service *CandidateService) GetAll() ([]model.GetCandidateResponse, *errortypes.ApiError) {
	candidates, err := service.CandidateRepo.FindAll()
	if err != nil {
		return nil, &errortypes.ApiError{
			Code:    http.StatusInternalServerError,
			Message: "fail to get candidates data",
			Data:    err,
		}
	}

	var response []model.GetCandidateResponse
	for _, candidate := range candidates {
		response = append(response, model.GetCandidateResponse{
			ID:     candidate.ID,
			UserID: candidate.UserID,
			Votes:  candidate.VoteCount,
		})
	}

	return response, nil
}
