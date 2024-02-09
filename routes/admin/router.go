package adminRoute

import (
	"github.com/gin-gonic/gin"

	electionController "github.com/AkbarFikri/freepassBCC-2024/controllers/election"
)

func RegisterRoute(route *gin.RouterGroup) {
	route.POST("/election/:election_id/candidate", electionController.PromoteCandidate)
}
