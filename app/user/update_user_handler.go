package user

import (
	"context"
	"errors"
	"time"

	"github.com/muhammedkucukaslan/library-management-service/domain"
)

type UpdateUserRequest struct {
	ID         int       `json:"id"`
	FirstName  string    `json:"first_name,omitempty"`
	SecondName string    `json:"second_name,omitempty"`
	Email      string    `json:"email,omitempty"`
	Address    string    `json:"address,omitempty"`
	BirthDate  time.Time `json:"birth_date,omitzero"`
}

type UpdateUserResponse struct{}

type UpdateUserHandler struct {
	repo Repository
}

func NewUpdateUserHandler(repo Repository) *UpdateUserHandler {
	return &UpdateUserHandler{repo: repo}
}

func (h *UpdateUserHandler) Handle(ctx context.Context, req *UpdateUserRequest) (*UpdateUserResponse, error) {

	role := domain.GetRole(ctx)

	if role != "admin" && role != "user" {
		return nil, errors.New("unauthorized")
	}

	user, err := h.repo.GetUserByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	if req.FirstName != "" {
		user.FirstName = req.FirstName
	}
	if req.SecondName != "" {
		user.SecondName = req.SecondName
	}
	if req.Email != "" {

		if err := h.repo.CheckEmail(ctx, req.Email); err != nil {
			return nil, err
		}
		user.Email = req.Email
	}

	if req.Address != "" {
		user.Address = req.Address
	}
	if !(req.BirthDate.IsZero()) {
		user.BirthDate = req.BirthDate
	}

	user.UpdatedAt = time.Now()

	if err = h.repo.UpdateUser(ctx, user); err != nil {
		return nil, err
	}

	return &UpdateUserResponse{}, nil

}
