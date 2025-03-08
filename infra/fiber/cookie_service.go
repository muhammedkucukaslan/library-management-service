package fiber

import (
	"context"
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
)

type FiberCookieService struct{}

func NewFiberCookieService() *FiberCookieService {
	return &FiberCookieService{}
}

func (s *FiberCookieService) SetAuthCookies(ctx context.Context, token string, expiresAt time.Time) error {

	fiberCtx, ok := ctx.Value(FiberContextKey{}).(*fiber.Ctx)
	if !ok || fiberCtx == nil {
		return errors.New("fiber context not found or invalid")
	}

	fiberCtx.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    token,
		Expires:  expiresAt,
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
		Path:     "/",
	})
	return nil
}
