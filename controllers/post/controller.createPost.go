package postController

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/AkbarFikri/freepassBCC-2024/models"
	candidateRepositorys "github.com/AkbarFikri/freepassBCC-2024/repositorys/candidate"
	electionRepositorys "github.com/AkbarFikri/freepassBCC-2024/repositorys/election"
	postRepositorys "github.com/AkbarFikri/freepassBCC-2024/repositorys/post"
	"github.com/AkbarFikri/freepassBCC-2024/schemas"

)

func CreateNewCandidatePost(c *gin.Context) {
	getUser, _ := c.Get("user")
	user := getUser.(schemas.UserTokenData)
	electionID := c.Param("election_id")
	var data *models.Post

	if electionID == "" {
		res := schemas.ResponeData{Error: true, Message: "election_id as a Param is required", Data: nil}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if err := c.ShouldBindJSON(&data); err != nil {
		res := schemas.ResponeData{Error: true, Message: "Post data is Required", Data: nil}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if data.PictureUrl == "" || data.Caption == "" {
		res := schemas.ResponeData{Error: true, Message: "picture_url, caption is Required", Data: nil}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	election, err := electionRepositorys.FindSpecificElection(electionID)
	if err != nil {
		res := schemas.ResponeData{Error: true, Message: "Election with id provided is not found", Data: nil}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	candidate, err := candidateRepositorys.FindCandidateInElection(user.ID, election.ID)
	if candidate.ID == "" {
		res := schemas.ResponeData{Error: true, Message: "Your not allowed to create a post in this election", Data: nil}
		c.JSON(http.StatusAccepted, res)
		return
	}

	data.CandidateID = candidate.ID
	data.UserID = candidate.UserID
	data.ElectionID = election.ID

	post, err := postRepositorys.CreateNewPost(data)
	if err != nil {
		res := schemas.ResponeData{Error: true, Message: "Failed to create post, Something went wrong", Data: nil}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := schemas.ResponeData{Error: false, Message: "Successfully create post", Data: post}
	c.JSON(http.StatusOK, res)
	return
}
