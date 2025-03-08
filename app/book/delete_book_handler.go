package book

import (
	"context"
	"errors"

	"github.com/muhammedkucukaslan/library-management-service/domain"
)

type DeleteBookRequest struct {
	ID int `json:"id" param:"id"`
}

type DeleteBookResponse struct{}

type DeleteBookHandler struct {
	repo Repository
}

func NewDeleteBookHandler(repo Repository) *DeleteBookHandler {
	return &DeleteBookHandler{repo: repo}
}

func (h *DeleteBookHandler) Handle(ctx context.Context, req *DeleteBookRequest) (*DeleteBookResponse, error) {

	if !domain.IsPermissionGranted(ctx) {
		return nil, errors.New("unauthorized")
	}

	if err := h.repo.DeleteBook(ctx, req.ID); err != nil {
		return nil, err
	}

	return &DeleteBookResponse{}, nil
}
