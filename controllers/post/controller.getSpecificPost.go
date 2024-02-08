package postController

import (
	"net/http"

	"github.com/gin-gonic/gin"

	postRepositorys "github.com/AkbarFikri/freepassBCC-2024/repositorys/post"
	"github.com/AkbarFikri/freepassBCC-2024/schemas"

)

func GetSpecificPost(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		res := schemas.ResponeData{Error: true, Message: "id as a Param is required", Data: nil}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	post, err := postRepositorys.FindOne(id)

	if post.ID == "" {
		res := schemas.ResponeData{Error: true, Message: "No Data Found", Data: nil}
		c.JSON(http.StatusNotFound, res)
		return
	}

	if err != nil {
		res := schemas.ResponeData{Error: true, Message: "Something Went Wrong", Data: nil}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := schemas.ResponeData{Error: false, Message: "Find one post successfully", Data: post}
	c.JSON(http.StatusOK, res)
	return
}
