package Presentation

import (
	"github.com/golang-jwt/jwt/v5"
)

const jwt_key string = ""

func GenerateTokenJwt() string {
	jwt_token_config := jwt.New(jwt.SigningMethodES256)
	jwt_string, err := jwt_token_config.SignedString(jwt_key)

	if err == nil {
		return jwt_string
	} else {
		return err.Error()
	}
}
