package candidateRoute

import (
	"github.com/gin-gonic/gin"

	candidateController "github.com/AkbarFikri/freepassBCC-2024/controllers/candidate"

)

func RegisterRoute(route *gin.RouterGroup) {
	route.POST("/:candidate_id", candidateController.GetOneCandidateInformation)
}
