package Presentation

import (
	"api-back-end/api/UsersCase"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ControllerRegistration(response_write http.ResponseWriter, request *http.Request) {

	user_dto := DesserializeUserDTO(request)

	err := validator.New(validator.WithRequiredStructEnabled()).Struct(user_dto)

	if err != nil {
		BadRequest(response_write, err)
		return
	}

	if err := UsersCase.CreateStudent(user_dto.CPF, user_dto.CEP,
		user_dto.Username, user_dto.Password); err != nil {
		BadRequest(response_write, err)
		return
	}

	StatusCode200(response_write, "Aluno registrado com sucesso")

}

func ControllerLogin(response http.ResponseWriter, request *http.Request) {
	user := DesserializeLoginDTO(request)

	err_ := validator.New(validator.WithRequiredStructEnabled()).Struct(user)

	if err_ != nil {
		BadRequest(response, err_)
		return
	}

	token, err := UsersCase.Login(user.Username, user.Password)

	if err != nil {
		BadRequest(response, err)
		return
	}

	StatusCode200(response, token)
}

func ControllerGymCreation(response http.ResponseWriter, request *http.Request) {
	id_user, err := GetAndValidateJwt(*request)

	if err != nil {
		BadRequest(response, err)
	}

	gym_dto := DesserializeGymDTO(request.Body)

	error_ := validator.New(validator.WithRequiredStructEnabled()).Struct(gym_dto)

	if error_ != nil {
		BadRequest(response, error_)
		return
	}

	UsersCase.CreateGym(gym_dto.CNPJ, gym_dto.Name, id_user)

}
