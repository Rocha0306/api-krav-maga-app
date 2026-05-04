package entities

import "time"

type Address struct {
	ID          string `gorm:"primaryKey;column:ID;type:varchar(100)"`
	CEP         string `gorm:"column:cep;uniqueIndex;not null"`
	Logradouro  string `gorm:"column:logradouro"`
	Complemento string `gorm:"column:complemento"`
	Unidade     string `gorm:"column:unidade"`
	Bairro      string `gorm:"column:bairro"`
	Localidade  string `gorm:"column:localidade"`
	UF          string `gorm:"column:uf"`
	Estado      string `gorm:"column:estado"`
	Regiao      string `gorm:"column:regiao"`
	DDD         string `gorm:"column:ddd"`
}

type Users struct {
	ID          string    `gorm:"primaryKey;column:ID;type:varchar(100)"`
	Name        string    `gorm:"column:Name;not null"`
	Email       string    `gorm:"column:Email;uniqueIndex;not null"`
	Password    string    `gorm:"column:Password;not null"`
	Gender      string    `gorm:"column:Gender"`
	CPF         string    `gorm:"column:CPF;uniqueIndex;not null"`
	DateBirth   time.Time `gorm:"column:DateBirth"`
	AddressesID string    `gorm:"column:addresses_id"`
	Address     Address   `gorm:"foreignKey:AddressesID;references:ID"`
}

type Gyms struct {
	ID   string `gorm:"primaryKey;column:ID;type:varchar(100)"`
	Name string `gorm:"column:NAME;uniqueIndex;not null"`
	CNPJ string `gorm:"column:CNPJ;uniqueIndex;not null"`
}

type Teachers struct {
	IDTeacher       string `gorm:"primaryKey;column:id_teacher;type:varchar(100)"`
	IDGymsTeachers  string `gorm:"column:id_gyms_teachers"`
	IDUsersTeachers string `gorm:"column:id_users_teachers"`
	Gym             Gyms   `gorm:"foreignKey:IDGymsTeachers;references:ID"`
	User            Users  `gorm:"foreignKey:IDUsersTeachers;references:ID"`
}

type Invites struct {
	IDInvite  string `gorm:"primaryKey;column:id_invite;type:varchar(100)"`
	IDGym     string `gorm:"column:id_gym"`
	InviteKey string `gorm:"column:invite_key;uniqueIndex;not null"`
	Gym       Gyms   `gorm:"foreignKey:IDGym;references:ID"`
}

type InviteRequests struct {
	IDRequest string `gorm:"primaryKey;column:id_request;type:varchar(100)"`
	IDUser    string `gorm:"column:id_user"`
	IDGym     string `gorm:"column:id_gym"`
	User      Users  `gorm:"foreignKey:IDUser;references:ID"`
	Gym       Gyms   `gorm:"foreignKey:IDGym;references:ID"`
}

// Tabela associativa | muitos para muitos
type Gym_Student struct {
	ID      int
	Gym     int
	Student int
}
