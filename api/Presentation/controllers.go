package Presentation

import (
	"api-back-end/api/UsersCase"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ControllerRegistration(response_write http.ResponseWriter, request *http.Request) {

	user_dto := Deserialize[UserDTO](request)

	err := validator.New(validator.WithRequiredStructEnabled()).Struct(user_dto)

	if err != nil {
		BadRequest(response_write, err)
		return
	}

	if err := UsersCase.PrepareUser(user_dto.CPF, user_dto.CEP,
		user_dto.Email, user_dto.Password); err != nil {
		BadRequest(response_write, err)
		return
	}

	StatusCode200(response_write, "Um codigo de autententicao foi enviado a seu Email")

}

func ControllerRegistrationConfirm(response http.ResponseWriter, request *http.Request) {
	code_auth_number := Deserialize[CodeUserDTO](request)

	err_ := validator.New(validator.WithRequiredStructEnabled()).Struct(code_auth_number)

	if err_ != nil {
		BadRequest(response, err_)
		return
	}

	err := UsersCase.CreateUser(code_auth_number.CodeAuthNumber)

	if err != nil {
		BadRequest(response, err)
		return
	}

	StatusCode200(response, "Usuario autenticado e registrado com sucesso")
}

func ControllerLogin(response http.ResponseWriter, request *http.Request) {
	user := Deserialize[LoginDTO](request)

	err_ := validator.New(validator.WithRequiredStructEnabled()).Struct(user)

	if err_ != nil {
		BadRequest(response, err_)
		return
	}

	token, err := UsersCase.Login(user.Email, user.Password)

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
		return
	}

	gym_dto := Deserialize[GymDTO](request)

	error_ := validator.New(validator.WithRequiredStructEnabled()).Struct(gym_dto)

	if error_ != nil {
		BadRequest(response, error_)
		return
	}

	errw := UsersCase.CreateGym(gym_dto.CNPJ, gym_dto.Name, id_user)

	if errw != nil {
		BadRequest(response, errw)
		return
	}

}

func ControllerGenerateInvites(response http.ResponseWriter, request *http.Request) {
	id_user, err := GetAndValidateJwt(*request)

	if err != nil {
		BadRequest(response, err)
		return
	}

	err = UsersCase.GenerateInvite(id_user)

	if err != nil {
		BadRequest(response, err)
		return
	}

	StatusCode200(response, "Convite criado")

}

func ControllerShowInvites(response http.ResponseWriter, request *http.Request) {
	id_user, err := GetAndValidateJwt(*request)

	if err != nil {
		BadRequest(response, err)
		return
	}

	entity, err_ := UsersCase.ShowInvites(id_user)

	if err_ != nil {
		BadRequest(response, err_)
		return
	}

	StatusCode200(response, entity)
}

func ControllerGymsRequestsJoin(response http.ResponseWriter, request *http.Request) {
	id_user, err := GetAndValidateJwt(*request)

	if err != nil {
		BadRequest(response, err)
		return
	}

	invite_dto := Deserialize[InviteCodeUserDTO](request)
	error_ := validator.New(validator.WithRequiredStructEnabled()).Struct(invite_dto)

	if error_ != nil {
		BadRequest(response, error_)
		return
	}

	err = UsersCase.RequestToJoin(invite_dto.InviteUUID, id_user)

	if err != nil {
		BadRequest(response, error_)
		return
	}

	StatusCode200(response, "Solicitacao feita pra entrar na academia, aguarde a aprovacao do professor")

}

func ControllerListRequestsJoin(response http.ResponseWriter, request *http.Request) {
	id_user, err := GetAndValidateJwt(*request)

	if err != nil {
		BadRequest(response, err)
		return
	}

	users_entities, err := UsersCase.ListRequestsToJoin(id_user)

	if err != nil {
		BadRequest(response, err)
		return
	}

	StatusCode200(response, users_entities)

}
