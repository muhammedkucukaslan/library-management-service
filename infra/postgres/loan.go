package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/muhammedkucukaslan/library-management-service/domain"
)

func (r *Repository) CreateLoan(ctx context.Context, loan *domain.Loan) error {
	query := `INSERT INTO loans (book_id, user_id, started_at, due_date) VALUES ($1, $2, $3, $4)`
	_, err := r.db.ExecContext(ctx, query, loan.BookID, loan.UserID, loan.StartedDate, loan.DueDate)
	return err
}

func (r *Repository) CheckBookStocks(ctx context.Context, bookID int) (int, error) {
	query := `SELECT available_quantity FROM book_stocks bs WHERE  bs.book_id = $1`

	row := r.db.QueryRowContext(ctx, query, bookID)

	var available_quantity int
	if err := row.Scan(&available_quantity); err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("no book found with id %d", bookID)
		}
		return 0, err
	}

	return available_quantity, nil
}

func (r Repository) CheckUserHasAlreadyBorrowed(ctx context.Context, bookID, userID int) error {
	query := `SELECT EXISTS (
			SELECT 1 
			FROM loans l 
			WHERE l.user_id = $1 
			  AND l.book_id = $2 
			  AND l.status != 'RETURNED'
		  );`

	var exists bool
	err := r.db.QueryRowContext(ctx, query, userID, bookID).Scan(&exists)
	if err != nil {
		return err
	}

	if exists {
		return fmt.Errorf("user has already borrowed the book")
	}
	return nil

}
