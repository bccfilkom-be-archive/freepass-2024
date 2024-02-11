package commentController

import (
	"net/http"

	"github.com/gin-gonic/gin"

	commentRepositorys "github.com/AkbarFikri/freepassBCC-2024/repositorys/comment"
	"github.com/AkbarFikri/freepassBCC-2024/schemas"

)

func DeleteComment(c *gin.Context) {
	commentID := c.Param("id")
	getUser, _ := c.Get("user")
	user := getUser.(schemas.UserTokenData)

	if commentID == "" {
		res := schemas.ResponeData{Error: true, Message: "comment_id as a Param is required", Data: nil}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	comment, err := commentRepositorys.FindOne(commentID)
	if err != nil {
		res := schemas.ResponeData{Error: true, Message: "Comment with id provided is not found", Data: nil}
		c.JSON(http.StatusNotFound, res)
		return
	}

	if comment.UserID != user.ID && !user.IsAdmin {
		res := schemas.ResponeData{Error: true, Message: "Your not allowed to delete this comment", Data: nil}
		c.JSON(http.StatusNotAcceptable, res)
		return
	}

	if err := commentRepositorys.DeleteComment(comment.ID); err != nil {
		res := schemas.ResponeData{Error: true, Message: "Failed to delete post, Something went wrong", Data: nil}
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	res := schemas.ResponeData{Error: false, Message: "Delete post successfully", Data: nil}
	c.JSON(http.StatusOK, res)
}
