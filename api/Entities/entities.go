package entities

import "time"

type Endereco struct {
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

func (Endereco) TableName() string { return "enderecos" }

type Usuarios struct {
	ID             string    `gorm:"primaryKey;column:ID;type:varchar(100)"`
	Nome           string    `gorm:"column:nome;not null"`
	Email          string    `gorm:"column:email;uniqueIndex;not null"`
	Senha          string    `gorm:"column:senha;not null"`
	Genero         string    `gorm:"column:genero"`
	CPF            string    `gorm:"column:CPF;uniqueIndex;not null"`
	DataNascimento time.Time `gorm:"column:data_nascimento"`
	EnderecoID     string    `gorm:"column:endereco_id"`
	Endereco       Endereco  `gorm:"foreignKey:EnderecoID;references:ID"`
}

func (Usuarios) TableName() string { return "usuarios" }

type Academias struct {
	ID   string `gorm:"primaryKey;column:ID;type:varchar(100)"`
	Nome string `gorm:"column:nome;uniqueIndex;not null"`
	CNPJ string `gorm:"column:CNPJ;uniqueIndex;not null"`
}

func (Academias) TableName() string { return "academias" }

type Professores struct {
	IDProfessor         string    `gorm:"primaryKey;column:id_professor;type:varchar(100)"`
	IDAcademiaProfessor string    `gorm:"column:id_academia_professor"`
	IDUsuarioProfessor  string    `gorm:"column:id_usuario_professor"`
	Academia            Academias `gorm:"foreignKey:IDAcademiaProfessor;references:ID"`
	Usuario             Usuarios  `gorm:"foreignKey:IDUsuarioProfessor;references:ID"`
}

func (Professores) TableName() string { return "professores" }

type Convites struct {
	IDConvite    string    `gorm:"primaryKey;column:id_convite;type:varchar(100)"`
	IDAcademia   string    `gorm:"column:id_academia"`
	ChaveConvite string    `gorm:"column:chave_convite;uniqueIndex;not null"`
	Academia     Academias `gorm:"foreignKey:IDAcademia;references:ID"`
}

func (Convites) TableName() string { return "convites" }

type SolicitacoesConvite struct {
	IDSolicitacao string    `gorm:"primaryKey;column:id_solicitacao;type:varchar(100)"`
	IDUsuario     string    `gorm:"column:id_usuario"`
	IDAcademia    string    `gorm:"column:id_academia"`
	Usuario       Usuarios  `gorm:"foreignKey:IDUsuario;references:ID"`
	Academia      Academias `gorm:"foreignKey:IDAcademia;references:ID"`
}

func (SolicitacoesConvite) TableName() string { return "solicitacoes_convite" }

type Alunos struct {
	IDAluno         string    `gorm:"primaryKey;column:id_aluno;type:varchar(100)"`
	Faixa           string    `gorm:"column:faixa"`
	IDUsuarioAluno  string    `gorm:"column:id_usuario_aluno"`
	IDAcademiaAluno string    `gorm:"column:id_academia_aluno"`
	Usuario         Usuarios  `gorm:"foreignKey:IDUsuarioAluno;references:ID"`
	Academia        Academias `gorm:"foreignKey:IDAcademiaAluno;references:ID"`
}

func (Alunos) TableName() string { return "alunos" }

type Aulas struct {
	IDAula      string    `gorm:"primaryKey;column:id_aula;type:varchar(100)"`
	DataAula    time.Time `gorm:"column:data_aula"`
	Conteudo    string    `gorm:"column:conteudo;type:varchar(255)"`
	IDAcademia  string    `gorm:"column:id_academia"`
	IDInstrutor string    `gorm:"column:id_instrutor"`
	Academia    Academias `gorm:"foreignKey:IDAcademia;references:ID"`
	Instrutor   Usuarios  `gorm:"foreignKey:IDInstrutor;references:ID"`
}

func (Aulas) TableName() string { return "aulas" }

type Presencas struct {
	IDPresenca string    `gorm:"primaryKey;column:id_presenca;type:varchar(100)"`
	IDAluno    string    `gorm:"column:id_aluno"`
	IDAula     string    `gorm:"column:id_aula"`
	CheckinEm  time.Time `gorm:"column:checkin_em"`
	Aluno      Alunos    `gorm:"foreignKey:IDAluno;references:IDAluno"`
	Aula       Aulas     `gorm:"foreignKey:IDAula;references:IDAula"`
}

func (Presencas) TableName() string { return "presencas" }

type Instrutores struct {
	IDInstrutor         string    `gorm:"primaryKey;column:id_instrutor;type:varchar(100)"`
	IDUsuarioInstrutor  string    `gorm:"column:id_usuario_instrutor"`
	IDAcademiaInstrutor string    `gorm:"column:id_academia_instrutor"`
	Usuario             Usuarios  `gorm:"foreignKey:IDUsuarioInstrutor;references:ID"`
	Academia            Academias `gorm:"foreignKey:IDAcademiaInstrutor;references:ID"`
}

func (Instrutores) TableName() string { return "instrutores" }

// Tabela associativa | muitos para muitos
type Academia_Aluno struct {
	ID       int
	Academia int
	Aluno    int
}
