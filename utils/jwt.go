package utils

import (
	"os"
	"fmt"
	"github.com/golang-jwt/jwt"
)

func CreateToken(username string) string{
	var secret = (os.Getenv("JWT_SECRET"))
	claims := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"username": username,
	})

	token, err := claims.SignedString([]byte(secret))
	if err != nil {
		fmt.Println("Error creating token")
		panic(err)
	}

	return token
}