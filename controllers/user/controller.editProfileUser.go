package userController

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/AkbarFikri/freepassBCC-2024/models"
	profileRepositorys "github.com/AkbarFikri/freepassBCC-2024/repositorys/profile"
	"github.com/AkbarFikri/freepassBCC-2024/schemas"
)

func EditProfileUser(c *gin.Context) {
	var profile *models.Profile
	getUser, _ := c.Get("user")
	user := getUser.(schemas.UserTokenData)

	if err := c.ShouldBindJSON(&profile); err != nil {
		res := schemas.ResponeData{Error: true, Message: "No Data Found", Data: nil}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	profile.UserID = user.ID

	if err := profileRepositorys.EditProfile(profile); err != nil {
		res := schemas.ResponeData{Error: true, Message: "Something went wrong", Data: nil}
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	res := schemas.ResponeData{Error: false, Message: "Berhasil", Data: profile}
	c.JSON(http.StatusOK, res)
}
