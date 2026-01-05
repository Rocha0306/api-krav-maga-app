package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func ErrorValidation(content string, response_error http.ResponseWriter) {
	response_error.WriteHeader(400)
	json.NewEncoder(response_error).Encode("Sem payload")
}

func ValidationBody(response_write http.ResponseWriter, request *http.Request) {
	bodyBytes, err := io.ReadAll(request.Body)
	if err == nil && len(bodyBytes) <= 0 {
		ErrorValidation("Sem payload", response_write)
	}
}

func ValidationFormat(content string) {

}

func ValidationLenght(content string, maxsize int, name string, operator string, response http.ResponseWriter) {

	switch operator {
	case "Upper":
		if strings.Count(content, "") > maxsize {
			ErrorValidation(fmt.Sprintf("O campo %s foi superior ao limite de %d caracteres", name, maxsize), response)
		}

	case "Different":
		if strings.Count(content, "") != maxsize {
			ErrorValidation(fmt.Sprintf("O campo %s foi diferente do limite de %d caracteres", name, maxsize), response)
		}
	}

}

func ValidationFieldsStudent(studentdto StudentDTO, response http.ResponseWriter) {
	ValidationLenght(studentdto.Name, 20, "Name", "Upper", response)
	ValidationLenght(studentdto.CPF, 11, "CPF", "Different", response)
	ValidationLenght(studentdto.RG, 9, "RG", "Different", response)
	ValidationLenght(studentdto.Username, 9, "Username", "Different", response)
	ValidationLenght(studentdto.CEP, 8, "CEP", "Different", response)
	ValidationLenght(studentdto.DateBirth, 10, "Date Birth", "Different", response)

}
