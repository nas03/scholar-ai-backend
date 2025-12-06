package controllers

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nas03/scholar-ai/backend/internal/consts"
	"github.com/nas03/scholar-ai/backend/internal/models"
	"github.com/nas03/scholar-ai/backend/internal/services"
	"github.com/nas03/scholar-ai/backend/pkg/response"
)

type AuthController struct {
	authService services.IAuthService
}

func NewAuthController(authService services.IAuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (c *AuthController) Login(ctx *gin.Context) {
	var req models.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(ctx, response.CodeInvalidParams, err.Error())
		return
	}

	tokenPair, code := c.authService.Login(ctx, req.Email, req.Password)
	if code != response.CodeSuccess {
		response.ErrorResponse(ctx, code, "")
		return
	}
	// TODO: Save device ID and create login session
	// Set token to response
	ctx.Header("Authorization", tokenPair.AccessToken)
	ctx.SetCookie(consts.REFRESH_TOKEN_COOKIE, tokenPair.RefreshToken, int(30*24*time.Hour.Seconds()), "/", "", true, true)
	response.SuccessResponse(ctx, response.CodeSuccess, nil)
}

func (c *AuthController) RotateAuthToken(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	accessToken := strings.TrimPrefix(authHeader, "Bearer ")

	if accessToken == "" {
		response.ErrorResponse(ctx, response.CodeTokenInvalid, "")
		return
	}

	refreshToken, err := ctx.Cookie(consts.REFRESH_TOKEN_COOKIE)
	if err != nil || refreshToken == "" {
		response.ErrorResponse(ctx, response.CodeTokenInvalid, "Invalid refresh token")
		return
	}

	tokenPair, code := c.authService.RotateAuthToken(ctx, accessToken, refreshToken)
	if code != response.CodeSuccess {
		response.ErrorResponse(ctx, code, "")
		return
	}

	ctx.Header("Authorization", tokenPair.AccessToken)
	ctx.SetCookie(consts.REFRESH_TOKEN_COOKIE, tokenPair.RefreshToken, int(30*24*time.Hour.Seconds()), "/", "", true, true)
	response.SuccessResponse(ctx, response.CodeSuccess, nil)
}
