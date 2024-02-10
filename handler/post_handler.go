package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nathakusuma/bcc-be-freepass-2024/service"
	"github.com/nathakusuma/bcc-be-freepass-2024/util/apiresponse"
	"github.com/nathakusuma/bcc-be-freepass-2024/util/binding"
	"github.com/nathakusuma/bcc-be-freepass-2024/util/errortypes"
	"net/http"
)

type PostHandler struct {
	PostService *service.PostService
}

func NewPostHandler(postService *service.PostService) *PostHandler {
	return &PostHandler{postService}
}

func (handler *PostHandler) Get(ctx *gin.Context) {
	if postId, err := binding.ShouldUintQuery(ctx, "postId"); err == nil {
		response, apiErr := handler.PostService.GetById(postId)
		if apiErr != nil {
			apiresponse.ApiError(ctx, apiErr)
			return
		}
		apiresponse.Success(ctx, http.StatusOK, "successfully retrieved post data", response)
	} else if candidateId, err := binding.ShouldUintQuery(ctx, "candidateId"); err == nil {
		response, apiErr := handler.PostService.GetByCandidateId(candidateId)
		if apiErr != nil {
			apiresponse.ApiError(ctx, apiErr)
			return
		}
		apiresponse.Success(ctx, http.StatusOK, "successfully retrieved post data", response)
	} else {
		apiresponse.ApiError(ctx, &errortypes.ApiError{
			Code:    http.StatusBadRequest,
			Message: "invalid query",
			Data:    gin.H{},
		})
		return
	}
}
