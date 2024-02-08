package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"

	"github.com/AkbarFikri/freepassBCC-2024/schemas"
	"github.com/AkbarFikri/freepassBCC-2024/utils"
)

func Auth() gin.HandlerFunc {

	return gin.HandlerFunc(func(c *gin.Context) {
		var res schemas.ResponeData

		if c.GetHeader("Authorization") == "" {
			res.Error = true
			res.Message = "Authorization is required for this endpoint"
			res.Data = nil
			c.AbortWithStatusJSON(http.StatusForbidden, res)
		}

		token, err := utils.VerifyTokenHeader(c, "JWT_SECRET")

		if err != nil {
			res.Error = true
			res.Message = "accessToken invalid or expired"
			res.Data = nil
			c.AbortWithStatusJSON(http.StatusUnauthorized, res)
		} else {
			claims := token.Claims.(jwt.MapClaims)
			user := schemas.UserTokenData{ID: claims["id"].(string), Email: claims["email"].(string)}
			c.Set("user", user)
			c.Next()
		}
	})

}
