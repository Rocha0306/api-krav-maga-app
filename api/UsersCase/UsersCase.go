package UsersCase

import (
	"api-back-end/api/Entities"
	"api-back-end/api/InterfaceAdapters"
)

func createentity(user_data InterfaceAdapters.CpfApiResponse, user_cep InterfaceAdapters.CepApiResponse) StudentEntity {
	student := &entities.StudentEntity{}

	student.Id = InterfaceAdapters.GenerateId()
	student.Name = user_data.Name
	student.CPF = user_data.CPF
	student.Gender = user_data.Gender

	address := &entities.Address{}

	address.CEP = user_cep.CEP
	address.Logradouro = user_cep.Logradouro
	address.Complemento = user_cep.Complemento
	address.Bairro = user_cep.Bairro
	address.Localidade = user_cep.Localidade
	address.UF = user_cep.UF
	address.Estado = user_cep.Estado
	address.Regiao = user_cep.Regiao
	address.DDD = user_cep.DDD

	return student

}

func CreateStudent(cpf string, cep string, username string, password string) error {
	err, response_api_cpf := InterfaceAdapters.CpfApi(cpf)

	if err != nil {
		return err
	}

	err, response_api_cep := InterfaceAdapters.CepApi(cep)

	if err != nil {
		return err
	}

	InterfaceAdapters.InsertStudent(createentity(response_api_cpf, response_api_cep))

	return nil

}
