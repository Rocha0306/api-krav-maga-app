package Presentation

type LoginDTO struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserDTO struct {
	Email    string `json:"email" validate:"required,max=100"`
	Password string `json:"password" validate:"required,max=20"`
	CPF      string `json:"cpf"      validate:"required,len=11"`
	CEP      string `json:"cep"      validate:"required,len=8"`
}

type CodeUserDTO struct {
	CodeAuthNumber string `json:"CodeAuthNumber" validate:"required,max=6"`
}

type InviteCodeUserDTO struct {
	InviteUUID string `json:"InviteUUID" validate:"required,max=100"`
}

type GymDTO struct {
	CNPJ string `json:"cnpj" validate:"required,len=14"`
	Name string `json:"Name" validate:"required,max=100"`
}

type ResponseControllerDTO struct {
	Message string `json:"message"`
}

type ResponseErrorDTO struct {
	Message string `json:"message"`
}
