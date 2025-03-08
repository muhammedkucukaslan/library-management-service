package bookcategory

import (
	"context"
	"errors"

	"github.com/muhammedkucukaslan/library-management-service/domain"
)

type CreateCategoryRequest struct {
	Title string `json:"title"`
}

type CreateCategoryResponse struct{}

type CreateCategoryHandler struct {
	repo Repository
}

func NewCreateCategoryHandler(repo Repository) *CreateCategoryHandler {
	return &CreateCategoryHandler{repo: repo}
}

func (h *CreateCategoryHandler) Handle(ctx context.Context, req *CreateCategoryRequest) (*CreateCategoryResponse, error) {
	role := domain.GetRole(ctx)

	if role != "admin" && role != "moderator" {
		return nil, errors.New("unauthorized")
	}

	if err := h.repo.CreateCategory(ctx, req.Title); err != nil {
		return nil, err
	}

	return nil, nil
}
