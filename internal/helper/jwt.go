package helper

import (
	"context"
	"maps"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTDefaultClaim struct {
	jwt.StandardClaims
}
type IJWTHelper interface {
	GenerateAuthToken(ctx context.Context, data any, exp time.Duration) (string, error)
	GetClaims(token string) (map[string]interface{}, error)
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

	// TODO: Load RSA private key and sign the token
	// privateKey, err := loadPrivateKey()
	// if err != nil {
	//     return "", err
	// }
	// return token.SignedString(privateKey)

	_ = token // Placeholder - token will be signed when private key is available
	return "", nil
}

func (h *JWTHelper) GetClaims(token string) (map[string]interface{}, error) {
	return nil, nil
}
