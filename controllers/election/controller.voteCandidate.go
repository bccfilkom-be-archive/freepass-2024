package electionController

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/AkbarFikri/freepassBCC-2024/models"
	candidateRepositorys "github.com/AkbarFikri/freepassBCC-2024/repositorys/candidate"
	electionRepositorys "github.com/AkbarFikri/freepassBCC-2024/repositorys/election"
	voteRepositorys "github.com/AkbarFikri/freepassBCC-2024/repositorys/vote"
	"github.com/AkbarFikri/freepassBCC-2024/schemas"

)

func VoteCandidate(c *gin.Context) {
	ElectionID := c.Param("election_id")
	CandidateID := c.Param("candidate_id")
	getUser, _ := c.Get("user")
	user := getUser.(schemas.UserTokenData)
	var data models.Vote

	if user.IsAdmin {
		res := schemas.ResponeData{Error: true, Message: "Admin is can't vote candidate, Use a user account to vote", Data: nil}
		c.JSON(http.StatusNotAcceptable, res)
		return
	}

	if ElectionID == "" || CandidateID == "" {
		res := schemas.ResponeData{Error: true, Message: "election_id, candidate_id as a Param is required", Data: nil}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	election, err := electionRepositorys.FindSpecificElection(ElectionID)
	if err != nil {
		res := schemas.ResponeData{Error: true, Message: "Election is not found, Something went wrong", Data: nil}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	startTime, err := time.Parse("2006-01-02T15:04:05-07:00", election.StartTime)
	if err != nil {
		res := schemas.ResponeData{Error: true, Message: "Wrong start_time format", Data: election}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	endTime, err := time.Parse("2006-01-02T15:04:05-07:00", election.EndTime)
	if err != nil {
		res := schemas.ResponeData{Error: true, Message: "Wrong end_time format", Data: nil}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	timeNow := time.Now().Local()
	if startTime.After(timeNow) || endTime.Before(timeNow) {
		res := schemas.ResponeData{Error: true, Message: "You can just vote during the election period", Data: nil}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	candidate, err := candidateRepositorys.FindSpecificCandidate(CandidateID)
	if err != nil {
		res := schemas.ResponeData{Error: true, Message: "Candidate is not found, Something went wrong", Data: nil}
		c.JSON(http.StatusNotFound, res)
		return
	}

	vote, err := voteRepositorys.FindUserVoteForCandidate(user.ID, candidate.ID)
	if err != nil {
		res := schemas.ResponeData{Error: true, Message: "Failed to vote, Something went wrong", Data: nil}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	if vote.ID == 0 {
		res := schemas.ResponeData{Error: true, Message: "You're already voted, Something went wrong", Data: nil}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	data.CandidateID = candidate.ID
	data.ElectionID = election.ID
	data.UserID = user.ID

	if err := voteRepositorys.CreateNewVote(&data); err != nil {
		res := schemas.ResponeData{Error: true, Message: "Failed to vote, Something went wrong", Data: nil}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	if err := candidateRepositorys.VoteCandidate(candidate.ID, candidate.Vote); err != nil {
		res := schemas.ResponeData{Error: true, Message: "Failed to vote, Something went wrong", Data: nil}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := schemas.ResponeData{Error: false, Message: "Successfully vote candidate", Data: data}
	c.JSON(http.StatusOK, res)
	return
}
