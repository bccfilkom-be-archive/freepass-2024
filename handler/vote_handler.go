package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nathakusuma/bcc-be-freepass-2024/model"
	"github.com/nathakusuma/bcc-be-freepass-2024/service"
	"github.com/nathakusuma/bcc-be-freepass-2024/util/apiresponse"
	"github.com/nathakusuma/bcc-be-freepass-2024/util/binding"
	"net/http"
)

type VoteHandler struct {
	VoteService *service.VoteService
}

func NewVoteHandler(voteService *service.VoteService) *VoteHandler {
	return &VoteHandler{voteService}
}

func (handler *VoteHandler) AddVote(ctx *gin.Context) {
	candidateId, err := binding.ShouldUintQuery(ctx, "candidateId")
	if err != nil {
		apiresponse.ApiError(ctx, err)
		return
	}

	claimsTemp, _ := ctx.Get("user")
	voterId := claimsTemp.(model.UserClaims).ID

	servErr := handler.VoteService.AddVote(candidateId, voterId)
	if servErr != nil {
		apiresponse.ApiError(ctx, servErr)
		return
	}

	apiresponse.Success(ctx, http.StatusCreated, "vote recorded", gin.H{})
}
