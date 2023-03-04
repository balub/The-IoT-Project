package utils

import (
	"os"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateToken(user_id uint) (string, error) {

	// token_lifespan, err := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))
	token_lifespan, err := strconv.Atoi("3000")

	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("API_SECRET")))

}
