package electionController

import (
	"net/http"

	"github.com/gin-gonic/gin"

	candidateRepositorys "github.com/AkbarFikri/freepassBCC-2024/repositorys/candidate"
	electionRepositorys "github.com/AkbarFikri/freepassBCC-2024/repositorys/election"
	"github.com/AkbarFikri/freepassBCC-2024/schemas"
)

func GetAllCandidateListByElection(c *gin.Context) {
	electionID := c.Param("election_id")

	if electionID == "" {
		res := schemas.ResponeData{Error: true, Message: "id as a Path Param is required", Data: nil}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	election, err := electionRepositorys.FindSpecificElection(electionID)
	if err != nil {
		res := schemas.ResponeData{Error: true, Message: "Election with id provided is not found", Data: nil}
		c.JSON(http.StatusNotFound, res)
		return
	}

	candidates, err := candidateRepositorys.FindAllCandidateInElection(election.ID)
	if err != nil {
		res := schemas.ResponeData{Error: true, Message: "Failed to find candidates, something went wrong", Data: nil}
		c.JSON(http.StatusNotFound, res)
		return
	}

	if lenOfSlice := len(candidates); lenOfSlice == 0 {
		res := schemas.ResponeData{Error: true, Message: "No Data Found", Data: nil}
		c.JSON(http.StatusNotFound, res)
		return
	}

	res := schemas.ResponeData{Error: false, Message: "Find all posts by election_id successfully", Data: candidates}
	c.JSON(http.StatusOK, res)
	return
}
