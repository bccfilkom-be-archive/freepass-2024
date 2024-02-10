package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nathakusuma/bcc-be-freepass-2024/entity"
	"github.com/nathakusuma/bcc-be-freepass-2024/service"
	"github.com/nathakusuma/bcc-be-freepass-2024/util/apiresponse"
	"github.com/nathakusuma/bcc-be-freepass-2024/util/binding"
	"gorm.io/gorm"
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

func (handler *CandidateHandler) Promote(ctx *gin.Context) {
	userId, err := binding.ShouldUintQuery(ctx, "userId")
	if err != nil {
		apiresponse.ApiError(ctx, err)
		return
	}

	response, servErr := handler.CandidateService.Promote(&entity.User{
		Model: gorm.Model{ID: userId},
	})
	if servErr != nil {
		apiresponse.ApiError(ctx, servErr)
		return
	}

	apiresponse.Success(ctx, http.StatusCreated, "target promoted from user to candidate", response)

}
