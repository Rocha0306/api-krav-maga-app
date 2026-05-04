package InterfaceAdapters

import (
	"errors"
	"math/rand/v2"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	mail "github.com/xhit/go-simple-mail/v2"
	"golang.org/x/crypto/bcrypt"
)

func GenerateUUID() string {
	return uuid.NewString()
}

func GenerateId() string {
	return strconv.Itoa(int(uuid.New().ID()))
}

var jwt_key = []byte("lorenzo_teste")

func GenerateTokenJwt(id_user string) string {
	token_claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
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

func ValidateTokenJwt(tokenString string) (string, error) {
	claims := &jwt.RegisteredClaims{}

	jwt_parseado, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwt_key, nil
	})

	if err != nil || !jwt_parseado.Valid {
		return "", errors.New("jwt invalido")
	}

	if len(claims.Audience) == 0 {
		return "", errors.New("jwt invalido")
	}

	return claims.Audience[0], nil
}

func HashPassword(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(hash)
}

func GenerateAuthNumber() int {
	return rand.IntN(1000000)
}

func SendEmail(content string, who []string) error {
	email_from := "l.m.p.rocha2005@gmail.com"
	smtp_server := mail.NewSMTPClient()
	smtp_server.Host = "smtp.gmail.com"
	smtp_server.Port = 587
	smtp_server.Encryption = mail.EncryptionSTARTTLS
	smtp_server.Username = email_from
	smtp_server.Password = "zaug agoi wmtw xceg"

	smtp_client, err := smtp_server.Connect()
	if err != nil {
		return err
	}

	msg := mail.NewMSG()
	msg.SetFrom(email_from)
	msg.AddTo(who...)
	msg.SetSubject("Autenticao aplicativo Krav")
	msg.SetBody(mail.TextPlain, content)

	return msg.Send(smtp_client)
}
