package electionRoute

import (
	"github.com/gin-gonic/gin"

	electionController "github.com/AkbarFikri/freepassBCC-2024/controllers/election"
)

func RegisterRoute(route *gin.RouterGroup) {
	route.GET("/candidate/:candidate_id", electionController.GetOneCandidateInformation)
	route.POST("/:election_id/candidate/:candidate_id/vote", electionController.VoteCandidate)
	route.POST("/:election_id/candidate", electionController.PromoteCandidate)
}
