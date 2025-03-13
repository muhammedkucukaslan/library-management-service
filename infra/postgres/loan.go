package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/muhammedkucukaslan/library-management-service/domain"
)

func (r *Repository) CreateLoan(ctx context.Context, loan *domain.Loan) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	insertQuery := `INSERT INTO loans (book_id, user_id, started_at, due_date) 
	                VALUES ($1, $2, $3, $4)`
	_, err = tx.ExecContext(ctx, insertQuery, loan.BookID, loan.UserID, loan.StartedDate, loan.DueDate)
	if err != nil {
		return err
	}

	updateQuery := `UPDATE book_stocks 
	                SET available_quantity = available_quantity - 1 
	                WHERE book_id = $1`
	_, err = tx.ExecContext(ctx, updateQuery, loan.BookID)
	if err != nil {
		return err
	}

	return tx.Commit()
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
			  AND l.status = 'BORROWED'
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

func (r Repository) GetLoan(ctx context.Context, loanID int) (*domain.Loan, error) {
	query := `SELECT id, book_id,  status
	          FROM loans 
	          WHERE id = $1`

	row := r.db.QueryRowContext(ctx, query, loanID)

	var loan domain.Loan
	if err := row.Scan(&loan.ID, &loan.BookID, &loan.Status); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &loan, nil
}

func (r Repository) UpdateLoan(ctx context.Context, loan *domain.Loan) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := `UPDATE loans
	          SET returned_at = $1, status = $2, updated_at = NOW()
	          WHERE id = $3`
	_, err = tx.ExecContext(ctx, query, loan.ReturnedDate, loan.Status, loan.ID)
	if err != nil {
		return err
	}

	updateQuery := `UPDATE book_stocks
	                SET available_quantity = available_quantity + 1,
					update_at = NOW()
	                WHERE book_id = $1`
	_, err = tx.ExecContext(ctx, updateQuery, loan.BookID)
	if err != nil {
		return err
	}

	return tx.Commit()
}
