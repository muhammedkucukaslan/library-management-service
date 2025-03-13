package loan

import (
	"context"
	"errors"
)

type ReturnBookRequest struct {
	LoanID int `json:"id" params:"id"`
}

type ReturnBookResponse struct{}

type ReturnBookHandler struct {
	repo Repository
}

func NewReturnBookHandler(repo Repository) *ReturnBookHandler {
	return &ReturnBookHandler{repo: repo}
}

func (h *ReturnBookHandler) Handle(ctx context.Context, req *ReturnBookRequest) (*ReturnBookResponse, error) {

	loan, err := h.repo.GetLoan(ctx, req.LoanID)
	if err != nil {
		return nil, err
	}

	if loan == nil {
		return nil, errors.New("loan not found")
	}

	if loan.IsReturned() {
		return nil, errors.New("loan already returned")
	}

	if loan.IsOverdue() {
		return nil, errors.New("loan is overdue")
	}

	loan.Return()
	if err := h.repo.UpdateLoan(ctx, loan); err != nil {
		return nil, err
	}
	return &ReturnBookResponse{}, nil
}
