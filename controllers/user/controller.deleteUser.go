package userController

import (
	"net/http"

	"github.com/gin-gonic/gin"

	userRepositorys "github.com/AkbarFikri/freepassBCC-2024/repositorys/user"
	"github.com/AkbarFikri/freepassBCC-2024/schemas"
)

func DeleteUser(c *gin.Context) {
	userID := c.Param("id")
	getUser, _ := c.Get("user")
	user := getUser.(schemas.UserTokenData)

	if userID == "" {
		res := schemas.ResponeData{Error: true, Message: "id as a Path Param is required", Data: nil}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	userData, err := userRepositorys.FindOne(userID)
	if err != nil {
		res := schemas.ResponeData{Error: true, Message: "User with id provided is not found", Data: nil}
		c.JSON(http.StatusNotFound, res)
		return
	}

	if !user.IsAdmin {
		res := schemas.ResponeData{Error: true, Message: "You're not allowed to delete this data", Data: nil}
		c.JSON(http.StatusNotFound, res)
		return
	}

	if err := userRepositorys.DeleteUser(userData.ID); err != nil {
		res := schemas.ResponeData{Error: true, Message: "Failed to delete user, something went wrong", Data: nil}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := schemas.ResponeData{Error: false, Message: "Delete user successfully", Data: nil}
	c.JSON(http.StatusOK, res)
}
