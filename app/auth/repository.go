package auth

import (
	"context"

	"github.com/muhammedkucukaslan/library-management-service/domain"
)

type Repository interface {
	CreateUser(ctx context.Context, user *domain.User) (int, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
}
