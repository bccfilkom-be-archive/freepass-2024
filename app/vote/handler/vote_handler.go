package handler

import (
	"freepass-bcc/app/vote/usecase"
	"freepass-bcc/help"
	"strconv"

	"github.com/gin-gonic/gin"
)

type VoteHandler struct {
	voteUsecase usecase.IVoteUsecase
}

func NewVoteHandler(voteUsecase usecase.IVoteUsecase) *VoteHandler {
	return &VoteHandler{voteUsecase}
}

func (h *VoteHandler) Vote(c *gin.Context) {
	userIdString := c.Param("userId")
	userId, _ := strconv.Atoi(userIdString)

	voted, errorObject := h.voteUsecase.Vote(c, userId)
	if errorObject != nil {
		errorObject := errorObject.(help.ErrorObject)
		help.FailedResponse(c, errorObject.Code, errorObject.Message, errorObject.Err)
		return
	}

	help.SuccessResponse(c, "success to vote", voted)
}
