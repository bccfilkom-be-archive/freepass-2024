package service

import (
	"github.com/nathakusuma/bcc-be-freepass-2024/model"
	"github.com/nathakusuma/bcc-be-freepass-2024/repository"
	"github.com/nathakusuma/bcc-be-freepass-2024/util/errortypes"
	"net/http"
)

type ElectionPeriodService struct {
	PeriodRepo *repository.ElectionPeriodRepository
}

func NewElectionPeriodService(periodRepo *repository.ElectionPeriodRepository) *ElectionPeriodService {
	return &ElectionPeriodService{periodRepo}
}

func (service *ElectionPeriodService) GetPeriod() (*model.GetElectionPeriodResponse, *errortypes.ApiError) {
	start, end, err := service.PeriodRepo.GetPeriod()
	if err != nil {
		return nil, &errortypes.ApiError{
			Code:    http.StatusInternalServerError,
			Message: "fail to get election period",
			Data:    err,
		}
	}
	return &model.GetElectionPeriodResponse{
		Start: start,
		End:   end,
	}, nil
}
