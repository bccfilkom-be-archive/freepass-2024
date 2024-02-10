package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nathakusuma/bcc-be-freepass-2024/service"
	"github.com/nathakusuma/bcc-be-freepass-2024/util/apiresponse"
	"net/http"
)

type ElectionPeriodHandler struct {
	PeriodService *service.ElectionPeriodService
}

func NewElectionPeriodHandler(serv *service.ElectionPeriodService) *ElectionPeriodHandler {
	return &ElectionPeriodHandler{serv}
}

func (handler *ElectionPeriodHandler) Get(ctx *gin.Context) {
	response, err := handler.PeriodService.GetPeriod()
	if err != nil {
		apiresponse.ApiError(ctx, err)
		return
	}

	apiresponse.Success(ctx, http.StatusOK, "successfully retrieved election period data", response)
}
