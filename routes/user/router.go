package userRoute

import (
	"github.com/gin-gonic/gin"

	userController "github.com/AkbarFikri/freepassBCC-2024/controllers/user"

)

func RegisterRoute(route *gin.RouterGroup) {
	route.POST("/profile", userController.EditProfileUser)
}