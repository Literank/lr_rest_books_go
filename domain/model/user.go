package model

import "time"

// UserPermission represents different levels of user permissions.
type UserPermission uint8

// UserPermission levels
const (
	PermNone UserPermission = iota
	PermUser
	PermAuthor
	PermAdmin
)

// User represents an app user
type User struct {
	ID        uint      `json:"id,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	Salt      string    `json:"salt,omitempty"`
	IsAdmin   bool      `json:"is_admin,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
