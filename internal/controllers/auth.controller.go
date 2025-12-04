package controllers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nas03/scholar-ai/backend/internal/models"
	"github.com/nas03/scholar-ai/backend/internal/services"
	"github.com/nas03/scholar-ai/backend/pkg/response"
)

type AuthController struct {
	authService services.IAuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{
		authService: &authService,
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

	ctx.Header("Authorization", tokenPair.AccessToken)
	ctx.SetCookie("REFRESH_TOKEN", tokenPair.RefreshToken, int(30*24*time.Hour.Seconds()), "/", "", true, true)
	response.SuccessResponse(ctx, response.CodeSuccess, tokenPair)
}
