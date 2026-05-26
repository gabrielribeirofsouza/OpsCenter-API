package config

import (
	"os"
	"time"
	dto "github.com/gabrielribeirofsouza/OpsCenter-API/DTO"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct{
	UserID  string `json:"user_id"`
	UserEmail string `json:"user_aemail"`
	Exp int64 `json:"exp"`
	jwt.RegisteredClaims
}

func GenerateToken(email string, id string) (dto.AuthResponse, error){
	SECRET_KEY := os.Getenv("SECRET_KEY")
	claims := Claims{
		UserID: id,
		UserEmail: email,
		Exp: time.Now().Add(time.Hour * 300).Unix(),
	}
	token:= jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil{
		return dto.AuthResponse{}, err
	}

	return dto.AuthResponse{
		AccessToken: signedToken,
		TokenType: "Bearer",
		ExpiresIn: time.Now().Add(time.Hour * 300).Unix(),
		Email: email,
		ID: id,
	}, nil
}