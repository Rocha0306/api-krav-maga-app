package InterfaceAdapters

import (
	"bytes"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func GenerateId() int {
	return int(uuid.New().ID())
}

const jwt_key string = "lorenzo_teste"

func GenerateTokenJwt(role string) string {
	token_claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(<-time.After(time.Duration(time.Duration(10).Minutes()))),
		Audience:  jwt.ClaimStrings{role},
	}

	bytes_key := bytes.NewBufferString(jwt_key).Bytes()

	jwt_token_config := jwt.NewWithClaims(jwt.SigningMethodHS256, token_claims)

	jwt_string, err := jwt_token_config.SignedString(bytes_key)

	if err == nil {
		return jwt_string
	} else {
		return err.Error()
	}
}
