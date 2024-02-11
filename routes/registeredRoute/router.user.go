package RegisteredRoute

import (
	"github.com/gin-gonic/gin"

	userController "github.com/AkbarFikri/freepassBCC-2024/controllers/user"
)

func UserRoute(route *gin.RouterGroup) {
	route.GET("/", userController.GetAllUser)
	route.PATCH("/profile", userController.EditProfileUser)
}
