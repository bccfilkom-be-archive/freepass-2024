package electionController

import (
	"net/http"

	"github.com/gin-gonic/gin"

	electionRepositorys "github.com/AkbarFikri/freepassBCC-2024/repositorys/election"
	"github.com/AkbarFikri/freepassBCC-2024/schemas"
)

func GetAllElection(c *gin.Context) {
	elections, err := electionRepositorys.FindAllElection()
	if err != nil {
		res := schemas.ResponeData{Error: true, Message: "Failed to find election, Something Went Wrong", Data: nil}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	if lenOfSlice := len(elections); lenOfSlice == 0 {
		res := schemas.ResponeData{Error: true, Message: "No Data Found", Data: nil}
		c.JSON(http.StatusNotFound, res)
		return
	}

	res := schemas.ResponeData{Error: false, Message: "Find all elections successfully", Data: elections}
	c.JSON(http.StatusOK, res)
	return
}
