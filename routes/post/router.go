package postRoute

import (
	"github.com/gin-gonic/gin"

	postController "github.com/AkbarFikri/freepassBCC-2024/controllers/post"
)

func RegisterRoute(route *gin.RouterGroup) {
	route.GET("/", postController.GetAllPost)
	route.GET("/:id", postController.GetSpecificPost)
	route.PATCH("/:id", postController.UpdatePost)
	route.DELETE("/:id", postController.DeletePost)
	route.GET("/election/:election_id", postController.GetAllPostByElection)
	route.POST("/election/:election_id", postController.CreateNewCandidatePost)
	route.GET("/candidate/:candidate_id", postController.GetAllCandidatePosts)
	route.GET("/election/:election_id/candidate/:candidate_id", postController.GetAllCandidatepostsByElection)
}
