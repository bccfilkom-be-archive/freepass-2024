package RegisteredRoute

import (
	"github.com/gin-gonic/gin"

	electionController "github.com/AkbarFikri/freepassBCC-2024/controllers/election"

)

func ElectionRoute(route *gin.RouterGroup) {
	route.GET("/", electionController.GetAllElection)
	route.GET("/candidate/:candidate_id", electionController.GetOneCandidateInformation)
	route.POST("/:election_id/candidate/:candidate_id/vote", electionController.VoteCandidate)
	route.GET("/:election_id/candidate", electionController.GetAllCandidateListByElection)
}
