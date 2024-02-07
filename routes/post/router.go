package postRoute

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/AkbarFikri/freepassBCC-2024/schemas"
)

func RegisterRoute(route *gin.RouterGroup) {
	route.GET("/", func(c *gin.Context) {
		data, _ := c.Get("user")
		res := schemas.ResponeData{Error: false, Message: "Menyalaa abangkuhhh", Data: data}
		c.JSON(http.StatusOK, res)
	})
}
