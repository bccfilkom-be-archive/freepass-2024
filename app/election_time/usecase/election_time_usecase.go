package usecase

import (
	"freepass-bcc/app/election_time/repository"
	"freepass-bcc/domain"
	"freepass-bcc/help"
	"net/http"
	"time"
)

type IElectionTimeUsecase interface {
	SetStartAndEndTime(electionTime domain.ElectionTimesRequest) (domain.ElectionTimes, any)
}

type ElectionTimeUsecase struct {
	electionTimeRepository repository.IElectionTimeRepository
}

func NewElectionTimeUsecase(electionTimeRepository repository.IElectionTimeRepository) *ElectionTimeUsecase {
	return &ElectionTimeUsecase{electionTimeRepository}
}

func (u *ElectionTimeUsecase) SetStartAndEndTime(electionTime domain.ElectionTimesRequest) (domain.ElectionTimes, any) {
	startTime, err := time.Parse("2006-01-02", electionTime.StartTime)
	if err != nil {
		return domain.ElectionTimes{}, help.ErrorObject{
			Code: http.StatusBadRequest,
			Message: "time format must YYYY-MM-DD",
			Err: err,
		}
	}

	endTime, err := time.Parse("2006-01-02", electionTime.EndTime)
	if err != nil {
		return domain.ElectionTimes{}, help.ErrorObject{
			Code: http.StatusBadRequest,
			Message: "time format must YYYY-MM-DD",
			Err: err,
		}
	}

	var setElectionTime domain.ElectionTimes
	setElectionTime.StartTime = startTime
	setElectionTime.EndTime = endTime
	
	err = u.electionTimeRepository.SetStartAndEndTime(&setElectionTime)
	if err != nil {
		return domain.ElectionTimes{}, help.ErrorObject{
			Code: http.StatusInternalServerError,
			Message: "error occured when set start and end election time",
			Err: err,
		}
	}

	return setElectionTime, nil
}
