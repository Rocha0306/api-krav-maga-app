package InterfaceAdapters

import (
	entities "api-back-end/api/Entities"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Migrations() {
	dsn := ConnectionStringMySQL()
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.AutoMigrate(
		&entities.Endereco{},
		&entities.Usuarios{},
		&entities.Academias{},
		&entities.Professores{},
		&entities.Convites{},
		&entities.SolicitacoesConvite{},
		&entities.Alunos{},
		&entities.Instrutores{},
		&entities.Aulas{},
		&entities.Presencas{},
		&entities.Pagamentos{},
		&entities.Produtos{},
		&entities.Interesses{},
	)
}
