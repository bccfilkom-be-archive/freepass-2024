package authController

import (
	"net/http"

	"github.com/gin-gonic/gin"

	userRepositorys "github.com/AkbarFikri/freepassBCC-2024/repositorys/user"
	"github.com/AkbarFikri/freepassBCC-2024/schemas"
	"github.com/AkbarFikri/freepassBCC-2024/utils"
)

func LoginUser(c *gin.Context) {
	var request *schemas.UserLoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		res := schemas.ResponeData{Error: true, Message: "Email, Name, Password is Required", Data: nil}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if request.Email == "" || request.Password == "" {
		res := schemas.ResponeData{Error: true, Message: "Email, Name, Password is Required", Data: nil}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	user, err := userRepositorys.FindUser(request)
	if err != nil {
		res := schemas.ResponeData{Error: true, Message: "Invalid Email or Password", Data: nil}
		c.JSON(http.StatusNotFound, res)
		return
	}

	if err := utils.ComparePassword(user.Password, request.Password); err != nil {
		res := schemas.ResponeData{Error: true, Message: "Invalid Email or Password", Data: nil}
		c.JSON(http.StatusNotFound, res)
		return
	}

	accessData := map[string]interface{}{"id": user.ID, "email": user.Email}
	accessToken, err := utils.SignJWT(accessData, "JWT_SECRET", 24)

	if err != nil {
		res := schemas.ResponeData{Error: true, Message: "Something went wrong", Data: nil}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := schemas.ResponeData{Error: false, Message: "Login successfully", Data: gin.H{"token": accessToken}}
	c.JSON(http.StatusOK, res)
}
