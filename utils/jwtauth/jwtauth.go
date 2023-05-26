package jwtauth

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JwtCustomClaims struct {
	UserId   int
	Username string
	jwt.RegisteredClaims
}

func GenerateToken() {
	now := time.Now()
	end := now.Add(time.Duration(60) * time.Second)

	claims := JwtCustomClaims{
		Username: "xiaojia",
		UserId:   11011,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "test",
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(end),
		},
	}
	mySigningKey := []byte("AllYourBase")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	fmt.Printf("token: %v, err: %v", ss, err)
}

func ValidToken(tokenStr string) (*JwtCustomClaims, error) {

	token, err := jwt.ParseWithClaims(tokenStr, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	}, jwt.WithLeeway(10*time.Second))

	if claims, ok := token.Claims.(*JwtCustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
