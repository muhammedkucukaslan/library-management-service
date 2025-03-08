package book

import (
	"context"
	"time"
)

type GetBookRequest struct {
	ID int `json:"id" param:"id"`
}

type GetBookResponse struct {
	ID                int       `json:"id"`
	Title             string    `json:"title"`
	Author            string    `json:"author"`
	Category          string    `json:"category"`
	ISBN              string    `json:"isbn"`
	TotalQuantity     int       `json:"total_quantity"`
	AvailableQuantity int       `json:"available_quantity"`
	Publisher         string    `json:"publisher"`
	Edition           int       `json:"edition"`
	Description       string    `json:"description"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type GetBookHandler struct {
	repo Repository
}

func NewGetBookHandler(repo Repository) *GetBookHandler {
	return &GetBookHandler{repo: repo}
}

func (h *GetBookHandler) Handle(ctx context.Context, req *GetBookRequest) (*GetBookResponse, error) {

	book, err := h.repo.GetBook(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &GetBookResponse{
		ID:                book.ID,
		Title:             book.Title,
		Author:            book.Author,
		Category:          book.Category,
		ISBN:              book.ISBN,
		TotalQuantity:     book.TotalQuantity,
		AvailableQuantity: book.AvailableQuantity,
		Publisher:         book.Publisher,
		Edition:           book.Edition,
		Description:       book.Description,
	}, nil

}
