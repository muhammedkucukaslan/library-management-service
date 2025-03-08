package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/muhammedkucukaslan/library-management-service/app/book"
	"github.com/muhammedkucukaslan/library-management-service/domain"
)

func (r *Repository) GetBook(ctx context.Context, id int) (*book.GetBookResponse, error) {
	query := `select  
	b.id,
	b.title,
	CONCAT (a.first_name, ' ', a.second_name) AS author,
	c.title as category,
	b.isbn,
	b.publisher, 
	b.edition,
	b.description,
	bs.total_quantity,
	bs.available_quantity,
	b.created_at,
	b.updated_at
from 
	books b
	join book_categories bc on b.id= bc.book_id
	join categories	 c on bc.category_id = c.id
	join authors a on b.author_id = a.id 
	join book_stocks bs on bs.book_id = b.id
	where b.id = $1`
	row := r.db.QueryRowContext(ctx, query, id)

	if err := row.Err(); err != nil {
		return nil, err
	}

	var book book.GetBookResponse
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Category, &book.ISBN, &book.Publisher, &book.Edition, &book.Description, &book.TotalQuantity, &book.AvailableQuantity, &book.CreatedAt, &book.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no book found with id %d", id)
		}
		return nil, err
	}
	return &book, nil
}

func (r *Repository) CreateBook(ctx context.Context, book *domain.Book) (int, error) {
	query := `INSERT INTO books (title, author_id, isbn, description, publisher, edition) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	err := r.db.QueryRowContext(ctx, query, book.Title, book.AuthorID, book.ISBN, book.Description, book.Publisher, book.Edition).Scan(&book.ID)

	if err != nil {
		return 0, err
	}
	return book.ID, nil
}

func (r *Repository) CreateStock(ctx context.Context, stock *domain.BookStock) error {
	query := `INSERT INTO book_stocks (book_id, total_quantity, available_quantity) VALUES ($1, $2, $3)`

	_, err := r.db.ExecContext(ctx, query, stock.BookID, stock.Total, stock.Available)

	return err
}

func (r *Repository) CreateBookCategory(ctx context.Context, bookID, categoryID int) error {
	query := `INSERT INTO book_categories (book_id, category_id) VALUES ($1, $2)`

	_, err := r.db.ExecContext(ctx, query, bookID, categoryID)

	return err
}

func (r *Repository) DeleteBook(ctx context.Context, id int) error {
	query := `DELETE FROM books WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, id)

	return err
}
