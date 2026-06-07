package Presentation

import (
	"api-back-end/api/InterfaceAdapters"
	"errors"
	"net/http"
)

func ValidarJwt(request http.Request) (string, error) {
	token := request.Header.Get("Authorization")

	if token == "" {
		return "", errors.New("Token JWT ausente")
	}

	id_usuario, err := InterfaceAdapters.ValidarTokenJwt(token)

	if err != nil {
		return "", errors.New("token jwt invalido")
	}

	return id_usuario, nil
}
