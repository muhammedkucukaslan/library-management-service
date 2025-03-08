package user

import (
	"context"
	"errors"
	"time"

	"github.com/muhammedkucukaslan/library-management-service/domain"
)

type GetUserRequest struct {
	ID int `json:"id"`
}

type GetUserResponse struct {
	ID         int       `json:"id"`
	FirstName  string    `json:"first_name"`
	SecondName string    `json:"second_name"`
	Email      string    `json:"email"`
	Role       string    `json:"role"`
	Address    string    `json:"address"`
	BirthDate  time.Time `json:"birth_date"`
}

type GetUserHandler struct {
	repo Repository
}

func NewGetUserHandler(repo Repository) *GetUserHandler {
	return &GetUserHandler{repo: repo}
}

func (h *GetUserHandler) Handle(ctx context.Context, req *GetUserRequest) (*GetUserResponse, error) {

	role := domain.GetRole(ctx)

	if role != "admin" && role != "moderator" {
		return nil, errors.New("unauthorized")
	}

	user, err := h.repo.GetUserByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &GetUserResponse{
		ID:         user.ID,
		FirstName:  user.FirstName,
		SecondName: user.SecondName,
		Email:      user.Email,
		Role:       user.Role,
		Address:    user.Address,
		BirthDate:  user.BirthDate,
	}, nil
}
