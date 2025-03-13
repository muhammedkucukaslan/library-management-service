package punishment

import (
	"context"
	"errors"
	"time"

	"github.com/muhammedkucukaslan/library-management-service/domain"
)

type PunishUserRequest struct {
	UserID  int       `json:"user_id" params:"id"`
	Reason  string    `json:"reason"`
	EndDate time.Time `json:"end_date"`
}

type PunishUserResponse struct{}

type PunishUserHandler struct {
	repo Repository
}

func NewPunishUserHandler(repo Repository) *PunishUserHandler {
	return &PunishUserHandler{repo: repo}
}

func (h *PunishUserHandler) Handle(ctx context.Context, req *PunishUserRequest) (*PunishUserResponse, error) {

	if !domain.IsPermissionGranted(ctx) {
		return nil, errors.New("permission denied")
	}

	punisherID := domain.GetUserID(ctx)

	punishment := domain.NewPunishment(req.UserID, punisherID, req.Reason, req.EndDate)

	if err := h.repo.PunishUser(ctx, punishment); err != nil {
		return nil, err
	}

	return nil, nil
}
