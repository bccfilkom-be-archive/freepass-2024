package postController

import (
	"net/http"

	"github.com/gin-gonic/gin"

	postRepositorys "github.com/AkbarFikri/freepassBCC-2024/repositorys/post"
	"github.com/AkbarFikri/freepassBCC-2024/schemas"

)

func GetAllPost(c *gin.Context) {
	posts, err := postRepositorys.FindAll()

	if err != nil {
		res := schemas.ResponeData{Error: true, Message: "Something Went Wrong", Data: nil}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	if lenOfSlice := len(posts); lenOfSlice == 0 {
		res := schemas.ResponeData{Error: true, Message: "No Data Found", Data: nil}
		c.JSON(http.StatusNotFound, res)
		return
	}

	res := schemas.ResponeData{Error: false, Message: "Find all posts successfully", Data: posts}
	c.JSON(http.StatusOK, res)
	return
}
