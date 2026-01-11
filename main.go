package main

import (
	"api-back-end/api/InterfaceAdapters"
	"api-back-end/api/Presentation"
	"api-back-end/api/UsersCase"
	"net/http"
)

func StatusCode200(response http.ResponseWriter, message string) {
	response.WriteHeader(200)
	Presentation.SerializeMessageResponse(message, response)
}

func BadRequest(response http.ResponseWriter, err error) {
	response.WriteHeader(400)
	Presentation.SerializeErrorMessageResponse(err.Error(), response)
}

func ControllerRegistration(response_write http.ResponseWriter, request *http.Request) {

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

func ControllerLogin(response http.ResponseWriter, request *http.Request) {
	student := Presentation.DesserializeLoginDTO(request)
	err := Presentation.ValidationFieldsLogin(student)
	if err != nil {
		BadRequest(response, err)
		return
	}

	entity, err := UsersCase.VerifyStudent(student.Username, student.Password)

	if err != nil {
		BadRequest(response, err)
		return
	}

	StatusCode200(response, InterfaceAdapters.GenerateTokenJwt(entity.Role))
}

func main() {

	http.HandleFunc("/Student/Auth", ControllerLogin)
	http.HandleFunc("/Student/Registration", ControllerRegistration)
	http.ListenAndServe(":8080", nil)

}
