package models

import "time"

type UserSchema struct {
	Id             string    `json:"id" db:"id"`
	Username       string    `json:"username" db:"username"`
	FullName       string    `json:"full_name" db:"full_name"`
	Email          string    `json:"email" db:"email"`
	HashedPassword string    `json:"hashed_password" db:"hashed_password"`
	IsDeleted      bool      `json:"is_deleted" db:"is_deleted"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}
