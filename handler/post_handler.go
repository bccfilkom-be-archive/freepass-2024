package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nathakusuma/bcc-be-freepass-2024/model"
	"github.com/nathakusuma/bcc-be-freepass-2024/service"
	"github.com/nathakusuma/bcc-be-freepass-2024/util/apiresponse"
	"github.com/nathakusuma/bcc-be-freepass-2024/util/binding"
	"github.com/nathakusuma/bcc-be-freepass-2024/util/errortypes"
	"github.com/nathakusuma/bcc-be-freepass-2024/util/roles"
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

func (handler *PostHandler) Create(ctx *gin.Context) {
	var request model.CreatePostRequest
	if err := binding.ShouldBindJSON(ctx, &request); err != nil {
		apiresponse.ApiError(ctx, err)
		return
	}

	claimsTemp, _ := ctx.Get("user")
	claims := claimsTemp.(model.UserClaims)
	issuerId := claims.ID

	response, err := handler.PostService.Create(&request, issuerId)
	if err != nil {
		apiresponse.ApiError(ctx, err)
		return
	}

	apiresponse.Success(ctx, http.StatusCreated, "post created", response)
}

func (handler *PostHandler) Update(ctx *gin.Context) {
	postId, err := binding.ShouldUintQuery(ctx, "postId")
	if err != nil {
		apiresponse.ApiError(ctx, err)
		return
	}

	claimsTemp, _ := ctx.Get("user")
	claims := claimsTemp.(model.UserClaims)
	issuerId := claims.ID

	var request model.UpdatePostRequest
	if err := binding.ShouldBindJSON(ctx, &request); err != nil {
		apiresponse.ApiError(ctx, err)
		return
	}

	if err := handler.PostService.Update(postId, issuerId, &request); err != nil {
		apiresponse.ApiError(ctx, err)
		return
	}

	apiresponse.Success(ctx, http.StatusCreated, "successfully edited post", gin.H{})
}

func (handler *PostHandler) Delete(ctx *gin.Context) {
	postId, err := binding.ShouldUintQuery(ctx, "postId")
	if err != nil {
		apiresponse.ApiError(ctx, err)
		return
	}

	role, _ := ctx.Get("role")
	role = role.(string)
	isAdmin := role == roles.Admin

	var issuerId uint
	if !isAdmin {
		claimsTemp, _ := ctx.Get("user")
		claims := claimsTemp.(model.UserClaims)
		issuerId = claims.ID
	}

	if err := handler.PostService.DeleteById(postId, issuerId, isAdmin); err != nil {
		apiresponse.ApiError(ctx, err)
		return
	}

	apiresponse.Success(ctx, http.StatusOK, "successfully deleted post", gin.H{})
}
