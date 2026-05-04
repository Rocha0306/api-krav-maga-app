package InterfaceAdapters

import (
	entities "api-back-end/api/Entities"
	"fmt"
	"time"
)

func MapStudent(email string, password string, user_data CpfApiResponse, user_cep CepApiResponse) *entities.Users {
	address := &entities.Address{}
	address.ID = GenerateId()
	address.CEP = user_cep.CEP
	address.Logradouro = user_cep.Logradouro
	address.Complemento = user_cep.Complemento
	address.Unidade = user_cep.Unidade
	address.Bairro = user_cep.Bairro
	address.Localidade = user_cep.Localidade
	address.UF = user_cep.UF
	address.Estado = user_cep.Estado
	address.Regiao = user_cep.Regiao
	address.DDD = user_cep.DDD

	user := &entities.Users{}
	user.ID = GenerateId()
	user.Name = user_data.Data.Name
	user.CPF = user_data.Data.CPF
	user.Gender = user_data.Data.Gender
	date, _ := time.Parse("02/01/2006", user_data.Data.BirthDate)
	user.DateBirth = date
	user.Email = email
	user.Password = HashPassword(password)
	user.AddressesID = address.ID
	user.Address = *address

	return user
}

func MapGym(cnpj string, name string) *entities.Gyms {
	gym := &entities.Gyms{
		ID:   GenerateId(),
		CNPJ: cnpj,
		Name: name,
	}

	return gym

}

func MapTeacher(id_user string, id_gym string) entities.Teachers {
	teacher := entities.Teachers{
		IDTeacher:       GenerateId(),
		IDGymsTeachers:  id_gym,
		IDUsersTeachers: id_user,
	}

	return teacher
}

func MapInvite(id_gym string) *entities.Invites {
	return &entities.Invites{
		IDInvite:  GenerateId(),
		IDGym:     id_gym,
		InviteKey: fmt.Sprintf("URLFrontComParametro?invite=%s", GenerateUUID()),
	}
}

func MapJoinInvite(id_gym string, id_user string) *entities.InviteRequests {
	return &entities.InviteRequests{
		IDRequest: GenerateId(),
		IDUser:    id_user,
		IDGym:     id_gym,
	}
}
