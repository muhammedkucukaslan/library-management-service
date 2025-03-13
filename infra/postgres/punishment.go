package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/muhammedkucukaslan/library-management-service/app/punishment"
	"github.com/muhammedkucukaslan/library-management-service/domain"
)

func (r *Repository) PunishUser(ctx context.Context, punishment *domain.Punishment) error {
	fmt.Println("punishing user", punishment)
	_, err := r.db.ExecContext(ctx, "INSERT INTO punishments (user_id, punisher_id, reason, end_at) VALUES ($1, $2, $3, $4)", punishment.UserID, punishment.PunisherID, punishment.Reason, punishment.EndDate)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) PunishOverdueLoans(ctx context.Context) ([]int, error) {

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	query := `SELECT l.id, l.user_id, l.book_id, l.status
			  FROM loans l
			  WHERE l.due_at < NOW() AND l.status = 'BORROWED'`

	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var overdueLoans []punishment.OverdueLoan
	for rows.Next() {
		var loan punishment.OverdueLoan
		err := rows.Scan(&loan.LoanID, &loan.UserID, &loan.BookID, &loan.Status)
		if err != nil {
			return nil, err
		}
		overdueLoans = append(overdueLoans, loan)
	}

	fmt.Println(overdueLoans)

	for _, loan := range overdueLoans {
		_, err := tx.ExecContext(ctx, "UPDATE loans SET status = 'OVERDUE' WHERE id = $1", loan.LoanID)
		if err != nil {
			return nil, err
		}

		punishment := domain.NewPunishment(loan.UserID, 22, "Overdue loan", time.Now().AddDate(0, 0, 15))
		fmt.Println(punishment)
		err = r.PunishUser(ctx, punishment)
		if err != nil {
			fmt.Println("AAA", err)
			return nil, err
		}
	}

	var userIDs []int
	for _, loan := range overdueLoans {
		userIDs = append(userIDs, loan.UserID)
	}

	if tx.Commit() != nil {
		return nil, err
	}

	return userIDs, nil
}
