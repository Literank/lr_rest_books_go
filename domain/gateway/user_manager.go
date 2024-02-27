package gateway

import (
	"context"

	"literank.com/rest-books/domain/model"
)

// UserManager manages users
type UserManager interface {
	CreateUser(ctx context.Context, u *model.User) (uint, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
}

// PermissionManager manage user permissions by tokens
type PermissionManager interface {
	GenerateToken(userID uint, email string, perm model.UserPermission) (string, error)
	HasPermission(tokenResult string, perm model.UserPermission) (bool, error)
}
