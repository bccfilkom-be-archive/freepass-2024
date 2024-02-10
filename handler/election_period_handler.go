package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nathakusuma/bcc-be-freepass-2024/model"
	"github.com/nathakusuma/bcc-be-freepass-2024/service"
	"github.com/nathakusuma/bcc-be-freepass-2024/util/apiresponse"
	"github.com/nathakusuma/bcc-be-freepass-2024/util/binding"
	"github.com/nathakusuma/bcc-be-freepass-2024/util/errortypes"
	"net/http"
	"time"
)

type ElectionPeriodHandler struct {
	PeriodService *service.ElectionPeriodService
}

func NewElectionPeriodHandler(serv *service.ElectionPeriodService) *ElectionPeriodHandler {
	return &ElectionPeriodHandler{serv}
}

func (handler *ElectionPeriodHandler) Set(ctx *gin.Context) {
	var request model.SetElectionPeriodRequest
	if err := binding.ShouldBindJSON(ctx, &request); err != nil {
		apiresponse.ApiError(ctx, err)
		return
	}

	start, err1 := time.Parse(time.RFC3339, request.Start)
	end, err2 := time.Parse(time.RFC3339, request.End)
	if err1 != nil || err2 != nil {
		apiresponse.ApiError(ctx, &errortypes.ApiError{
			Code:    http.StatusBadRequest,
			Message: "invalid date time format",
			Data: gin.H{
				"err_start": err1,
				"err_end":   err2,
			},
		})
		return
	}

	err := handler.PeriodService.SetPeriod(start, end)
	if err != nil {
		apiresponse.ApiError(ctx, err)
		return
	}

	apiresponse.Success(ctx, http.StatusCreated, "successfully saved election period", gin.H{})
}

func (handler *ElectionPeriodHandler) Get(ctx *gin.Context) {
	response, err := handler.PeriodService.GetPeriod()
	if err != nil {
		apiresponse.ApiError(ctx, err)
		return
	}

	apiresponse.Success(ctx, http.StatusOK, "successfully retrieved election period data", response)
}
