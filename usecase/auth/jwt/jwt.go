package jwt

import (
	"errors"
	"log"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

type Secret string

type String string

func (jwt Secret) asByte() []byte {
	return []byte(jwt)
}

func (jwt String) String() string {
	return string(jwt)
}

type JWT struct {
	Secret
}

func (auth *JWT) GenerateJWTString(username string) (String, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"iat":      time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	token_String, err := token.SignedString(auth.Secret.asByte())

	if err != nil {
		return "", err
	}
	return String(token_String), nil
}

var ErrorTokenStringSigningMethodMismatch = errors.New("token string signing method mismatch")
var ErrorInvalidToken = errors.New("ERROR: [engine::core::auth] invalid token")
var ErrorParsingClaims = errors.New("ERROR: [engine::core::auth] error parsing claims")

func (auth *JWT) validateTokenString(tokenString String) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString.String(), func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrorTokenStringSigningMethodMismatch
		}
		return auth.Secret.asByte(), nil
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
