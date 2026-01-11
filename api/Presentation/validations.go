package Presentation

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func ValidationBody(response_write http.ResponseWriter, request *http.Request) error {
	bodyBytes, _ := io.ReadAll(request.Body)
	if string(bodyBytes) == "{\n\t\n}" {
		return ThrowException("Sem payload")
	}

	return nil
}

func ValidationLenght(content string, maxsize int, name string, operator string) error {

	switch operator {
	case "Upper":
		if strings.Count(content, "") > maxsize {
			return ThrowException(fmt.Sprintf("O campo %s foi superior ao limite de %d caracteres", name, maxsize))
		}

	case "Different":
		if count := strings.Count(content, ""); count != maxsize {
			return ThrowException(fmt.Sprintf("O campo %s foi diferente do limite de %d caracteres", name, maxsize))
		}
	}

	return nil
}

func ValidationIsNullOrWhiteSpace(content string) error {
	if content == "" {
		return ThrowException("Nao existe em informacao em um dos campos enviados")
	}

	return nil
}

func ValidationFieldsLogin(login_dto LoginDTO) error {
	
	if(ValidationIsNullOrWhiteSpace(login_dto.Username) != nil) {
		return ThrowException("Campo Username nulo ou vazio")
	}

	if(ValidationIsNullOrWhiteSpace(login_dto.Password) != nil) {
		return ThrowException("Campo Password nulo ou vazio")
	}
	
	return nil
}

func ValidationFieldsStudent(studentdto UserDTO) error {
	var err error

	if ValidationIsNullOrWhiteSpace(studentdto.CPF) != nil ||
		ValidationLenght(studentdto.CPF, 12, "CPF", "Different") != nil {
		return errors.New("CPF nulo ou diferente de 11 caracteres, o cpf deve conter 11 caracteres")
	}

	if ValidationIsNullOrWhiteSpace(studentdto.Username) !=
		nil || ValidationLenght(studentdto.Username, 15, "Username", "Upper") != nil {
		return errors.New("Username nulo ou superior a 30 caracteres")
	}

	if ValidationIsNullOrWhiteSpace(studentdto.Password) != nil ||
		ValidationLenght(studentdto.Username, 20, "Password", "Upper") != nil {
		return errors.New("Password nulo ou superior a 20 caracteres")
	}

	if ValidationIsNullOrWhiteSpace(studentdto.Password) != nil ||
		ValidationLenght(studentdto.Username, 8, "CEP", "Different") != nil {
		return errors.New("CEP diferente de 8 caracteres")
	}

	return err

}
