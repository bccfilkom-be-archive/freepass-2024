package handler

import (
	"freepass-bcc/app/post/usecase"
	"freepass-bcc/domain"
	"freepass-bcc/help"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	postUsecase usecase.IPostUsecase
}

func NewPostHandler(usecase usecase.IPostUsecase) *PostHandler {
	return &PostHandler{usecase}
}

func (h *PostHandler) GetAllPost(c *gin.Context) {
	posts, errorObject := h.postUsecase.GetAllPost()
	if errorObject != nil {
		errorObject := errorObject.(help.ErrorObject)
		help.FailedResponse(c, errorObject.Code, errorObject.Message, errorObject.Err)
		return
	}

	help.SuccessResponse(c, "succes get all post", posts)
}

func (h *PostHandler) GetPost(c *gin.Context) {
	postIdString := c.Param("postId")
	postId, _ := strconv.Atoi(postIdString)

	post, errorObject := h.postUsecase.GetPost(postId)
	if errorObject != nil {
		errorObject := errorObject.(help.ErrorObject)
		help.FailedResponse(c, errorObject.Code, errorObject.Message, errorObject.Err)
		return
	}

	help.SuccessResponse(c, "success get post", post)
}

func (h *PostHandler) CreatePost(c *gin.Context) {
	var postRequest domain.PostRequest
	err := c.ShouldBindJSON(&postRequest)
	if err != nil {
		help.FailedResponse(c, http.StatusBadRequest, "failed to bind post request", err)
		return
	}

	post, errorObject := h.postUsecase.CreatePost(c, postRequest)
	if errorObject != nil {
		errorObject := errorObject.(help.ErrorObject)
		help.FailedResponse(c, errorObject.Code, errorObject.Message, errorObject.Err)
		return
	}

	help.SuccessResponse(c, "success create a post", post)
}

func (h *PostHandler) UpdatePost(c *gin.Context) {
	postIdString := c.Param("postId")
	postId, _ := strconv.Atoi(postIdString)

	var postRequest domain.PostRequest
	err := c.ShouldBindJSON(&postRequest)
	if err != nil {
		help.FailedResponse(c, http.StatusBadRequest, "failed to bind post request", err)
		return
	}

	updatedPost, errorObject := h.postUsecase.UpdatePost(c, postRequest, postId)
	if errorObject != nil {
		errorObject := errorObject.(help.ErrorObject)
		help.FailedResponse(c, errorObject.Code, errorObject.Message, errorObject.Err)
		return
	}

	help.SuccessResponse(c, "success update post", updatedPost)
}

func (h *PostHandler) DeletePost(c *gin.Context) {
	postIdString := c.Param("postId")
	postId, _ := strconv.Atoi(postIdString)

	deletedPost, errorObject := h.postUsecase.DeletePost(c, postId)
	if errorObject != nil {
		errorObject := errorObject.(help.ErrorObject)
		help.FailedResponse(c, errorObject.Code, errorObject.Message, errorObject.Err)
		return
	}

	help.SuccessResponse(c, "success delete post", deletedPost)
}
