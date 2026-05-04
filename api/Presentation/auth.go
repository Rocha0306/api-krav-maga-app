package Presentation

import (
	"api-back-end/api/InterfaceAdapters"
	"errors"
	"net/http"
)

func GetAndValidateJwt(request http.Request) (string, error) {
	token := request.Header.Get("Authorization")

	if token == "" {
		return "", errors.New("Token JWT ausente")
	}

	id_user, err := InterfaceAdapters.ValidateTokenJwt(token)

	if err != nil {
		return "", errors.New("token jwt invalido")
	}

	return id_user, nil
}
