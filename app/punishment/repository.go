package punishment

import (
	"context"

	"github.com/muhammedkucukaslan/library-management-service/domain"
)

type Repository interface {
	PunishUser(ctx context.Context, punishment *domain.Punishment) error
	PunishOverdueLoans(ctx context.Context) ([]int, error)
}
