package bookcategory

import (
	"context"
	"errors"

	"github.com/muhammedkucukaslan/library-management-service/domain"
)

type DeleteCategoryRequest struct {
	ID int `json:"id"`
}

type DeleteCategoryResponse struct{}

type DeleteCategoryHandler struct {
	repo Repository
}

func NewDeleteCategoryHandler(repo Repository) *DeleteCategoryHandler {
	return &DeleteCategoryHandler{repo: repo}
}

func (h *DeleteCategoryHandler) Handle(ctx context.Context, req *DeleteCategoryRequest) (*DeleteCategoryResponse, error) {
	role := domain.GetRole(ctx)

	if role != "admin" && role != "moderator" {
		return nil, errors.New("unauthorized")
	}

	if err := h.repo.DeleteCategory(ctx, req.ID); err != nil {
		return nil, err
	}

	return nil, nil
}
