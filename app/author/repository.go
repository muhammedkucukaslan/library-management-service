package author

import (
	"context"

	"github.com/muhammedkucukaslan/library-management-service/domain"
)

type Repository interface {
	CreateAuthor(ctx context.Context, author *domain.Author) error
	DeleteAuthor(ctx context.Context, id int) error
}
