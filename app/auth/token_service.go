package auth

type TokenPayload struct {
	UserID int    `json:"user_id"`
	Role   string `json:"role"`
}

type TokenService interface {
	GenerateToken(userID int, role string) (string, error)
	ValidateToken(token string) (TokenPayload, error)
}
