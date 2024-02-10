package RegisteredRoute

import (
	"github.com/gin-gonic/gin"

	electionController "github.com/AkbarFikri/freepassBCC-2024/controllers/election"
	userController "github.com/AkbarFikri/freepassBCC-2024/controllers/user"
)

func AdminRoute(route *gin.RouterGroup) {
	route.POST("/election/:election_id/candidate", electionController.PromoteCandidate)
	route.POST("/election", electionController.CreateElection)
	route.PATCH("/election/:id", electionController.EditElection)
	route.DELETE("/election/:id", electionController.DeleteElection)
	route.DELETE("/candidate/:id", electionController.DeleteCandidate)
	route.DELETE("/user/:id", userController.DeleteUser)
}
