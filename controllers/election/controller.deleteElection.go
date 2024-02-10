package electionController

import (
	"net/http"

	"github.com/gin-gonic/gin"

	electionRepositorys "github.com/AkbarFikri/freepassBCC-2024/repositorys/election"
	"github.com/AkbarFikri/freepassBCC-2024/schemas"

)

func DeleteElection(c *gin.Context) {
	electionID := c.Param("id")
	getUser, _ := c.Get("user")
	user := getUser.(schemas.UserTokenData)

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

	if !user.IsAdmin {
		res := schemas.ResponeData{Error: true, Message: "You're not allowed to delete this data", Data: nil}
		c.JSON(http.StatusNotFound, res)
		return
	}

	if err := electionRepositorys.DeleteElection(election.ID); err != nil {
		res := schemas.ResponeData{Error: true, Message: "Failed to delete election, something went wrong", Data: nil}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := schemas.ResponeData{Error: false, Message: "Delete election successfully", Data: nil}
	c.JSON(http.StatusOK, res)
}
