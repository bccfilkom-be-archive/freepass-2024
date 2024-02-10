package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/nathakusuma/bcc-be-freepass-2024/repository"
	"github.com/nathakusuma/bcc-be-freepass-2024/util/errortypes"
	"gorm.io/gorm"
	"net/http"
)

type VoteService struct {
	UserRepo      *repository.UserRepository
	CandidateRepo *repository.CandidateRepository
	PeriodService *ElectionPeriodService
}

func NewVoteService(userRepo *repository.UserRepository, candidateRepo *repository.CandidateRepository,
	periodServ *ElectionPeriodService) *VoteService {

	return &VoteService{userRepo, candidateRepo, periodServ}
}

func (service *VoteService) AddVote(candidateId uint, voterId uint) *errortypes.ApiError {
	isPeriod, err := service.PeriodService.IsInPeriod()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &errortypes.ApiError{
				Code:    http.StatusNotFound,
				Message: "election period is not set",
				Data:    err,
			}
		}
		return &errortypes.ApiError{
			Code:    http.StatusInternalServerError,
			Message: "fail to check election period data",
			Data:    err,
		}
	}
	if !isPeriod {
		return &errortypes.ApiError{
			Code:    http.StatusForbidden,
			Message: "not in election period",
			Data:    gin.H{},
		}
	}

	candidate, err := service.CandidateRepo.FindById(candidateId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &errortypes.ApiError{
				Code:    http.StatusNotFound,
				Message: "no candidate with such id",
				Data:    err,
			}
		}
		return &errortypes.ApiError{
			Code:    http.StatusInternalServerError,
			Message: "fail to check candidate data",
			Data:    err,
		}
	}

	voter, err := service.UserRepo.FindById(voterId)
	if err != nil {
		return &errortypes.ApiError{
			Code:    http.StatusInternalServerError,
			Message: "fail to check voter data",
			Data:    err,
		}
	}
	if !voter.CanVote {
		return &errortypes.ApiError{
			Code:    http.StatusForbidden,
			Message: "user not allowed to vote",
			Data:    gin.H{},
		}
	}

	if err := service.CandidateRepo.AddVote(candidate, voter); err != nil {
		return &errortypes.ApiError{
			Code:    http.StatusInternalServerError,
			Message: "fail to add vote",
			Data:    err,
		}
	}
	return nil
}
