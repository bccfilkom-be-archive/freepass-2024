package adminRoute

import (
	"github.com/gin-gonic/gin"

	authController "github.com/AkbarFikri/freepassBCC-2024/controllers/auth"
	electionController "github.com/AkbarFikri/freepassBCC-2024/controllers/election"
)

func RegisterRoute(route *gin.RouterGroup) {
	route.POST("/login", authController.LoginAdmin)
	route.POST("/register", authController.RegisterAdmin)
	route.POST("/election/:election_id/candidate", electionController.PromoteCandidate)
}
