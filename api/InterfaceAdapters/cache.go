package InterfaceAdapters

import (
	entities "api-back-end/api/Entities"
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

func abrirRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "enchanting-place-reaction-20529.db.redis.io:12800",
		Username: "default",
		Password: "dSzsiTMprSKooLGEN8LDpbitcalCdf9y",
		DB:       0,
	})
}

func SalvarUsuarioCache(codigo_auth string, usuario entities.Usuarios) error {
	bytes, _ := json.Marshal(usuario)
	return abrirRedis().Set(context.Background(), codigo_auth, bytes, 2*time.Minute).Err()
}

func PegarUsuarioCache(chave string) (entities.Usuarios, error) {
	usuario := entities.Usuarios{}

	val, err := abrirRedis().Get(context.Background(), chave).Bytes()
	if err == redis.Nil {
		return usuario, errors.New("conteudo nao encontrado em cache")
	}

	err = json.Unmarshal(val, &usuario)
	return usuario, err
}

func RemoverValorCache(chave string) {
	abrirRedis().Del(context.Background(), chave)
}
