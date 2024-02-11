package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nathakusuma/bcc-be-freepass-2024/model"
	"github.com/nathakusuma/bcc-be-freepass-2024/service"
	"github.com/nathakusuma/bcc-be-freepass-2024/util/apiresponse"
	"github.com/nathakusuma/bcc-be-freepass-2024/util/binding"
	"github.com/nathakusuma/bcc-be-freepass-2024/util/roles"
	"net/http"
)

type CommentHandler struct {
	CommentService *service.CommentService
}

func NewCommentHandler(commService *service.CommentService) *CommentHandler {
	return &CommentHandler{commService}
}

func (handler *CommentHandler) Add(ctx *gin.Context) {
	postId, err := binding.ShouldUintQuery(ctx, "postId")
	if err != nil {
		apiresponse.ApiError(ctx, err)
		return
	}

	claimsTemp, _ := ctx.Get("user")
	claims := claimsTemp.(model.UserClaims)
	issuerId := claims.ID

	var request model.AddCommentRequest
	if err := binding.ShouldBindJSON(ctx, &request); err != nil {
		apiresponse.ApiError(ctx, err)
		return
	}

	response, apiErr := handler.CommentService.Add(&request, postId, issuerId)
	if apiErr != nil {
		apiresponse.ApiError(ctx, apiErr)
		return
	}

	apiresponse.Success(ctx, http.StatusOK, "comment added", response)
}

func (handler *CommentHandler) Delete(ctx *gin.Context) {
	commentId, err := binding.ShouldUintQuery(ctx, "commentId")
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

	if err := handler.CommentService.DeleteById(commentId, issuerId, isAdmin); err != nil {
		apiresponse.ApiError(ctx, err)
		return
	}

	apiresponse.Success(ctx, http.StatusOK, "successfully deleted comment", gin.H{})
}
