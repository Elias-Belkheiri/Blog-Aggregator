package utils

import (
	// "os"
	// "time"

	"os"

	"github.com/golang-jwt/jwt"
)



func CreateToken(username string) (string, error) {
	var secretKey = []byte(os.Getenv("JWT_SECRET"))
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, 
        jwt.MapClaims{ 
        "username": username, 
        })

    return token.SignedString(secretKey)
}