package jwtauth

import (
	"RudderMaster/models/auth"
	"RudderMaster/settings"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JwtCustomClaims struct {
	UserId   uint
	Username string
	RoleId   uint
	jwt.RegisteredClaims
}

func NewClaims(user *auth.User) *JwtCustomClaims {
	now := time.Now()
	end := now.Add(time.Duration(settings.Config.Application.SecretExpire) * time.Second)
	return &JwtCustomClaims{
		Username: user.Username,
		UserId:   user.ID,
		RoleId:   user.RoleId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "test",
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(end),
		},
	}
}

func GenerateToken(user *auth.User) (tokenStr string, err error) {
	claims := NewClaims(user)
	mySigningKey := []byte(settings.Config.Application.Secret)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err = token.SignedString(mySigningKey)
	return
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
