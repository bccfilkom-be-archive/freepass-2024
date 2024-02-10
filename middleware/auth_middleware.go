package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/nathakusuma/bcc-be-freepass-2024/model"
	"github.com/nathakusuma/bcc-be-freepass-2024/util/apiresponse"
	"github.com/nathakusuma/bcc-be-freepass-2024/util/auth"
	"github.com/nathakusuma/bcc-be-freepass-2024/util/errortypes"
	"net/http"
	"os"
	"strings"
)

func Auth(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("Authorization")
	if !strings.HasPrefix(authorization, "Bearer ") {
		ctx.Abort()
		apiresponse.ApiError(ctx, &errortypes.ApiError{
			Code:    http.StatusUnauthorized,
			Message: "user not authenticated",
			Data:    gin.H{},
		})
		return
	}
	token := authorization[7:]
	claims := model.UserClaims{}
	KEY := os.Getenv("JWT_SECRET_KEY")
	if err := auth.DecodeToken(token, &claims, KEY); err != nil {
		ctx.Abort()
		apiresponse.ApiError(ctx, &errortypes.ApiError{
			Code:    http.StatusUnauthorized,
			Message: "user not authenticated",
			Data:    gin.H{},
		})
		return
	}
	ctx.Set("user", claims)
	ctx.Next()
}
