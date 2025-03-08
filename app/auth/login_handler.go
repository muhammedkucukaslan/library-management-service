package auth

import (
	"context"
	"time"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct{}

type LoginHandler struct {
	repo Repository
	ts   TokenService
	cs   CookieService
}

func NewLoginHandler(repo Repository, ts TokenService, cs CookieService) *LoginHandler {
	return &LoginHandler{repo: repo, ts: ts, cs: cs}
}

func (h *LoginHandler) Handle(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	user, err := h.repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	err = user.ValidatePassword(req.Password)
	if err != nil {
		return nil, err
	}

	token, err := h.ts.GenerateToken(user.ID, user.Role)
	if err != nil {
		return nil, err
	}

	err = h.cs.SetAuthCookies(ctx, token, time.Now().Add(time.Hour*24))
	if err != nil {
		return nil, err
	}

	return &LoginResponse{}, nil
}
