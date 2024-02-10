package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nathakusuma/bcc-be-freepass-2024/service"
	"github.com/nathakusuma/bcc-be-freepass-2024/util/apiresponse"
	"net/http"
)

type CandidateHandler struct {
	CandidateService *service.CandidateService
}

func NewCandidateHandler(candidateService *service.CandidateService) *CandidateHandler {
	return &CandidateHandler{candidateService}
}

func (handler *CandidateHandler) Get(ctx *gin.Context) {
	candidates, err := handler.CandidateService.GetAll()
	if err != nil {
		apiresponse.ApiError(ctx, err)
		return
	}

	apiresponse.Success(ctx, http.StatusOK, "successfully retrieved candidates data", candidates)
}
