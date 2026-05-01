package InterfaceAdapters

import (
	entities "api-back-end/api/Entities"
	"time"
)

func MapStudent(username string, password string, user_data CpfApiResponse, user_cep CepApiResponse) entities.User {
	student := entities.User{}

	student.Id = GenerateId()
	student.Name = user_data.Data.Name
	student.CPF = user_data.Data.CPF
	student.Gender = user_data.Data.Gender
	date, _ := time.Parse("02/01/2006", user_data.Data.BirthDate)
	student.DateBirth = date.Format("2006-01-02")
	student.Username = username
	student.Password = HashPassword(student.Password)

	address := &entities.Address{}

	address.Id = GenerateId()
	address.CEP = user_cep.CEP
	address.Logradouro = user_cep.Logradouro
	address.Complemento = user_cep.Complemento
	address.Bairro = user_cep.Bairro
	address.Localidade = user_cep.Localidade
	address.UF = user_cep.UF
	address.Estado = user_cep.Estado
	address.Regiao = user_cep.Regiao
	address.DDD = user_cep.DDD

	student.Address = *address

	return student

}
