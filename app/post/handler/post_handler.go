package handler

import (
	"freepass-bcc/app/post/usecase"
	"freepass-bcc/domain"
	"freepass-bcc/help"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	postUsecase usecase.IPostUsecase
}

func NewPostHandler(usecase usecase.IPostUsecase) *PostHandler {
	return &PostHandler{usecase}
}

func (h *PostHandler) CreatePost (c *gin.Context) {
	var postRequest domain.PostRequest
	err := c.ShouldBindJSON(&postRequest)
	if err != nil {
		return
	}

	post, errorObject := h.postUsecase.CreatePost(c, postRequest)
	if err != nil {
		errorObject := errorObject.(help.ErrorObject)
		help.FailedResponse(c, errorObject.Code, errorObject.Message, errorObject.Err)
		return 
	}

	help.SuccessResponse(c, "success create a post", post)
}
