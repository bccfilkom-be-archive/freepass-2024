package electionController

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/AkbarFikri/freepassBCC-2024/models"
	candidateRepositorys "github.com/AkbarFikri/freepassBCC-2024/repositorys/candidate"
	electionRepositorys "github.com/AkbarFikri/freepassBCC-2024/repositorys/election"
	userRepositorys "github.com/AkbarFikri/freepassBCC-2024/repositorys/user"
	"github.com/AkbarFikri/freepassBCC-2024/schemas"
)

func PromoteCandidate(c *gin.Context) {
	electionID := c.Param("election_id")
	var data *models.Candidate

	if err := c.ShouldBindJSON(&data); err != nil {
		res := schemas.ResponeData{Error: true, Message: "Candidate data is Required", Data: nil}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if electionID == "" {
		res := schemas.ResponeData{Error: true, Message: "election_id as a Param is required", Data: nil}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if data.UserID == "" || data.ElectionNum == 0 {
		res := schemas.ResponeData{Error: true, Message: "user_id, election_number is Required", Data: nil}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	userCandidate, err := userRepositorys.FindOne(data.UserID)
	if userCandidate.ID == "" {
		res := schemas.ResponeData{Error: true, Message: "User is not found", Data: nil}
		c.JSON(http.StatusNotFound, res)
		return
	}

	if err != nil {
		res := schemas.ResponeData{Error: true, Message: "Something went wrong", Data: nil}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	election, err := electionRepositorys.FindSpecificElection(electionID)
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

	checkUser, _ := candidateRepositorys.FindCandidateInElection(userCandidate.ID, election.ID)
	if checkUser != nil {
		res := schemas.ResponeData{Error: true, Message: "This user is already become candidate in this election", Data: checkUser}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if err := electionRepositorys.FindElectionNumber(data.ElectionNum, election.ID); err == nil {
		res := schemas.ResponeData{Error: true, Message: "Election number is already used", Data: nil}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	data.ElectionID = election.ID
	data.UserID = userCandidate.ID

	createCandidate, err := candidateRepositorys.CreatCandidate(data)
	if err != nil {
		res := schemas.ResponeData{Error: true, Message: "Failed to promote user to candidate, Something went wrong", Data: nil}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	if err := candidateRepositorys.CreateCandidateInformation(createCandidate.ID); err != nil {
		res := schemas.ResponeData{Error: true, Message: "Failed to create candidate information, Something went wrong", Data: nil}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := schemas.ResponeData{Error: false, Message: "Successfully create candidate", Data: createCandidate}
	c.JSON(http.StatusOK, res)
	return
}
