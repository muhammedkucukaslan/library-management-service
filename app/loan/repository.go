package loan

import (
	"context"

	"github.com/muhammedkucukaslan/library-management-service/domain"
)

type Repository interface {
	CreateLoan(ctx context.Context, loan *domain.Loan) error
	CheckBookStocks(ctx context.Context, bookID int) (int, error)
	CheckUserHasAlreadyBorrowed(ctx context.Context, bookID, userID int) error
}
