package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	repositorys "github.com/AkbarFikri/freepassBCC-2024/repositorys/user"
	"github.com/AkbarFikri/freepassBCC-2024/schemas"
)

func LoginController(c *gin.Context) {
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

	user, err := repositorys.FindUser(request)
	if err != nil {
		res := schemas.ResponeData{Error: true, Message: "Invalid Email or Password", Data: nil}
		c.JSON(http.StatusBadRequest, res)
		return
	}
}
