package Presentation

import (
	"api-back-end/api/InterfaceAdapters"
	"errors"
	"net/http"
)

func GetAndValidateJwt(request http.Request) (int, error) {
	token := request.Header.Get("Authorization")

	if token == "" {
		return 0, errors.New("Token JWT ausente")
	}

	id_user, err := InterfaceAdapters.ValidateTokenJwt(token)

	if err != nil {
		return 0, errors.New("token jwt invalido")
	}

	return id_user, nil
}
