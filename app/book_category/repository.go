package bookcategory

import (
	"context"
)

type Repository interface {
	GetCategories(ctx context.Context) (*GetBookCategoriesResponse, error)
	CreateCategory(ctx context.Context, title string) error
	DeleteCategory(ctx context.Context, id int) error
}
