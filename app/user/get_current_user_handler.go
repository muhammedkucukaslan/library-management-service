package user

import (
	"context"
	"time"

	"github.com/muhammedkucukaslan/library-management-service/domain"
)

type GetCurrentUserRequest struct{}

type GetCurrentUserResponse struct {
	ID          int          `json:"id"`
	FirstName   string       `json:"first_name"`
	SecondName  string       `json:"second_name"`
	Email       string       `json:"email"`
	Role        string       `json:"role"`
	Address     string       `json:"address"`
	BirthDate   time.Time    `json:"birth_date"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	Loans       []Loan       `json:"loans"`
	Punishments []Punishment `json:"punishments"`
}

type Punishment struct {
	ID        int       `json:"id"`
	StartedAt time.Time `json:"started_at"`
	EndedAt   time.Time `json:"ended_at"`
	Reason    string    `json:"reason"`
}

type Loan struct {
	ID         int       `json:"id"`
	StartedAt  time.Time `json:"started_at"`
	DueAt      time.Time `json:"due_at"`
	ReturnedAt time.Time `json:"returned_at"`
	Status     string    `json:"status"`
	Book       Book      `json:"book"`
}

type Book struct {
	ID       int    `json:"id"`
	ISBN     string `json:"isbn"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Category string `json:"category"`
}

type GetCurrentUserHandler struct {
	repo Repository
}

func NewGetCurrentUserHandler(repo Repository) *GetCurrentUserHandler {
	return &GetCurrentUserHandler{repo: repo}
}

func (h *GetCurrentUserHandler) Handle(ctx context.Context, req *GetCurrentUserRequest) (*GetCurrentUserResponse, error) {
	userID := domain.GetUserID(ctx)

	user, err := h.repo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	punishments, err := h.repo.GetUserPunishments(ctx, userID)
	if err != nil {
		return nil, err
	}

	loans, err := h.repo.GetCurrentUserLoans(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &GetCurrentUserResponse{
		ID:          user.ID,
		FirstName:   user.FirstName,
		SecondName:  user.SecondName,
		Email:       user.Email,
		Role:        user.Role,
		Address:     user.Address,
		BirthDate:   user.BirthDate,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
		Loans:       loans,
		Punishments: punishments,
	}, nil
}
