package author

import (
	"context"
	"errors"
	"time"

	"github.com/muhammedkucukaslan/library-management-service/domain"
)

type CreateAuthorRequest struct {
	FirstName  string    `json:"first_name"`
	SecondName string    `json:"second_name"`
	Bio        string    `json:"bio"`
	BirthDate  time.Time `json:"birth_date"`
}

type CreateAuthorResponse struct{}

type CreateAuthorHandler struct {
	repo Repository
}

func NewCreateAuthorHandler(repo Repository) *CreateAuthorHandler {
	return &CreateAuthorHandler{repo: repo}
}

func (h *CreateAuthorHandler) Handle(ctx context.Context, req *CreateAuthorRequest) (*CreateAuthorResponse, error) {
	role := domain.GetRole(ctx)

	if role != "admin" && role != "moderator" {
		return nil, errors.New("unauthorized")
	}

	author := &domain.Author{
		FirstName:  req.FirstName,
		SecondName: req.SecondName,
		Bio:        req.Bio,
		BirthDate:  req.BirthDate,
	}
	err := h.repo.CreateAuthor(ctx, author)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
