package domain

import "time"

type AdminID uint
type AuthTokenID string
type AuthToken struct {
	ID        AuthTokenID
	ExpiresAt time.Time
}

type AuthParams struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SaveParams struct {
	ID        AdminID `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Email     string  `json:"email"`
	Password  string  `json:"password"`
}
