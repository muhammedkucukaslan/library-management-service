package loan

import (
	"context"
	"time"

	"github.com/muhammedkucukaslan/library-management-service/domain"
)

type GetLoansRequest struct{}

type Loan struct {
	ID         int       `json:"id"`
	BookID     int       `json:"book_id"`
	UserID     int       `json:"user_id"`
	Status     string    `json:"status"`
	StartedAt  time.Time `json:"started_at"`
	DueAt      time.Time `json:"due_at"`
	ReturnedAt time.Time `json:"returned_at"`
}

type GetLoansResponse []Loan

type GetLoanHandler struct {
	repo Repository
}

func NewGetLoanHandler(repo Repository) *GetLoanHandler {
	return &GetLoanHandler{repo: repo}
}

func (h *GetLoanHandler) Handle(ctx context.Context, req *GetLoansRequest) (*GetLoansResponse, error) {

	role := domain.GetRole(ctx)
	var loans *GetLoansResponse
	var err error

	if role == "admin" || role == "moderator" {
		loans, err = h.repo.GetAllLoans(ctx)
		if err != nil {
			return nil, err
		}
	}

	if role == "user" {
		userID := domain.GetUserID(ctx)
		loans, err = h.repo.GetUserLoans(ctx, userID)
		if err != nil {
			return nil, err
		}
	}

	return loans, nil
}
