package services

import (
	"context"

	"github.com/nas03/scholar-ai/backend/global"
	"github.com/nas03/scholar-ai/backend/internal/repositories"
	errMessage "github.com/nas03/scholar-ai/backend/pkg/errors"
	"github.com/nas03/scholar-ai/backend/pkg/response"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	Login(ctx context.Context, email, password string) int
}

type AuthService struct {
	userRepo repositories.IUserRepository
}

func NewAuthService(userRepo repositories.IUserRepository) IAuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

func (s *AuthService) Login(ctx context.Context, email, password string) int {
	if email == "" {
		global.Log.Warn(errMessage.ErrInvalidEmail.Error(), zap.String("email", email))
		return response.CodeUserInvalidEmail
	}
	if password == "" {
		global.Log.Warn(errMessage.ErrEmptyPassword.Error(), zap.String("password", password))
		return response.CodeUserInvalidPassword
	}

	// Check if user existed
	user, err := s.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return response.CodeUserNotFound
	}

	// Verify user's password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return response.CodeInvalidCredentials
	}

	// TODO: Generate an access_token and refresh_token
	// TODO: Send access_token(local_storage) & refresh_token(cookies)
	return response.CodeSuccess
}
