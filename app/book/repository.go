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
	GetAllBooks(ctx context.Context, page, limit int) (*GetBooksResponse, error)
	GetBooksByAuthor(ctx context.Context, authorID int) (*GetBooksResponse, error)
	GetBooksByCategory(ctx context.Context, categoryID int) (*GetBooksResponse, error)
}
