package main

import (
	"fmt"
	"net/http"
)

type LoginDTO struct {
	Username string
	Password string
}

type StudentDTO struct {
	id        string
	Name      string
	Username  string
	RG        string
	CPF       string
	DateBirth string
	CEP       string
	Belt      string
}

func ControllerLogin(w http.ResponseWriter, r *http.Request) {
	ValidationBody(w, r)

}

func ControllerRegistration(response_write http.ResponseWriter, request *http.Request) {
	url := fmt.Sprintf("https://api.cpfhub.io/cpf/53201090832")
	/*
		ValidationBody(response_write, request)
		student_dto := DesserializeStudentDTO(request)
		ValidationFieldsStudent(student_dto, response_write)
	*/

	teste, err := http.Get(url)
	if err != nil {
		fmt.Print(teste)
	}

}

func main() {

	http.HandleFunc("/Students/Auth", ControllerLogin)
	http.HandleFunc("/Students/Registration", ControllerRegistration)
	http.ListenAndServe(":8080", nil)

}
