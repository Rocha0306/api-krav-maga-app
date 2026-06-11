package InterfaceAdapters

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"math/rand/v2"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
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
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(60 * time.Minute)),
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

// EnviarEmail envia via API HTTP do Brevo (porta 443), pois o Render bloqueia
// as portas de SMTP de saida (25/465/587) — por isso o SMTP direto dava timeout.
// O Brevo permite verificar so um email remetente (ate um gmail), sem exigir
// dominio proprio.
func EnviarEmail(conteudo string, para []string) error {
	destinatarios := make([]map[string]string, 0, len(para))
	for _, email := range para {
		destinatarios = append(destinatarios, map[string]string{"email": email})
	}

	corpo, _ := json.Marshal(map[string]interface{}{
		"sender":      map[string]string{"name": "KravConnect", "email": EmailRemetente()},
		"to":          destinatarios,
		"subject":     "Autenticacao aplicativo Krav",
		"textContent": conteudo,
	})

	requisicao, err := http.NewRequest("POST", "https://api.brevo.com/v3/smtp/email", bytes.NewBuffer(corpo))
	if err != nil {
		EscreverLogsMongoDb("Erro ao montar requisicao de email: "+err.Error(), "InterfaceAdapters/libs.go/EnviarEmail()")
		return err
	}
	requisicao.Header.Set("api-key", EmailApiKey())
	requisicao.Header.Set("Content-Type", "application/json")
	requisicao.Header.Set("Accept", "application/json")

	cliente := &http.Client{Timeout: 15 * time.Second}
	resposta, err := cliente.Do(requisicao)
	if err != nil {
		EscreverLogsMongoDb("Erro ao enviar email: "+err.Error(), "InterfaceAdapters/libs.go/EnviarEmail()")
		return err
	}
	defer resposta.Body.Close()

	if resposta.StatusCode >= 300 {
		detalhe, _ := io.ReadAll(resposta.Body)
		msg := fmt.Sprintf("Brevo retornou %d: %s", resposta.StatusCode, string(detalhe))
		EscreverLogsMongoDb(msg, "InterfaceAdapters/libs.go/EnviarEmail()")
		return errors.New(msg)
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
