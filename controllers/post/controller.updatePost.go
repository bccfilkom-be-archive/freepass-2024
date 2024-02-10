package postController

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/AkbarFikri/freepassBCC-2024/models"
	postRepositorys "github.com/AkbarFikri/freepassBCC-2024/repositorys/post"
	"github.com/AkbarFikri/freepassBCC-2024/schemas"
)

func UpdatePost(c *gin.Context) {
	postID := c.Param("id")
	var data *models.Post
	getUser, _ := c.Get("user")
	user := getUser.(schemas.UserTokenData)

	if err := c.ShouldBindJSON(&data); err != nil {
		res := schemas.ResponeData{Error: true, Message: "No Data Found", Data: nil}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if data.Caption == "" || data.PictureUrl == "" {
		res := schemas.ResponeData{Error: true, Message: "picture_url, caption is Required", Data: nil}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	post, err := postRepositorys.FindOne(postID)
	if err != nil {
		res := schemas.ResponeData{Error: true, Message: "Post with id provided is not found", Data: nil}
		c.JSON(http.StatusNotFound, res)
		return
	}

	if user.ID != post.UserID {
		res := schemas.ResponeData{Error: true, Message: "You're not allowed to edit this post", Data: nil}
		c.JSON(http.StatusNotAcceptable, res)
		return
	}

	post.Caption = data.Caption
	post.PictureUrl = data.PictureUrl
	post.Comments = nil

	if err := postRepositorys.EditPost(post); err != nil {
		res := schemas.ResponeData{Error: true, Message: "Something went wrong", Data: nil}
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	res := schemas.ResponeData{Error: false, Message: "Edit post data successfully", Data: post}
	c.JSON(http.StatusOK, res)
}
