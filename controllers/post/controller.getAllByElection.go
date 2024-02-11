package postController

import (
	"net/http"

	"github.com/gin-gonic/gin"

	postRepositorys "github.com/AkbarFikri/freepassBCC-2024/repositorys/post"
	"github.com/AkbarFikri/freepassBCC-2024/schemas"
)

func GetAllPostByElection(c *gin.Context) {
	electionID := c.Param("election_id")

	if electionID == "" {
		res := schemas.ResponeData{Error: true, Message: "election_id as a Param is required", Data: nil}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	posts, err := postRepositorys.FindAllByElectionID(electionID)
	if err != nil {
		res := schemas.ResponeData{Error: true, Message: "Something Went Wrong", Data: nil}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	if lenOfSlice := len(posts); lenOfSlice == 0 {
		res := schemas.ResponeData{Error: true, Message: "No Data Found", Data: nil}
		c.JSON(http.StatusNotFound, res)
		return
	}

	res := schemas.ResponeData{Error: false, Message: "Find all posts by election_id successfully", Data: posts}
	c.JSON(http.StatusOK, res)
	return
}
