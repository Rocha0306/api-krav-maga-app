package UsersCase

import (
	entities "api-back-end/api/Entities"
	"api-back-end/api/InterfaceAdapters"
	"api-back-end/api/Repository"
	"errors"
	"strconv"
)

/*
1 - Pega o Response da API de CPF e CEP
2 - Popula o struct que é a entidade no banco
3 - Insere no banco o Cep primeiro e depois o Student (Relacionamento)
*/
func CreateStudent(cpf string, cep string, username string, password string) error {

	entity := entities.User{}
	user_database := Repository.SelectWhere(entity, []int{1}, cpf).(entities.User)

	if user_database.CPF != "" {
		return errors.New("Aluno ja existe na base de dados")
	}

	err, response_api_cpf := InterfaceAdapters.CpfApi(cpf)

	if err != nil {
		return err
	}

	err, response_api_cep := InterfaceAdapters.CepApi(cep)

	if err != nil {
		return err
	}

	final_entity := InterfaceAdapters.MapStudent(username, password, response_api_cpf, response_api_cep)
	Repository.Insert(final_entity.Address, "Address", 11)
	Repository.Insert(final_entity, "Student", 10)

	return nil

}

/*
1 - Campos do struct desejados especificados em int
2 - Realiza o select no banco
3 - Verifica se é exatamente igual
*/
func Login(username string, password string) (string, error) {
	entity := &entities.User{}
	student := Repository.SelectWhere(entity, []int{2, 3, 8}, username).(entities.User)

	if student.Username == username && student.Password == password {
		return "", errors.New("Aluno nao existe no sistema")
	}

	return InterfaceAdapters.GenerateTokenJwt(strconv.Itoa(student.Id)), nil
}

/*

1. Busca o Socio com base no cnpj
2. Vai no banco pega o nome do usuario
3. Cruza informacoes

*/

func CreateGym(cnpj string, name_gym string, id_user int) {

}
