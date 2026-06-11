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

func SenhaAppGmail() string {
	return PegaVariavelAmbiente("EMAIL_APP_PASS")
}

func JwtSecret() string {
	return PegaVariavelAmbiente("JWT_SECRET")
}
