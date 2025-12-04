package services

import (
	"context"
	"time"

	"github.com/nas03/scholar-ai/backend/global"
	"github.com/nas03/scholar-ai/backend/internal/helper"
	"github.com/nas03/scholar-ai/backend/internal/models"
	"github.com/nas03/scholar-ai/backend/internal/repositories"
	errMessage "github.com/nas03/scholar-ai/backend/pkg/errors"
	"github.com/nas03/scholar-ai/backend/pkg/response"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	Login(ctx context.Context, email, password string) (*models.AuthTokenPair, int)
}

type AuthService struct {
	userRepo repositories.IUserRepository
}

func NewAuthService(userRepo repositories.IUserRepository) IAuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

func (s *AuthService) Login(ctx context.Context, email, password string) (*models.AuthTokenPair, int) {
	if email == "" {
		global.Log.Warn(errMessage.ErrInvalidEmail.Error(), zap.String("email", email))
		return nil, response.CodeUserInvalidEmail
	}
	if password == "" {
		global.Log.Warn(errMessage.ErrEmptyPassword.Error(), zap.String("password", password))
		return nil, response.CodeUserInvalidPassword
	}

	// Check if user existed
	user, err := s.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, response.CodeUserNotFound
	}

	// Verify user's password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, response.CodeInvalidCredentials
	}

	// TODO: Generate an access_token and refresh_token
	// TODO: Send access_token(local_storage) & refresh_token(cookies)
	jwtHelper := helper.NewJWTHelper()
	claim := map[string]interface{}{
		"UserID": user.UserID,
		"Email":  user.Email,
	}
	accessToken, err := jwtHelper.GenerateAuthToken(ctx, claim, 24*time.Hour)
	if err != nil {
		global.Log.Error("Failed to generate access token", zap.Error(err))
		return nil, response.CodeServerBusy
	}
	refreshToken, err := jwtHelper.GenerateAuthToken(ctx, claim, 30*24*time.Hour)
	return &models.AuthTokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, response.CodeSuccess
}
