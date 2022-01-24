package models

import (
	"api/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type TokenModel struct {
	//DB will be used when we add refresh
	graphDb       *GraphDb
	Secret        string
	ExpireMinutes int
}

func NewTokenModel(gdb *GraphDb, cfg config.Jwt) *TokenModel {
	return &TokenModel{
		graphDb:       gdb,
		Secret:        cfg.Secret,
		ExpireMinutes: cfg.ExpireMinutes}
}

func (j *TokenModel) GenerateJwtToken(user User) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["user_id"] = user.Id
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(j.ExpireMinutes))

	tokenString, _ := token.SignedString([]byte(j.Secret))

	return tokenString, nil
}
