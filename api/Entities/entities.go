package entities

type StudentEntity struct {
	Id        uint32
	Name      string
	Username  string
	Gender    string
	CPF       string
	DateBirth string
	Belt      string
	Gym       Gym
	Adress    Address
}

type Gym struct {
	gym_name string
	cep      string
}

type Address struct {
	CEP         string
	Logradouro  string
	Complemento string
	Unidade     string
	Bairro      string
	Localidade  string
	UF          string
	Estado      string
	Regiao      string
	DDD         string
}
