package Presentation

type LoginDTO struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserDTO struct {
	Username string `json:"username" validate:"required,max=15"`
	Password string `json:"password" validate:"required,max=20"`
	CPF      string `json:"cpf"      validate:"required,len=11"`
	CEP      string `json:"cep"      validate:"required,len=8"`
}

type GymDTO struct {
	CNPJ string `json:"cnpj" validate:"required,len=14"`
	Name string `json:"Name" validate:"required,len=30"`
}

type ResponseControllerDTO struct {
	Message string `json:"message"`
}

type ResponseErrorDTO struct {
	Message string `json:"message"`
}
