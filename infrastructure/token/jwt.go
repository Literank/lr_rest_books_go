/*
Package token includes authorization tokens.
*/
package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"literank.com/rest-books/domain/model"
)

const (
	errInvalidToken  = "invalid token"
	errFailToConvert = "failed to convert token type"
)

// Keeper manages user tokens.
type Keeper struct {
	secretKey   []byte
	expireHours uint
}

// UserClaims includes user info.
type UserClaims struct {
	UserID     uint                 `json:"user_id,omitempty"`
	UserName   string               `json:"user_name,omitempty"`
	Permission model.UserPermission `json:"permission,omitempty"`
	jwt.RegisteredClaims
}

// NewTokenKeeper constructs a new JWT token keeper
func NewTokenKeeper(secretKey string, expireInHours uint) *Keeper {
	return &Keeper{[]byte(secretKey), expireInHours}
}

// GenerateToken generates a new JWT token.
func (t *Keeper) GenerateToken(userID uint, email string, perm model.UserPermission) (string, error) {
	claims := UserClaims{
		userID, email, perm,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(t.expireHours) * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenResult, err := token.SignedString(t.secretKey)
	if err != nil {
		return "", err
	}
	return tokenResult, nil
}

// ExtractToken extracts the token from the signed string.
func (t *Keeper) ExtractToken(tokenResult string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenResult, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return t.secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New(errInvalidToken)
	}
	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, errors.New(errFailToConvert)
	}
	return claims, nil
}

// HasPermission checks if user has the given permission.
func (t *Keeper) HasPermission(tokenResult string, perm model.UserPermission) (bool, error) {
	claims, err := t.ExtractToken(tokenResult)
	if err != nil {
		return false, err
	}
	return claims.Permission >= perm, nil
}
