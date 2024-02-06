package handler

import (
	"freepass-bcc/app/comment/usecase"
	"freepass-bcc/domain"
	"freepass-bcc/help"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	commentUsecase usecase.ICommentUsecase
}

func NewCommentHandler(commentUsecase usecase.ICommentUsecase) *CommentHandler {
	return &CommentHandler{commentUsecase}
}

func (h *CommentHandler) CreateComment(c *gin.Context) {
	postIdString := c.Param("postId")
	postId, _ := strconv.Atoi(postIdString)

	var commentRequest domain.CommentRequest
	err := c.ShouldBindJSON(&commentRequest)
	if err != nil {
		help.FailedResponse(c, http.StatusBadRequest, "failed to bind comment request", err)
		return
	}

	comment, errorObject := h.commentUsecase.CreateComment(c, commentRequest, postId)
	if errorObject != nil {
		errorObject := errorObject.(help.ErrorObject)
		help.FailedResponse(c, errorObject.Code, errorObject.Message, errorObject.Err)
		return
	}

	help.SuccessResponse(c, "success create comment", comment)
}
