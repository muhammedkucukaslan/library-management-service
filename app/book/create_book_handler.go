package book

import (
	"context"
	"errors"

	"github.com/muhammedkucukaslan/library-management-service/domain"
)

type CreateBookRequest struct {
	Title       string `json:"title"`
	AuthorID    int    `json:"author_id"`
	Description string `json:"description"`
	ISBN        string `json:"isbn"`
	CategoryID  int    `json:"category_id"`
	TotalStock  int    `json:"total_stock"`
	Publisher   string `json:"publisher"`
	Edition     int    `json:"edition"`
}

type CreateBookResponse struct{}

type CreateBookHandler struct {
	repo Repository
}

func NewCreateBookHandler(repo Repository) *CreateBookHandler {
	return &CreateBookHandler{repo: repo}
}

func (h *CreateBookHandler) Handle(ctx context.Context, req *CreateBookRequest) (*CreateBookResponse, error) {

	if !domain.IsPermissionGranted(ctx) {
		return nil, errors.New("unauthorized")
	}

	book := domain.NewBook(req.Title, req.ISBN, req.Description, req.Publisher, req.AuthorID, req.Edition)

	bookID, err := h.repo.CreateBook(ctx, book)
	if err != nil {
		return nil, err
	}

	stock := domain.NewStock(bookID, req.TotalStock)

	if err := h.repo.CreateStock(ctx, stock); err != nil {
		return nil, err
	}

	if err := h.repo.CreateBookCategory(ctx, bookID, req.CategoryID); err != nil {
		return nil, err
	}

	return &CreateBookResponse{}, nil
}
