package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nathakusuma/bcc-be-freepass-2024/entity"
	"github.com/nathakusuma/bcc-be-freepass-2024/model"
	"log"
	"os"
	"time"
)

func GenerateToken(payload entity.User) (string, error) {
	expString := os.Getenv("JWT_TTL")
	exp, err := time.ParseDuration(expString)
	if expString == "" || err != nil {
		exp = time.Hour
		log.Println("[WARNING] JWT_TTL is not set!")
	}

	unsignedJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, model.NewUserClaims(payload.ID, exp))
	signedJWT, err := unsignedJWT.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return signedJWT, nil
}

func DecodeToken(signedJwt string, claims jwt.Claims, KEY string) error {
	token, err := jwt.ParseWithClaims(signedJwt, claims, func(t *jwt.Token) (any, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return "", errors.New("wrong signing method")
		}
		return []byte(KEY), nil
	})

	if err != nil {
		return errors.New("error when parsing token")
	}

	if !token.Valid {
		return errors.New("invalid token")
	}

	return nil
}
