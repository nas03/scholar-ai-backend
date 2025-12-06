// Package consts
package consts

import "time"

var (
	UserAccountStatus = struct {
		INACTIVE int8
		ACTIVE   int8
	}{
		INACTIVE: 0,
		ACTIVE:   1,
	}

	Flag = struct {
		TRUE  int8
		FALSE int8
	}{
		TRUE:  1,
		FALSE: 0,
	}
	REDIS_OTP_EXPIRATION     = 60 * time.Second // 1 minute
	REDIS_DEFAULT_EXPIRATION = 60 * time.Minute // 1 hour

	REFRESH_TOKEN_COOKIE = "REFRESH_TOKEN"
)
