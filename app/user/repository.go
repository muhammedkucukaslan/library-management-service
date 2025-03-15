package user

import (
	"context"

	"github.com/muhammedkucukaslan/library-management-service/domain"
)

type Repository interface {
	GetUserByID(ctx context.Context, id int) (*domain.User, error)
	CheckEmail(ctx context.Context, email string) error
	DeleteUser(ctx context.Context, id int) error
	UpdateUser(ctx context.Context, user *domain.User) error
	GetUserPunishments(ctx context.Context, id int) ([]Punishment, error)
	GetCurrentUserLoans(ctx context.Context, id int) ([]Loan, error)
}
