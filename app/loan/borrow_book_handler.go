package loan

import (
	"context"
	"errors"

	"github.com/muhammedkucukaslan/library-management-service/domain"
)

type BorrowBookRequest struct {
	BookID int `json:"book_id" param:"book_id"`
}

type BorrowBookResponse struct{}

type BorrowBookHandler struct {
	repo Repository
}

func NewBorrowBookHandler(repo Repository) *BorrowBookHandler {
	return &BorrowBookHandler{repo: repo}
}

func (h *BorrowBookHandler) Handle(ctx context.Context, req *BorrowBookRequest) (*BorrowBookResponse, error) {
	userID := domain.GetUserID(ctx)

	// TODO i want to cover here to OOP. im not happy with this implementation

	if err := h.repo.CheckUserHasCurrentPunishment(ctx, userID); err != nil {
		return nil, err
	}

	if err := h.repo.CheckUserHasAlreadyBorrowed(ctx, req.BookID, userID); err != nil {
		return nil, err
	}

	stock, err := h.repo.CheckBookStocks(ctx, req.BookID)
	if err != nil {
		return nil, err
	}
	if stock == 0 {
		return nil, errors.New("book is out of stock")
	}

	loan := domain.NewLoan(req.BookID, userID)
	if err := h.repo.CreateLoan(ctx, loan); err != nil {
		return nil, err
	}
	return &BorrowBookResponse{}, nil
}
