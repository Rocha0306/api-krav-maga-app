package Presentation

import (
	"errors"
)

func LancarExcecao(mensagem_erro string) error {
	return errors.New(mensagem_erro)
}
