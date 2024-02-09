package postController

import (
	"net/http"

	"github.com/gin-gonic/gin"

	candidateRepositorys "github.com/AkbarFikri/freepassBCC-2024/repositorys/candidate"
	postRepositorys "github.com/AkbarFikri/freepassBCC-2024/repositorys/post"
	"github.com/AkbarFikri/freepassBCC-2024/schemas"
)

func DeletePost(c *gin.Context) {
	postID := c.Param("id")
	getUser, _ := c.Get("user")
	user := getUser.(schemas.UserTokenData)

	post, err := postRepositorys.FindOne(postID)
	if err != nil {
		res := schemas.ResponeData{Error: true, Message: "Post with id provided is not found", Data: nil}
		c.JSON(http.StatusNotFound, res)
		return
	}

	candidate, _ := candidateRepositorys.FindCandidateInElection(user.ID, post.ElectionID)
	if candidate.ID == "" && !user.IsAdmin {
		res := schemas.ResponeData{Error: true, Message: "Your not allowed to delete this post", Data: nil}
		c.JSON(http.StatusNotAcceptable, res)
		return
	}

	if candidate.ID != user.ID {
		res := schemas.ResponeData{Error: true, Message: "Your not allowed to delete this post", Data: nil}
		c.JSON(http.StatusNotAcceptable, res)
		return
	}

	if err := postRepositorys.DeletePost(post.ID); err != nil {
		res := schemas.ResponeData{Error: true, Message: "Failed to delete post, Something went wrong", Data: nil}
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	res := schemas.ResponeData{Error: false, Message: "Delete post successfully", Data: nil}
	c.JSON(http.StatusOK, res)
}
