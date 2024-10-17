package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Payload struct {
	Id        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	UserId    int64     `json:"user_id"`
	Email     string    `json:"email"`
	ExpiredAt time.Time `json:"expired_at"`
	CreatedAt time.Time `json:"created_at"`
	jwt.RegisteredClaims
}

func NewPayload(username string, email string, userId int64, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return &Payload{
		Id:        tokenID,
		UserId:    userId,
		Username:  username,
		Email:     email,
		ExpiredAt: time.Now().Add(duration),
		CreatedAt: time.Now(),
	}, nil

}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return fmt.Errorf("Token is expired")
	}
	return nil
}
