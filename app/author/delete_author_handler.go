package author

import (
	"context"
	"errors"

	"github.com/muhammedkucukaslan/library-management-service/domain"
)

type DeleteAuthorRequest struct {
	ID int `json:"id"`
}

type DeleteAuthorResponse struct{}

type DeleteAuthorHandler struct {
	repo Repository
}

func NewDeleteAuthorHandler(repo Repository) *DeleteAuthorHandler {
	return &DeleteAuthorHandler{repo: repo}
}

func (h *DeleteAuthorHandler) Handle(ctx context.Context, req *DeleteAuthorRequest) (*DeleteAuthorResponse, error) {

	role := domain.GetRole(ctx)

	if role != "admin" && role != "moderator" {
		return nil, errors.New("unauthorized")
	}

	err := h.repo.DeleteAuthor(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
