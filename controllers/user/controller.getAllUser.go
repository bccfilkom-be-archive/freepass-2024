package userController

import (
	"net/http"

	"github.com/gin-gonic/gin"

	userRepositorys "github.com/AkbarFikri/freepassBCC-2024/repositorys/user"
	"github.com/AkbarFikri/freepassBCC-2024/schemas"

)

func GetAllUser(c *gin.Context) {
	users, err := userRepositorys.FindAll()

	if err != nil {
		res := schemas.ResponeData{Error: true, Message: "Failed to find users, Something Went Wrong", Data: nil}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	if lenOfSlice := len(users); lenOfSlice == 0 {
		res := schemas.ResponeData{Error: true, Message: "No Data Found", Data: nil}
		c.JSON(http.StatusNotFound, res)
		return
	}

	res := schemas.ResponeData{Error: false, Message: "Find all posts successfully", Data: users}
	c.JSON(http.StatusOK, res)
	return
}
