package main

import (
	"api-back-end/api/Presentation"
	"api-back-end/api/UsersCase"
	"net/http"
)

func ControllerLogin(w http.ResponseWriter, r *http.Request) {
	Presentation.ValidationBody(w, r)
}

func StatusCode200(response http.ResponseWriter, message string) {
	response.WriteHeader(200)
	Presentation.SerializeMessageResponse(message, response)
}

func BadRequest(response http.ResponseWriter, err error) {
	response.WriteHeader(400)
	Presentation.SerializeErrorMessageResponse(err.Error(), response)

}

func ControllerRegistration(response_write http.ResponseWriter, request *http.Request) {

	if err := Presentation.ValidationBody(response_write, request); err != nil {
		BadRequest(response_write, err)
		return
	}

	student_dto := Presentation.DesserializeStudentDTO(request)

	if err := Presentation.ValidationFieldsStudent(student_dto); err != nil {
		BadRequest(response_write, err)
		return
	}

	if err := UsersCase.CreateStudent(student_dto.CPF, student_dto.CEP,
		student_dto.Username, student_dto.Password); err != nil {
		BadRequest(response_write, err)
		return
	}

	StatusCode200(response_write, "Aluno registrado com sucesso")

}

func main() {

	http.HandleFunc("/Students/Auth", ControllerLogin)
	http.HandleFunc("/Students/Registration", ControllerRegistration)
	http.ListenAndServe(":8080", nil)

}
