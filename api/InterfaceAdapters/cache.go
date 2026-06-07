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
		Addr:     "redis-15478.crce296.us-east-1-6.ec2.cloud.redislabs.com:15478",
		Username: "default",
		Password: "QVnLZwuKjLiVDme0NtVjMnfP3cVdgqny",
		DB:       0,
	})
}

func SalvarUsuarioCache(codigo_auth string, usuario entities.Usuarios) error {
	bytes, _ := json.Marshal(usuario)
	return abrirRedis().Set(context.Background(), codigo_auth, bytes, 15*time.Minute).Err()
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
