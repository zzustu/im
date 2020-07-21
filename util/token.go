package util

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var secret = []byte{0xFF, 0x12}

func NewToken(username string) (string, error) {

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Audience:  username,
		ExpiresAt: 0,
		Id:        "",
		IssuedAt:  time.Now().Unix(),
		Issuer:    "",
		NotBefore: 0,
		Subject:   "",
	})

	return claims.SignedString(secret)
}

func ParseToken(token string) (*jwt.StandardClaims, error) {
	tk, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	if err = tk.Claims.Valid(); err != nil {
		return nil, err
	}
	if sc, ok := tk.Claims.(*jwt.StandardClaims); ok {
		return sc, nil
	} else {
		return nil, errors.New("不是标准的JWT")
	}
}
