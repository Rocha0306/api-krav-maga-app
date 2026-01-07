package Presentation

type LoginDTO struct {
	Username string
	Password string
}

type UserDTO struct {
	Username string
	Password string
	CPF      string
	CEP      string
}

type ResponseControllerDTO struct {
	message string
}

type ResponseErrorDTO struct {
	message string
}
