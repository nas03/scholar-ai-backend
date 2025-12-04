package helper

import (
	"context"
	"errors"
	"fmt"
	"maps"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nas03/scholar-ai/backend/global"
	"github.com/nas03/scholar-ai/backend/internal/models"
	"go.uber.org/zap"
)

type IJWTHelper interface {
	GenerateAuthToken(ctx context.Context, data any, exp time.Duration) (string, error)
	GetClaims(token string) (map[string]interface{}, error)
	ValidateAuthToken(ctx context.Context, token string) (*models.AuthTokenClaim, error)
}

type JWTHelper struct{}

func NewJWTHelper() IJWTHelper {
	return &JWTHelper{}
}

func (h *JWTHelper) GenerateAuthToken(ctx context.Context, data any, exp time.Duration) (string, error) {
	// Create claims with default expiration and custom data
	claims := jwt.MapClaims{
		"exp": time.Now().Add(exp).Unix(),
		"iat": time.Now().Unix(),
	}

	if dataMap, ok := data.(map[string]interface{}); ok {
		maps.Copy(claims, dataMap)
	} else {
		claims["data"] = data
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	privateKey, err := NewRSAHelper().LoadPrivateKey(ctx, "")
	if err != nil {
		return "", err
	}
	return token.SignedString(privateKey)

}

func (h *JWTHelper) GetClaims(token string) (map[string]interface{}, error) {
	return nil, nil
}

func (h *JWTHelper) ValidateAuthToken(ctx context.Context, token string) (*models.AuthTokenClaim, error) {
	rsaHelper := NewRSAHelper()
	privateKey, err := rsaHelper.LoadPrivateKey(ctx, "")
	if err != nil {
		global.Log.Error("Failed loading private key", zap.Error(err))
		return nil, err
	}
	publicKey := NewRSAHelper().GetPublicKey(ctx, privateKey)

	parsedToken, err := jwt.ParseWithClaims(token, &models.AuthTokenClaim{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return publicKey, nil
	})

	if err != nil {
		global.Log.Fatal("Token verification failed: %v", zap.Error(err))
		return nil, err
	}

	// Access the claims if valid
	if claims, ok := parsedToken.Claims.(*models.AuthTokenClaim); ok && parsedToken.Valid {
		return claims, nil
	} else {
		return nil, errors.New("token is invalid")
	}
}
