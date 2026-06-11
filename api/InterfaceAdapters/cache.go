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
		Addr:     ConnectionStringRedis(),
		Username: UserRedis(),
		Password: SenhaRedis(),
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

func SalvarValorCache(chave string, valor string) error {
	return abrirRedis().Set(context.Background(), chave, valor, 10*time.Minute).Err()
}

func PegarValorCache(chave string) (string, error) {
	val, err := abrirRedis().Get(context.Background(), chave).Result()
	if err == redis.Nil {
		return "", errors.New("conteudo nao encontrado em cache")
	}
	return val, err
}

func ThrottleAtingido(chave string, janela time.Duration) bool {
	ok, _ := abrirRedis().SetNX(context.Background(), chave, "1", janela).Result()
	return !ok
}
