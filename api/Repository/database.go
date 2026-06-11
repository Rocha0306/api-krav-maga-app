package Repository

import (
	entities "api-back-end/api/Entities"
	"api-back-end/api/InterfaceAdapters"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	dsn := InterfaceAdapters.ConnectionStringMySQL()
	conexao, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, _ := conexao.DB()
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)

	db = conexao
}

func connect() *gorm.DB {
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

func SelectPresencasComAula(id_aluno string) []entities.Presencas {
	var resultados []entities.Presencas
	connect().
		Joins("Aula").
		Where("presencas.id_aluno = ?", id_aluno).
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

func ExistsTwo[T any](campo1 string, valor1 string, campo2 string, valor2 string) bool {
	var modelo T
	var contagem int64
	connect().Model(&modelo).Where(campo1+" = ? AND "+campo2+" = ?", valor1, valor2).Count(&contagem)
	return contagem > 0
}

func UpdateLocationGym(id_academia string, latitude float64, longitude float64) {
	connect().Model(&entities.Academias{}).
		Where("ID = ?", id_academia).
		Updates(map[string]any{"latitude": latitude, "longitude": longitude})
}

func UpdateProduct(id_produto string, nome string, preco float64, tamanho string, quantidade int) {
	connect().Model(&entities.Produtos{}).
		Where("id_produto = ?", id_produto).
		Updates(map[string]any{"nome": nome, "preco": preco, "tamanho": tamanho, "quantidade": quantidade})
}

func UpdateAlunoFaixa(id_aluno string, faixa string) {
	connect().Model(&entities.Alunos{}).
		Where("id_aluno = ?", id_aluno).
		Updates(map[string]any{"faixa": faixa})
}

func UpdateSenhaUsuario(id_usuario string, senha_hash string) {
	connect().Model(&entities.Usuarios{}).
		Where("ID = ?", id_usuario).
		Updates(map[string]any{"senha": senha_hash})
}
