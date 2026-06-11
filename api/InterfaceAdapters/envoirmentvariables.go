package InterfaceAdapters

import (
	"fmt"
	"os"
)

func PegaVariavelAmbiente(chave string) string {
	valor, existe := os.LookupEnv(chave)
	if !existe || valor == "" {
		LogFatal(fmt.Sprintf("Erro Crítico: A variável de ambiente '%s' não foi definida no Sistema Operacional!", chave))
	}
	return valor
}

func ConnectionStringMySQL() string {
	return PegaVariavelAmbiente("DB_MYSQL")
}

func ConnectionStringMongo() string {
	return PegaVariavelAmbiente("DB_MONGODB")
}

func ConnectionStringRedis() string {
	return PegaVariavelAmbiente("REDIS_URL")
}

func UserRedis() string {
	return os.Getenv("REDIS_USER")
}

func SenhaRedis() string {
	return PegaVariavelAmbiente("REDIS_PASS")
}

// Chave da API do Brevo. Aceita BREVO_API_KEY; se nao existir, reaproveita
// a EMAIL_APP_PASS que ja estava configurada, pra nao exigir env nova.
func EmailApiKey() string {
	if chave := os.Getenv("BREVO_API_KEY"); chave != "" {
		return chave
	}
	return os.Getenv("EMAIL_APP_PASS")
}

// Email remetente verificado no Brevo. Pode ser ate um gmail (via "single
// sender verification"), sem precisar de dominio proprio. Setar em EMAIL_FROM.
func EmailRemetente() string {
	return os.Getenv("EMAIL_FROM")
}

func JwtSecret() string {
	return PegaVariavelAmbiente("JWT_SECRET")
}
