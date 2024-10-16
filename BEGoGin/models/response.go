package models

import "time"

//USER APIS
type CreateUserResponse struct {
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	// Token     string    `json:"token"`
	// ExpiredAt time.Time `json:"expired_at"`
}

type LoginUserResponse struct {
	Username  string    `json:"username"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expired_at"`
}
