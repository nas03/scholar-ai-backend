package services

import "context"

type IAuthService interface {
	Login(ctx context.Context, email, password string) int
}
