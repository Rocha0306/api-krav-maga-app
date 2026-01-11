package UsersCase

import (
	entities "api-back-end/api/Entities"
	"api-back-end/api/InterfaceAdapters"
	"api-back-end/api/Repository"
	"errors"
	"time"
)

func createentity(username string, password string, user_data InterfaceAdapters.CpfApiResponse, user_cep InterfaceAdapters.CepApiResponse) entities.Student {
	student := entities.Student{}

	student.Id = InterfaceAdapters.GenerateId()
	student.Name = user_data.Data.Name
	student.CPF = user_data.Data.CPF
	if user_data.Data.Gender == "M" {
		user_data.Data.Gender = "Masculino"
	} else {
		user_data.Data.Gender = "Feminino"
	}
	date, _ := time.Parse("02/01/2006", user_data.Data.BirthDate)
	student.DateBirth = date.Format("2006-01-02")
	student.Username = username
	student.Password = password
	student.Role = "Student"

	address := &entities.Address{}

	address.Id = InterfaceAdapters.GenerateId()
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

/*
1 - Pega o Response da API de CPF e CEP
2 - Popula o struct que é a entidade no banco
3 - Insere no banco o Cep primeiro e depois o Student (Relacionamento)
*/
func CreateStudent(cpf string, cep string, username string, password string) error {
	err, response_api_cpf := InterfaceAdapters.CpfApi(cpf)

	if err != nil {
		return err
	}

	err, response_api_cep := InterfaceAdapters.CepApi(cep)

	if err != nil {
		return err
	}

	final_entity := createentity(username, password, response_api_cpf, response_api_cep)
	Repository.InsertAddress(final_entity.Address, "Address", 11)
	Repository.InsertStudent(final_entity, "Student", 10)

	return nil

}

/*
1 - Campos do struct desejados especificados em int
2 - Realiza o select no banco
3 - Verifica se é exatamente igual
*/
func VerifyStudent(username string, password string) (entities.Student, error) {
	entity := entities.Student{}
	array := [3]int{2, 3, 8}
	student := Repository.SelectOneStudent(entity, array[:], username)
	result := entities.VerifyIsStudentExist(username, password, student)

	if result == false {
		return student, errors.New("Aluno nao existe no sistema")
	} else {
		return student, nil
	}
}
