package candidateController

import (
	"net/http"

	"github.com/gin-gonic/gin"

	candidateRepositorys "github.com/AkbarFikri/freepassBCC-2024/repositorys/candidate"
	"github.com/AkbarFikri/freepassBCC-2024/schemas"
)

func GetOneCandidateInformation(c *gin.Context) {
	candidateID := c.Param("candidate_id")

	if candidateID == "" {
		res := schemas.ResponeData{Error: true, Message: "candidate_id as a Param is required", Data: nil}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	candidate, err := candidateRepositorys.FindSpecificCandidate(candidateID)
	if candidate.ID == "" {
		res := schemas.ResponeData{Error: true, Message: "Candidate is not found", Data: nil}
		c.JSON(http.StatusNotFound, res)
		return
	}

	if err != nil {
		res := schemas.ResponeData{Error: true, Message: "Something went wrong", Data: nil}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	informations, err := candidateRepositorys.FindCandidateInformations(candidate.ID)
	if informations.ID == "" {
		res := schemas.ResponeData{Error: true, Message: "Candidate Informations is not found", Data: nil}
		c.JSON(http.StatusNotFound, res)
		return
	}

	if err != nil {
		res := schemas.ResponeData{Error: true, Message: "Something went wrong", Data: nil}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := schemas.ResponeData{Error: false, Message: "Successfully find candidate informations", Data: informations}
	c.JSON(http.StatusOK, res)
	return

}
