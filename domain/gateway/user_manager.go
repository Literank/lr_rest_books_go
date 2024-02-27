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
