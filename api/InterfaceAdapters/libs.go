package InterfaceAdapters

import (
	"errors"
	"math"
	"math/rand/v2"
	"regexp"
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

var chave_jwt = []byte(JwtSecret())

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

func ComparaSenhaHash(hash string, senha string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(senha)) == nil
}

func GerarNumeroAuth() int {
	return rand.IntN(1000000)
}

func EnviarEmail(conteudo string, para []string) error {
	email_de := "kravconnect@gmail.com"
	servidor_smtp := mail.NewSMTPClient()
	servidor_smtp.Host = "smtp.gmail.com"
	servidor_smtp.Port = 587
	servidor_smtp.Encryption = mail.EncryptionSTARTTLS
	servidor_smtp.Username = email_de
	servidor_smtp.Password = SenhaAppGmail()

	cliente_smtp, err := servidor_smtp.Connect()
	if err != nil {
		EscreverLogsMongoDb("Erro ao conectar no SMTP do email: "+err.Error(), "InterfaceAdapters/libs.go/EnviarEmail()")
		return err
	}

	msg := mail.NewMSG()
	msg.SetFrom(email_de)
	msg.AddTo(para...)
	msg.SetSubject("Autenticao aplicativo Krav")
	msg.SetBody(mail.TextPlain, conteudo)
	err = msg.Send(cliente_smtp)

	if err != nil {
		EscreverLogsMongoDb("Erro ao enviar email: "+err.Error(), "InterfaceAdapters/libs.go/EnviarEmail()")
		return err
	}

	return nil
}

func DateTimeNow() time.Time {
	layout := "2006-01-02 15:04:05"
	date_time_now, _ := time.Parse(layout, time.Now().Format(layout))
	return date_time_now
}

var regexDominioEmail = regexp.MustCompile(`@(gmail|hotmail|outlook)\.com$`)

func RegexDominioEmail(email string) bool {
	return regexDominioEmail.MatchString(email)
}

func CalcularDistanciaMetros(lat1, lon1, lat2, lon2 float64) float64 {
	const raioTerra = 6371000.0
	lat1Rad := lat1 * math.Pi / 180
	lat2Rad := lat2 * math.Pi / 180
	deltaLat := (lat2 - lat1) * math.Pi / 180
	deltaLon := (lon2 - lon1) * math.Pi / 180
	a := math.Sin(deltaLat/2)*math.Sin(deltaLat/2) +
		math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Sin(deltaLon/2)*math.Sin(deltaLon/2)
	return raioTerra * 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
}
