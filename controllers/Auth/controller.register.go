package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/AkbarFikri/freepassBCC-2024/models"
	repositorys "github.com/AkbarFikri/freepassBCC-2024/repositorys/auth"
	"github.com/AkbarFikri/freepassBCC-2024/schemas"
	"github.com/AkbarFikri/freepassBCC-2024/utils"
)

func RegisterController(c *gin.Context) {
	var user *models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		res := schemas.ResponeData{Error: true, Message: "Email, Name, Password is Required", Data: nil}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if user.Email == "" || user.Password == "" || user.Name == "" {
		res := schemas.ResponeData{Error: true, Message: "Email, Name, Password is Required", Data: nil}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	hashPass, err := utils.HashPassword(user.Password)
	if err != nil {
		res := schemas.ResponeData{Error: true, Message: "Something went wrong", Data: nil}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	user.Password = string(hashPass)
	user.IsAdmin = false

	newUser, err := repositorys.RegisterUser(user)
	if err != nil {
		res := schemas.ResponeData{Error: true, Message: "Email alredy exist", Data: nil}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := schemas.ResponeData{Error: false, Message: "Success to create user", Data: gin.H{"user": newUser}}
	c.JSON(http.StatusOK, res)
}
