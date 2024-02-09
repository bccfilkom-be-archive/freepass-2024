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
	if election.ID == "" {
		res := schemas.ResponeData{Error: true, Message: "Election is not found", Data: nil}
		c.JSON(http.StatusNotFound, res)
		return
	}

	if err != nil {
		res := schemas.ResponeData{Error: true, Message: "Something went wrong", Data: nil}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	timeNow := time.Now().Local()
	if election.StartTime.After(timeNow) || election.EndTime.Before(timeNow) {
		res := schemas.ResponeData{Error: true, Message: "You can just vote during the election period", Data: nil}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	candidate, err := candidateRepositorys.FindSpecificCandidate(CandidateID)
	if candidate.ID == "" {
		res := schemas.ResponeData{Error: true, Message: "Candidate is not found, Something went wrong", Data: nil}
		c.JSON(http.StatusNotFound, res)
		return
	}

	if err != nil {
		res := schemas.ResponeData{Error: true, Message: "", Data: nil}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	var vote *models.Vote
	vote.ElectionID = election.ID
	vote.CandidateID = candidate.ID
	vote.UserID = user.ID

	if err := voteRepositorys.CreateNewVote(vote); err != nil {
		res := schemas.ResponeData{Error: true, Message: "Failed to vote, Something went wrong", Data: nil}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := schemas.ResponeData{Error: false, Message: "Successfully vote candidate", Data: vote}
	c.JSON(http.StatusOK, res)
	return
}
