package domain

import "context"

func IsPermissionGranted(ctx context.Context) bool {
	return GetRole(ctx) == "moderator" || GetRole(ctx) == "admin"
}
