package fiber

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

type FiberContextKey struct{}

func ContextMiddleware(c *fiber.Ctx) error {
	ctx := context.WithValue(c.UserContext(), FiberContextKey{}, c)
	c.SetUserContext(ctx)
	return c.Next()
}
