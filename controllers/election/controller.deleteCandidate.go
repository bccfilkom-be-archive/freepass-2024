package electionController

import (
	"net/http"

	"github.com/gin-gonic/gin"

	candidateRepositorys "github.com/AkbarFikri/freepassBCC-2024/repositorys/candidate"
	postRepositorys "github.com/AkbarFikri/freepassBCC-2024/repositorys/post"
	voteRepositorys "github.com/AkbarFikri/freepassBCC-2024/repositorys/vote"
	"github.com/AkbarFikri/freepassBCC-2024/schemas"

)

func DeleteCandidate(c *gin.Context) {
	candidateID := c.Param("id")
	getUser, _ := c.Get("user")
	user := getUser.(schemas.UserTokenData)

	if candidateID == "" {
		res := schemas.ResponeData{Error: true, Message: "id as a Path Param is required", Data: nil}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	candidate, err := candidateRepositorys.FindSpecificCandidate(candidateID)
	if err != nil {
		res := schemas.ResponeData{Error: true, Message: "Candidate with id provided is not found", Data: nil}
		c.JSON(http.StatusNotFound, res)
		return
	}

	if !user.IsAdmin {
		res := schemas.ResponeData{Error: true, Message: "You're not allowed to delete this data", Data: nil}
		c.JSON(http.StatusNotFound, res)
		return
	}

	if err := postRepositorys.DeleteCandidatePost(candidate.ID); err != nil {
		res := schemas.ResponeData{Error: true, Message: "Failed to delete candidate post, something went wrong", Data: nil}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	if err := voteRepositorys.DeleteCandidateVote(candidate.ID); err != nil {
		res := schemas.ResponeData{Error: true, Message: "Failed to delete candidate vote, something went wrong", Data: nil}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	if err := candidateRepositorys.DeleteCandidate(candidate.ID); err != nil {
		res := schemas.ResponeData{Error: true, Message: "Failed to delete candidate, something went wrong", Data: nil}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := schemas.ResponeData{Error: false, Message: "Delete candidate successfully", Data: nil}
	c.JSON(http.StatusOK, res)
}
