package utils

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string `json:"username"`
	UseID    int    `json:"userID"`
	IsAdmin  bool   `json:"is_admin"`
	jwt.StandardClaims
}

func GenerateToken(username string, userID int, isAdmin bool)(string,error) {
	expirationTime:=time.Now().Add(10 *time.Hour)
	claims:= &Claims{
		Username: username,
		UseID: userID,
		IsAdmin: isAdmin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}