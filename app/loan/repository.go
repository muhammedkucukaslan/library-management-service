package loan

import (
	"context"

	"github.com/muhammedkucukaslan/library-management-service/domain"
)

type Repository interface {
	GetLoan(ctx context.Context, loanID int) (*domain.Loan, error)
	CreateLoan(ctx context.Context, loan *domain.Loan) error
	CheckBookStocks(ctx context.Context, bookID int) (int, error)
	CheckUserHasAlreadyBorrowed(ctx context.Context, bookID, userID int) error
	CheckUserHasCurrentPunishment(ctx context.Context, userID int) error
	UpdateLoan(ctx context.Context, loan *domain.Loan) error
	GetAllLoans(ctx context.Context) (*GetLoansResponse, error)
	GetUserLoans(ctx context.Context, userID int) (*GetLoansResponse, error)
}
