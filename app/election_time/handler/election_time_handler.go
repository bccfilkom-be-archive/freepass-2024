package handler

import (
	"freepass-bcc/app/election_time/usecase"
	"freepass-bcc/domain"
	"freepass-bcc/help"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ElectionTimeHandler struct {
	electionTimeUsecase usecase.IElectionTimeUsecase
}

func NewElectionTimeHandler(electionTimeUsecase usecase.IElectionTimeUsecase) *ElectionTimeHandler {
	return &ElectionTimeHandler{electionTimeUsecase}
}

func (h *ElectionTimeHandler) SetStartAndEndTime(c *gin.Context) {
	var electionTime domain.ElectionTimesRequest
	err := c.ShouldBindJSON(&electionTime)
	if err != nil {
		help.FailedResponse(c, http.StatusBadRequest, "failed to bind election time", err)
	}

	setElectionTime, errorObject := h.electionTimeUsecase.SetStartAndEndTime(electionTime)
	if errorObject != nil {
		errorObject := errorObject.(help.ErrorObject)
		help.FailedResponse(c, errorObject.Code, errorObject.Message, errorObject.Err)
		return
	}

	help.SuccessResponse(c, "success set election time", setElectionTime)
}
