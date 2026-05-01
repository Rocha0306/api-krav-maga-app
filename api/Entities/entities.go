package entities

type User struct {
	Id        int
	Name      string
	Username  string
	Password  string
	Gender    string
	CPF       string
	DateBirth string
	Address   Address
}

type Gym struct {
	GymName string
	cep     string
}

// Tabela associativa | muitos para muitos
type Gym_Student struct {
	id      int
	Gym     int
	Student int
}

type Address struct {
	Id          int
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
