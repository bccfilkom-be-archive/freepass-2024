package RegisteredRoute

import (
	"github.com/gin-gonic/gin"

	commentController "github.com/AkbarFikri/freepassBCC-2024/controllers/comment"

)

func CommentRoute(route *gin.RouterGroup) {
	route.POST("/post/:post_id", commentController.CreateComment)
	route.DELETE("/:id", commentController.DeleteComment)
}
