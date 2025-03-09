package book

import (
	"context"

	"github.com/muhammedkucukaslan/library-management-service/domain"
)

type Repository interface {
	CreateBook(ctx context.Context, book *domain.Book) (int, error)
	CreateStock(ctx context.Context, stock *domain.BookStock) error
	CreateBookCategory(ctx context.Context, bookID, categoryID int) error
	GetBook(ctx context.Context, id int) (*GetBookResponse, error)
	DeleteBook(ctx context.Context, id int) error
	// GetBooksByCategory(ctx context.Context, categoryID int) ([]*domain.Book, error)
	// GetCategory(ctx context.Context, id int) (*domain.BookCategory, error)
	// GetBook(ctx context.Context, id int) (*domain.Book, error)
}
