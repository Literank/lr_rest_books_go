package executor

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"math/rand"
	"time"

	"literank.com/rest-books/application/dto"
	"literank.com/rest-books/domain/gateway"
	"literank.com/rest-books/domain/model"
)

const (
	saltLen          = 4
	errEmptyEmail    = "empty email"
	errEmptyPassword = "empty password"
)

type UserOperator struct {
	userManager gateway.UserManager
	permManager gateway.PermissionManager
}

func NewUserOperator(u gateway.UserManager, p gateway.PermissionManager) *UserOperator {
	return &UserOperator{userManager: u, permManager: p}
}

// CreateUser creates a new user
func (u *UserOperator) CreateUser(ctx context.Context, uc *dto.UserCredential) (*dto.User, error) {
	if uc.Email == "" {
		return nil, errors.New(errEmptyEmail)
	}
	if uc.Password == "" {
		return nil, errors.New(errEmptyPassword)
	}
	salt := randomString(saltLen)
	user := &model.User{
		Email:    uc.Email,
		Password: sha1Hash(uc.Password + salt),
		Salt:     salt,
	}
	uid, err := u.userManager.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return &dto.User{
		ID:    uid,
		Email: uc.Email,
	}, nil
}

// SignIn signs an user in
func (u *UserOperator) SignIn(ctx context.Context, email, password string) (*dto.UserToken, error) {
	if email == "" {
		return nil, errors.New(errEmptyEmail)
	}
	if password == "" {
		return nil, errors.New(errEmptyPassword)
	}
	user, err := u.userManager.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	passwordHash := sha1Hash(password + user.Salt)
	if user.Password != passwordHash {
		return nil, errors.New("wrong password")
	}
	token, err := u.permManager.GenerateToken(user.ID, user.Email, calcPerm(user.IsAdmin))
	if err != nil {
		return nil, err
	}

	return &dto.UserToken{
		User: dto.User{
			ID:    user.ID,
			Email: user.Email,
		},
		Token: token,
	}, nil
}

// HasPermission checks if user has the given permission.
func (u *UserOperator) HasPermission(tokenResult string, perm model.UserPermission) (bool, error) {
	return u.permManager.HasPermission(tokenResult, perm)
}

func calcPerm(isAdmin bool) model.UserPermission {
	perm := model.PermUser
	if isAdmin {
		perm = model.PermAdmin
	}
	return perm
}

func randomString(length int) string {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[random.Intn(len(charset))]
	}
	return string(result)
}

func sha1Hash(input string) string {
	h := sha1.New()
	h.Write([]byte(input))
	hashBytes := h.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	return hashString
}
