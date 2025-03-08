package auth

import (
	"context"
	"time"

	"github.com/muhammedkucukaslan/library-management-service/domain"
)

type SignupRequest struct {
	FirstName  string    `json:"first_name"`
	SecondName string    `json:"second_name"`
	Role       string    `json:"role"`
	Password   string    `json:"password"`
	Email      string    `json:"email"`
	Address    string    `json:"address"`
	BirthDate  time.Time `json:"birth_date"`
}

type SignupResponse struct{}

type SignupHandler struct {
	repo Repository
	ts   TokenService
	cs   CookieService
}

func NewSignupHandler(repo Repository, ts TokenService, cs CookieService) *SignupHandler {
	return &SignupHandler{repo: repo, ts: ts, cs: cs}
}

func (h *SignupHandler) Handle(ctx context.Context, req *SignupRequest) (*SignupResponse, error) {
	user := domain.NewUser(req.FirstName, req.SecondName, req.Role, req.Password, req.Email, req.Address, req.BirthDate)

	userID, err := h.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	token, err := h.ts.GenerateToken(userID, user.Role)
	if err != nil {
		return nil, err
	}

	err = h.cs.SetAuthCookies(ctx, token, time.Now().Add(time.Hour*24))
	if err != nil {
		return nil, err
	}

	return &SignupResponse{}, nil
}
