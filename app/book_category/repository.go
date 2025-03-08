package bookcategory

import (
	"context"
)

type Repository interface {
	CreateCategory(ctx context.Context, title string) error
	DeleteCategory(ctx context.Context, id int) error
}
