package commentController

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/AkbarFikri/freepassBCC-2024/models"
	commentRepositorys "github.com/AkbarFikri/freepassBCC-2024/repositorys/comment"
	postRepositorys "github.com/AkbarFikri/freepassBCC-2024/repositorys/post"
	"github.com/AkbarFikri/freepassBCC-2024/schemas"
)

func CreateComment(c *gin.Context) {
	var comment *models.Comment
	PostID := c.Param("post_id")
	getUser, _ := c.Get("user")
	user := getUser.(schemas.UserTokenData)

	if PostID == "" {
		res := schemas.ResponeData{Error: true, Message: "post_id as a Param is required", Data: nil}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if err := c.ShouldBindJSON(&comment); err != nil {
		res := schemas.ResponeData{Error: true, Message: "No data provided", Data: nil}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if comment.Message == "" {
		res := schemas.ResponeData{Error: true, Message: "Comment message is required", Data: nil}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	post, err := postRepositorys.FindOne(PostID)
	if err != nil {
		res := schemas.ResponeData{Error: true, Message: "Post with id provided is not found", Data: nil}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	comment.PostID = post.ID
	comment.UserID = user.ID

	if err := commentRepositorys.CreateComment(comment); err != nil {
		res := schemas.ResponeData{Error: true, Message: "Failed to create comment, Something went wrong", Data: nil}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := schemas.ResponeData{Error: false, Message: "Create comment successfully", Data: comment}
	c.JSON(http.StatusCreated, res)
	return
}
