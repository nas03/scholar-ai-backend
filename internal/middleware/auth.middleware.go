package middleware

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nas03/scholar-ai/backend/internal/helper"
	"github.com/nas03/scholar-ai/backend/pkg/response"
)

type IAuthMiddleware interface {
	Auth() gin.HandlerFunc
}

type AuthMiddleware struct {
	jwtHelper helper.IJWTHelper
}

func NewAuthMiddleware(jwtHelper helper.IJWTHelper) IAuthMiddleware {
	return &AuthMiddleware{jwtHelper: jwtHelper}
}

func (m *AuthMiddleware) Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			response.ErrorResponse(ctx, response.CodeTokenInvalid, "missing Authorization header")
			ctx.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := m.jwtHelper.ValidateAuthToken(ctx, token)
		if err != nil {
			// Decide whether token is expired or otherwise invalid
			if errors.Is(err, jwt.ErrTokenExpired) {
				response.ErrorResponse(ctx, response.CodeTokenExpired, "")
			} else {
				response.ErrorResponse(ctx, response.CodeTokenInvalid, err.Error())
			}
			ctx.Abort()
			return
		}

		// attach user info to context for handlers
		ctx.Set("userID", claims.UserID)

		ctx.Next()
	}
}
