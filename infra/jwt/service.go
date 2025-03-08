package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/muhammedkucukaslan/library-management-service/app/auth"
)

type TokenService struct {
	secretKey     string
	tokenDuration time.Duration
}

func NewTokenService(secretKey string, tokenDuration time.Duration) *TokenService {
	return &TokenService{secretKey: secretKey, tokenDuration: tokenDuration}
}

func (s *TokenService) GenerateToken(userID int, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"role":    role,
	})

	return token.SignedString([]byte(s.secretKey))
}

func (s *TokenService) ValidateToken(tokenString string) (auth.TokenPayload, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verify signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.secretKey), nil
	})
	if err != nil {
		return auth.TokenPayload{}, err
	}

	if !token.Valid {
		return auth.TokenPayload{}, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return auth.TokenPayload{}, fmt.Errorf("invalid token claims")
	}

	userIDClaim, exists := claims["user_id"]
	if !exists || userIDClaim == nil {
		return auth.TokenPayload{}, fmt.Errorf("user_id claim is missing or nil")
	}

	roleClaim, exists := claims["role"]
	if !exists || roleClaim == nil {
		return auth.TokenPayload{}, fmt.Errorf("role claim is missing or nil")
	}

	userIDFloat, ok := userIDClaim.(float64)
	if !ok {
		return auth.TokenPayload{}, fmt.Errorf("invalid user_id type")
	}
	userID := int(userIDFloat)

	role, ok := roleClaim.(string)
	if !ok {
		return auth.TokenPayload{}, fmt.Errorf("invalid role type")
	}

	return auth.TokenPayload{
		UserID: userID,
		Role:   role,
	}, nil
}
