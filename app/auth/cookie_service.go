package auth

import (
	"context"
	"time"
)

type CookieService interface {
	SetAuthCookies(ctx context.Context, token string, expiresAt time.Time) error
}
