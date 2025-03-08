package user

import (
	"context"
	"errors"

	"github.com/muhammedkucukaslan/library-management-service/domain"
)

type DeleteUserRequest struct {
	ID int `json:"id"`
}

type DeleteUserResponse struct{}

type DeleteUserHandler struct {
	repo Repository
}

func NewDeleteUserHandler(repo Repository) *DeleteUserHandler {
	return &DeleteUserHandler{repo: repo}
}

func (h *DeleteUserHandler) Handle(ctx context.Context, req *DeleteUserRequest) (*DeleteUserResponse, error) {

	role := domain.GetRole(ctx)

	if role != "admin" && role != "user" {
		return nil, errors.New("unauthorized")
	}

	err := h.repo.DeleteUser(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &DeleteUserResponse{}, nil
}
