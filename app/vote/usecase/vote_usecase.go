package usecase

import (
	"errors"
	election_time_repository "freepass-bcc/app/election_time/repository"
	user_repository "freepass-bcc/app/user/repository"
	vote_repository "freepass-bcc/app/vote/repository"
	"freepass-bcc/domain"
	"freepass-bcc/help"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type IVoteUsecase interface {
	Vote(c *gin.Context, userId int) (domain.Votes, any)
}

type VoteUsecase struct {
	voteRepository         vote_repository.IVoteRepository
	userRepository         user_repository.IUserRepository
	electionTimeRepository election_time_repository.IElectionTimeRepository
}

func NewVoteUsecase(voteRepository vote_repository.IVoteRepository, userRepository user_repository.IUserRepository,
	electionTimeRepository election_time_repository.IElectionTimeRepository) *VoteUsecase {
	return &VoteUsecase{voteRepository, userRepository, electionTimeRepository}
}

func (u *VoteUsecase) Vote(c *gin.Context, userId int) (domain.Votes, any) {
	loginUser, err := help.GetLoginUser(c)
	if err != nil {
		return domain.Votes{}, help.ErrorObject{
			Code:    http.StatusNotFound,
			Message: "account not found",
			Err:     err,
		}
	}

	var candidate domain.Users
	err = u.userRepository.GetUserByCondition(&candidate, "id = ?", userId)
	if err != nil {
		return domain.Votes{}, help.ErrorObject{
			Code:    http.StatusNotFound,
			Message: "candidate not found",
			Err:     err,
		}
	}

	if candidate.Role != "CANDIDATE" {
		return domain.Votes{}, help.ErrorObject{
			Code:    http.StatusBadRequest,
			Message: "vote declined",
			Err:     errors.New("only can vote candidate"),
		}
	}

	var electionTime domain.ElectionTimes
	err = u.electionTimeRepository.GetCurrentElectionTime(&electionTime)
	if err != nil {
		return domain.Votes{}, help.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "election time has not been set yet",
			Err:     err,
		}
	}

	currentTime := time.Now()

	if currentTime.Before(electionTime.StartTime) || currentTime.After(electionTime.EndTime){
		return domain.Votes{}, help.ErrorObject{
			Code:    http.StatusBadRequest,
			Message: "can't vote right now",
			Err:     errors.New("election not started yet"),
		}
	}

	var voted domain.Votes
	voted.UserId = loginUser.ID
	voted.Choice = candidate.Name
	voted.VoteTime = currentTime

	err = u.voteRepository.Vote(&voted)
	if err != nil {
		return domain.Votes{}, help.ErrorObject{
			Code: http.StatusInternalServerError,
			Message: "user only can vote 1 time",
			Err: err,
		}
	}

	return voted, nil
}
