package dto

// UserCredential represents the user's sign-in email and password
type UserCredential struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

// User is used as result of a successful sign-in
type User struct {
	ID    uint   `json:"id,omitempty"`
	Email string `json:"email,omitempty"`
}

// UserToken is an combination of User struct and token field
type UserToken struct {
	User  User   `json:"user,omitempty"`
	Token string `json:"token,omitempty"`
}
