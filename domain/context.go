package domain

import "context"

type ContextKey string

const (
	UserIDKey ContextKey = "userID"
	RoleKey   ContextKey = "role"
	TokenKey  ContextKey = "token"
)

func GetUserID(ctx context.Context) int {
	return ctx.Value(UserIDKey).(int)

}

func GetRole(ctx context.Context) string {
	return ctx.Value(RoleKey).(string)
}
