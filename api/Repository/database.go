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

func conectarBanco() *gorm.DB {
	dsn := "root:lorenzo05*@tcp(localhost)/KRAVMAGAAPP?parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Print(err)
	}
	return db

}

func Inserir[T any](entidade *T) {
	conectarBanco().Create(entidade)
}

func BuscarOnde[T any](campo string, valor string) *T {
	var resultado T
	conectarBanco().Where(campo+" = ?", valor).First(&resultado)
	return &resultado
}

func BuscarListaOnde[T any](campo string, valor string) []T {
	var resultados []T
	conectarBanco().Where(campo+" = ?", valor).Find(&resultados)
	return resultados
}

func BuscarJoin[T any](join string, campo string, valor string) []T {
	var resultados []T
	conectarBanco().Joins(join).Where(campo+" = ?", valor).Find(&resultados)
	return resultados
}

func BuscarSolicitacoesPorAcademia(id_academia string) []entities.SolicitacoesConvite {
	var resultados []entities.SolicitacoesConvite
	conectarBanco().
		Joins("Usuario").
		Joins("Usuario.Endereco").
		Where("solicitacoes_convite.id_academia = ?", id_academia).
		Find(&resultados)
	return resultados
}

// Uma unica query: alunos LEFT JOIN usuarios, ja preenchendo o Usuario de cada aluno.
func BuscarAlunosPorAcademia(id_academia string) []entities.Alunos {
	var resultados []entities.Alunos
	conectarBanco().
		Joins("Usuario").
		Where("alunos.id_academia_aluno = ?", id_academia).
		Find(&resultados)
	return resultados
}

// Aulas de uma academia dentro de um intervalo de tempo (ex: o dia inteiro).
func BuscarAulasDoDia(id_academia string, inicio time.Time, fim time.Time) []entities.Aulas {
	var resultados []entities.Aulas
	conectarBanco().
		Where("id_academia = ? AND data_aula >= ? AND data_aula < ?", id_academia, inicio, fim).
		Find(&resultados)
	return resultados
}

func Deletar[T any](campo string, valor string) {
	var modelo T
	conectarBanco().Where(campo+" = ?", valor).Delete(&modelo)
}

func Contar[T any](campo string, valor string) int64 {
	var modelo T
	var contagem int64
	conectarBanco().Model(&modelo).Where(campo+" = ?", valor).Count(&contagem)
	return contagem
}

func Existe[T any](campo string, valor string) bool {
	var modelo T
	var contagem int64
	conectarBanco().Model(&modelo).Where(campo+" = ?", valor).Count(&contagem)
	return contagem > 0
}
