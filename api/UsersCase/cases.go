package UsersCase

import (
	entities "api-back-end/api/Entities"
	"api-back-end/api/InterfaceAdapters"
	"api-back-end/api/Repository"
	"errors"
	"fmt"
	"strconv"
)

/*
1 - Pega o Response da API de CPF e CEP
2 - Popula o struct que é a entidade no banco
3 - Insere no banco o Cep e coloca o User em cache
*/
func PrepareUser(cpf string, cep string, email string, password string) error {

	if Repository.Exists[entities.Users]("Email", email) {
		return errors.New("Usuario ja cadastrado")
	}

	err, response_api_cpf := InterfaceAdapters.CpfApi(cpf)
	if err != nil {
		return err
	}

	err, response_api_cep := InterfaceAdapters.CepApi(cep)
	if err != nil {
		return err
	}

	final_entity := InterfaceAdapters.MapStudent(email, password, response_api_cpf, response_api_cep)
	Repository.Insert(&final_entity.Address)
	auth_number := strconv.Itoa(InterfaceAdapters.GenerateAuthNumber())
	InterfaceAdapters.PutUserCache(auth_number, *final_entity)
	InterfaceAdapters.SendEmail(fmt.Sprintf("Codigo de Autenticacao: %s", auth_number), []string{final_entity.Email})
	return nil
}

/*
1. Pega o User do cache
2. Insere no banco
*/
func CreateUser(key_code string) error {
	user_database, err := InterfaceAdapters.GetUserCache(key_code)
	if err != nil {
		return errors.New("codigo de autenticacao invalido ou expirado")
	}

	Repository.Insert(&user_database)
	InterfaceAdapters.RemoveValueFromCache(key_code)
	return nil
}

func Login(email string, password string) (string, error) {
	user := Repository.SelectWhere[entities.Users]("Email", email)
	password_hashed := InterfaceAdapters.HashPassword(password)

	if user.Email != email && user.Password != password_hashed {
		return "", errors.New("Aluno nao existe no sistema")
	}

	return InterfaceAdapters.GenerateTokenJwt(user.ID), nil
}

func CreateGym(cnpj string, name_gym string, id_user string) error {
	gym_entity := InterfaceAdapters.MapGym(cnpj, name_gym)
	Repository.Insert(gym_entity)
	teacher_entity := InterfaceAdapters.MapTeacher(id_user, gym_entity.ID)
	Repository.Insert(&teacher_entity)
	return nil
}

func GenerateInvite(id_user string) error {
	if !Repository.Exists[entities.Teachers]("id_users_teachers", id_user) {
		return errors.New("Voce nao pode criar convites ja que nao eh professor")
	}

	teacher := Repository.SelectWhere[entities.Teachers]("id_users_teachers", id_user)
	invite_entity := InterfaceAdapters.MapInvite(teacher.IDGymsTeachers)
	Repository.Insert(invite_entity)
	return nil
}

func ShowInvites(id_user string) (entities.Invites, error) {
	if !Repository.Exists[entities.Teachers]("id_users_teachers", id_user) {
		return entities.Invites{}, errors.New("Voce nao pode visualizar invites ja que nao eh um professor")
	}

	teacher := Repository.SelectWhere[entities.Teachers]("id_users_teachers", id_user)
	invite := Repository.SelectWhere[entities.Invites]("id_gym", teacher.IDGymsTeachers)
	return *invite, nil
}

func RequestToJoin(invite_code string, id_user string) error {
	if !Repository.Exists[entities.Invites]("invite_key", invite_code) {
		return errors.New("invite nao encontrado")
	}

	invite := Repository.SelectWhere[entities.Invites]("invite_key", invite_code)
	request_join := InterfaceAdapters.MapJoinInvite(invite.IDGym, id_user)
	Repository.Insert(request_join)
	return nil
}

type UserRequestInfo struct {
	Name string
	CPF  string
	CEP  string
}

func ListRequestsToJoin(id_user string) ([]UserRequestInfo, error) {
	if !Repository.Exists[entities.Teachers]("id_users_teachers", id_user) {
		return nil, errors.New("Voce nao pode visualizar convites ja que nao eh um professor")
	}

	teacher := Repository.SelectWhere[entities.Teachers]("id_users_teachers", id_user)
	list_requests := Repository.SelectWhereList[entities.InviteRequests]("id_gym", teacher.IDGymsTeachers)

	result := []UserRequestInfo{}
	for i := 0; i < len(list_requests); i++ {
		user := Repository.SelectWhere[entities.Users]("ID", list_requests[i].IDUser)
		address := Repository.SelectWhere[entities.Address]("ID", user.AddressesID)
		result = append(result, UserRequestInfo{
			Name: user.Name,
			CPF:  user.CPF,
			CEP:  address.CEP,
		})
	}

	return result, nil
}
