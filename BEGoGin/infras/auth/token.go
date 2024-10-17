package auth

import (
	"MLMW/BEGoGin/utils"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenMaker struct {
	secretKey string
}

func NewTokenMaker(secretKey string) (TokenMaker, error) {
	if len(secretKey) < utils.MIN_LENGTH_SECRET_KEY {
		return TokenMaker{}, fmt.Errorf("Invalid secret key size: secret key must at least %d characters", utils.MIN_LENGTH_SECRET_KEY)
	}
	return TokenMaker{
		secretKey: secretKey,
	}, nil
}

func (tokenMaker TokenMaker) NewToken(username string, email string, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(username, email, duration)
	if err != nil {
		return "", &Payload{}, err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, *payload)
	token, err := jwtToken.SignedString([]byte(tokenMaker.secretKey))
	if err != nil {
		return "", &Payload{}, err
	}

	return token, payload, err
}

func (tokenMaker TokenMaker) ValidToken(token string) (*Payload, error) {
	var keyFunc = func(tk *jwt.Token) (any, error) {
		//Token was not the same picked security
		if _, ok := tk.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("token is invalid")
		}
		return []byte(tokenMaker.secretKey), nil
	}
	//parse
	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		return nil, fmt.Errorf("token is invalid")
	}
	// get payload
	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, fmt.Errorf("token is invalid")
	}
	return payload, nil
}
