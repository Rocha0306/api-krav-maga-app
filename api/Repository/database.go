package Repository

import (
	entities "api-back-end/api/Entities"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const enderecoBanco string = "localhost"
const usuarioBanco string = "root"
const senhaBanco string = "lorenzo05*"
const nome_banco string = "KRAVMAGAAPP"

func connect() *gorm.DB {
	dsn := "root:lorenzo05*@tcp(localhost)/KRAVMAGAAPP?parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Print(err)
	}
	return db

}

func Inserir[T any](entidade *T) {
	connect().Create(entidade)
}

func SelectWhere[T any](campo string, valor string) *T {
	var resultado T
	connect().Where(campo+" = ?", valor).First(&resultado)
	return &resultado
}

func SelectWhereList[T any](campo string, valor string) []T {
	var resultados []T
	connect().Where(campo+" = ?", valor).Find(&resultados)
	return resultados
}

func SelectJoin[T any](join string, campo string, valor string) []T {
	var resultados []T
	connect().Joins(join).Where(campo+" = ?", valor).Find(&resultados)
	return resultados
}

func SelectInvitesByGym(id_academia string) []entities.SolicitacoesConvite {
	var resultados []entities.SolicitacoesConvite
	connect().
		Joins("Usuario").
		Joins("Usuario.Endereco").
		Where("solicitacoes_convite.id_academia = ?", id_academia).
		Find(&resultados)
	return resultados
}

// Uma unica query: alunos LEFT JOIN usuarios, ja preenchendo o Usuario de cada aluno.
func SelectStudentsByGym(id_academia string) []entities.Alunos {
	var resultados []entities.Alunos
	connect().
		Joins("Usuario").
		Where("alunos.id_academia_aluno = ?", id_academia).
		Find(&resultados)
	return resultados
}

// Aulas de uma academia dentro de um intervalo de tempo (ex: o dia inteiro).
func SelectDayPresences(id_academia string, inicio time.Time, fim time.Time) []entities.Aulas {
	var resultados []entities.Aulas
	connect().
		Where("id_academia = ? AND data_aula >= ? AND data_aula < ?", id_academia, inicio, fim).
		Find(&resultados)
	return resultados
}

func Delete[T any](campo string, valor string) {
	var modelo T
	connect().Where(campo+" = ?", valor).Delete(&modelo)
}

func Count[T any](campo string, valor string) int64 {
	var modelo T
	var contagem int64
	connect().Model(&modelo).Where(campo+" = ?", valor).Count(&contagem)
	return contagem
}

func Exists[T any](campo string, valor string) bool {
	var modelo T
	var contagem int64
	connect().Model(&modelo).Where(campo+" = ?", valor).Count(&contagem)
	return contagem > 0
}
