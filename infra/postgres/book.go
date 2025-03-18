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

func (r *Repository) GetAllBooks(ctx context.Context, page, limit int) (*book.GetBooksResponse, error) {
	query := `SELECT
	b.id,
	b.title,
	(
		SELECT
			CONCAT (a.first_name, ' ', a.second_name)
		FROM
			authors a
		WHERE
			a.id = b.author_id
	) AS author,
	(
		SELECT
			c.title
		FROM
			categories c
			JOIN book_categories bc ON bc.category_id = c.id
		WHERE
			bc.book_id = b.id
		LIMIT
			1
	) AS category,
	b.publisher,
	b.description
	FROM books b
	OFFSET $1 LIMIT $2`

	if limit == 0 {
		limit = 10
	}

	if page == 0 {
		page = 1
	}

	rows, err := r.db.QueryContext(ctx, query, (page-1)*limit, limit)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var books book.GetBooksResponse
	for rows.Next() {
		var book book.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Category, &book.Publisher, &book.Description)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return &books, nil
}

func (r *Repository) GetBooksByAuthor(ctx context.Context, authorID int) (*book.GetBooksResponse, error) {
	query := `SELECT
	b.id,
	b.title,
	(
		SELECT
			CONCAT (a.first_name, ' ', a.second_name)
		FROM
			authors a
		WHERE
			a.id = b.author_id
	) AS author,
	(
		SELECT
			c.title
		FROM
			categories c
			JOIN book_categories bc ON bc.category_id = c.id
		WHERE
			bc.book_id = b.id
		LIMIT
			1
	) AS category,
	b.publisher,
	b.description
	FROM books b
	WHERE b.author_id = $1`

	rows, err := r.db.QueryContext(ctx, query, authorID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var books book.GetBooksResponse
	for rows.Next() {
		var book book.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Category, &book.Publisher, &book.Description)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return &books, nil
}

func (r *Repository) GetBooksByCategory(ctx context.Context, categoryID int) (*book.GetBooksResponse, error) {
	query := `SELECT
	b.id,
	b.title,
	(
		SELECT
			CONCAT (a.first_name, ' ', a.second_name)
		FROM
			authors a
		WHERE
			a.id = b.author_id
	) AS author,
	(
		SELECT
			c.title
		FROM
			categories c
			JOIN book_categories bc ON bc.category_id = c.id
		WHERE
			bc.book_id = b.id
		LIMIT
			1
	) AS category,
	b.publisher,
	b.description
	FROM books b
	JOIN book_categories bc ON bc.book_id = b.id
	WHERE bc.category_id = $1`

	rows, err := r.db.QueryContext(ctx, query, categoryID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var books book.GetBooksResponse
	for rows.Next() {
		var book book.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Category, &book.Publisher, &book.Description)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return &books, nil

}
