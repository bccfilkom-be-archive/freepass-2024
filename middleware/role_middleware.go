package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/nathakusuma/bcc-be-freepass-2024/model"
	"github.com/nathakusuma/bcc-be-freepass-2024/repository"
	"github.com/nathakusuma/bcc-be-freepass-2024/util/apiresponse"
	"github.com/nathakusuma/bcc-be-freepass-2024/util/errortypes"
	"net/http"
)

type RoleMiddleware struct {
	UserRepository *repository.UserRepository
}

func NewRoleMiddleware(repo *repository.UserRepository) *RoleMiddleware {
	return &RoleMiddleware{repo}
}

func (mid *RoleMiddleware) RequireRole(roles ...string) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		claimsTemp, _ := ctx.Get("user")
		claims := claimsTemp.(model.UserClaims)
		issuerId := claims.ID
		issuer, _ := mid.UserRepository.FindById(issuerId)
		if issuer == nil {
			ctx.Abort()
			apiresponse.ApiError(ctx, &errortypes.ApiError{
				Code:    http.StatusUnauthorized,
				Message: "invalid token/user deleted",
				Data:    gin.H{},
			})
			return
		}

		ctx.Set("role", issuer.Role)

		isValid := false
		for _, role := range roles {
			isValid = isValid || (issuer.Role == role)
		}

		if !isValid {
			ctx.Abort()
			apiresponse.ApiError(ctx, &errortypes.ApiError{
				Code:    http.StatusForbidden,
				Message: "no permission",
				Data:    gin.H{},
			})
			return
		}
		ctx.Next()
	}
}
