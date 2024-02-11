package electionController

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/AkbarFikri/freepassBCC-2024/models"
	electionRepositorys "github.com/AkbarFikri/freepassBCC-2024/repositorys/election"
	"github.com/AkbarFikri/freepassBCC-2024/schemas"
)

func EditElection(c *gin.Context) {
	electionID := c.Param("id")
	getUser, _ := c.Get("user")
	user := getUser.(schemas.UserTokenData)
	var electionData *models.Election

	if err := c.ShouldBindJSON(&electionData); err != nil {
		res := schemas.ResponeData{Error: true, Message: "No data provided", Data: nil}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if electionData.Title == "" || electionData.Description == "" {
		res := schemas.ResponeData{Error: true, Message: "Title, description data must be provided", Data: nil}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	startTime, err := time.Parse("2006-01-02 15:04:05", electionData.StartTime)
	if err != nil {
		res := schemas.ResponeData{Error: true, Message: "Wrong start_time format", Data: nil}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	endTime, err := time.Parse("2006-01-02 15:04:05", electionData.EndTime)
	if err != nil {
		res := schemas.ResponeData{Error: true, Message: "Wrong end_time format", Data: nil}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	if startTime.IsZero() || endTime.IsZero() {
		res := schemas.ResponeData{Error: true, Message: "Start_time, end_time must be provided", Data: nil}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	parseTimeNow := time.Now().Local().Format("2006-01-02 15:04:05")
	timeNow, err := time.Parse("2006-01-02 15:04:05", parseTimeNow)
	if err != nil {
		res := schemas.ResponeData{Error: true, Message: "Something went wrong", Data: nil}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	if endTime.Before(timeNow) && startTime.Before(timeNow) {
		res := schemas.ResponeData{Error: true, Message: "Election period is not valid", Data: nil}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if !user.IsAdmin {
		res := schemas.ResponeData{Error: true, Message: "You're not allowed to create election", Data: nil}
		c.JSON(http.StatusNotFound, res)
		return
	}

	if startTime.After(timeNow) && endTime.Before(timeNow) {
		electionData.Status = "Ongoing"
	} else {
		electionData.Status = "Pending"
	}

	election, err := electionRepositorys.FindSpecificElection(electionID)
	if err != nil {
		res := schemas.ResponeData{Error: true, Message: "Election with id provided is not found", Data: nil}
		c.JSON(http.StatusNotFound, res)
		return
	}

	election.Title = electionData.Title
	election.Description = electionData.Description
	election.StartTime = electionData.StartTime
	election.EndTime = electionData.EndTime

	if err := electionRepositorys.EditElection(election); err != nil {
		res := schemas.ResponeData{Error: true, Message: "Failed to edit election, Something went wrong", Data: nil}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := schemas.ResponeData{Error: false, Message: "Create new election successfully", Data: election}
	c.JSON(http.StatusOK, res)
}
