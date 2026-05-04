package InterfaceAdapters

import (
	entities "api-back-end/api/Entities"
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

func openredis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "redis-15478.crce296.us-east-1-6.ec2.cloud.redislabs.com:15478",
		Username: "default",
		Password: "QVnLZwuKjLiVDme0NtVjMnfP3cVdgqny",
		DB:       0,
	})
}

func PutUserCache(code_auth string, user entities.Users) error {
	bytes, _ := json.Marshal(user)
	return openredis().Set(context.Background(), code_auth, bytes, 15*time.Minute).Err()
}

func GetUserCache(key string) (entities.Users, error) {
	user := entities.Users{}

	val, err := openredis().Get(context.Background(), key).Bytes()
	if err == redis.Nil {
		return user, errors.New("conteudo nao encontrado em cache")
	}

	err = json.Unmarshal(val, &user)
	return user, err
}

func RemoveValueFromCache(key string) {
	openredis().Del(context.Background(), key)
}
