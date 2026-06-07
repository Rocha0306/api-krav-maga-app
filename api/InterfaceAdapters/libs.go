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

func GerarUUID() string {
	return uuid.NewString()
}

func GerarId() string {
	return strconv.Itoa(int(uuid.New().ID()))
}

var chave_jwt = []byte("lorenzo_teste")

func GerarTokenJwt(id_usuario string) string {
	claims_token := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
		Audience:  jwt.ClaimStrings{id_usuario},
	}

	config_token_jwt := jwt.NewWithClaims(jwt.SigningMethodHS256, claims_token)

	string_jwt, err := config_token_jwt.SignedString(chave_jwt)

	if err == nil {
		return string_jwt
	} else {
		return err.Error()
	}
}

func ValidarTokenJwt(stringToken string) (string, error) {
	claims := &jwt.RegisteredClaims{}

	token_parseado, err := jwt.ParseWithClaims(stringToken, claims, func(token *jwt.Token) (interface{}, error) {
		return chave_jwt, nil
	})

	if err != nil || !token_parseado.Valid {
		return "", errors.New("jwt invalido")
	}

	if len(claims.Audience) == 0 {
		return "", errors.New("jwt invalido")
	}

	return claims.Audience[0], nil
}

func HashSenha(senha string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(senha), 10)
	return string(hash)
}

func GerarNumeroAuth() int {
	return rand.IntN(1000000)
}

func EnviarEmail(conteudo string, para []string) error {
	email_de := "l.m.p.rocha2005@gmail.com"
	servidor_smtp := mail.NewSMTPClient()
	servidor_smtp.Host = "smtp.gmail.com"
	servidor_smtp.Port = 587
	servidor_smtp.Encryption = mail.EncryptionSTARTTLS
	servidor_smtp.Username = email_de
	servidor_smtp.Password = "zaug agoi wmtw xceg"

	cliente_smtp, err := servidor_smtp.Connect()
	if err != nil {
		return err
	}

	msg := mail.NewMSG()
	msg.SetFrom(email_de)
	msg.AddTo(para...)
	msg.SetSubject("Autenticao aplicativo Krav")
	msg.SetBody(mail.TextPlain, conteudo)

	return msg.Send(cliente_smtp)
}

func DateTimeNow() time.Time {
	date_time_now, _ := time.Parse("2026-06-07", time.Now().String())
	return date_time_now
}
