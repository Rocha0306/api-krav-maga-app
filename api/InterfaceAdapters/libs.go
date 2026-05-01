package InterfaceAdapters

import (
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func GenerateId() int {
	return int(uuid.New().ID())
}

var jwt_key = []byte("lorenzo_teste")

func GenerateTokenJwt(id_user string) string {
	token_claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(<-time.After(time.Duration(time.Duration(10).Minutes()))),
		Audience:  jwt.ClaimStrings{id_user},
	}

	jwt_token_config := jwt.NewWithClaims(jwt.SigningMethodHS256, token_claims)

	jwt_string, err := jwt_token_config.SignedString(jwt_key)

	if err == nil {
		return jwt_string
	} else {
		return err.Error()
	}
}

func ValidateTokenJwt(tokenString string) (int, error) {
	jwt_parseado, err := jwt.ParseWithClaims(tokenString, nil, func(token *jwt.Token) (interface{}, error) {
		return jwt_key, nil
	})

	if err != nil {
		return 0, errors.New("jwt invalido")
	}

	claim, _ := jwt_parseado.Claims.GetAudience()

	id, _ := strconv.Atoi(claim[2])

	return id, nil
}

func HashPassword(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(hash)
}
