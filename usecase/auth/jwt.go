package auth

import (
	"errors"
	"log"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

type JwtSecret string

func (jwt JwtSecret) asByte() []byte {
	return []byte(jwt)
}

type Auth struct {
	JwtSecret
}

func (auth *Auth) getJWTString(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"iat":      time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	return token.SignedString(auth.JwtSecret.asByte())
}

var ErrorTokenStringSigningMethodMismatch = errors.New("token string signing method mismatch")
var ErrorInvalidToken = errors.New("ERROR: [engine::core::auth] invalid token")
var ErrorParsingClaims = errors.New("ERROR: [engine::core::auth] error parsing claims")

func (auth *Auth) validateTokenString(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrorTokenStringSigningMethodMismatch
		}
		return auth.JwtSecret.asByte(), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, ErrorInvalidToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok {
		log.Printf("[engine::core::auth] claims %+v", claims)
	} else {
		log.Printf("ERROR [engine::core::auth] claims have invalid assertion types.")
	}
	return token, nil
}
