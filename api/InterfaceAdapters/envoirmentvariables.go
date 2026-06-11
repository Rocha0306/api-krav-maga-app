package InterfaceAdapters

import (
	"log"
	"os"
)

func PegaVariavelAmbiente(chave string) string {
	valor, existe := os.LookupEnv(chave)
	if !existe || valor == "" {
		log.Fatalf("Erro Crítico: A variável de ambiente '%s' não foi definida no Sistema Operacional!", chave)
	}
	return valor
}

func ConnectionStringBanco() string {
	return PegaVariavelAmbiente("DB_PASS")
}

func ConnectionStringMongo() string {
	return PegaVariavelAmbiente("DB_URL")
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

func SenhaAppGmail() string {
	return PegaVariavelAmbiente("EMAIL_APP_PASS")
}

func JwtSecret() string {
	return PegaVariavelAmbiente("JWT_SECRET")
}
