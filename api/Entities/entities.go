package entities

import "time"

type Endereco struct {
	ID          string `gorm:"primaryKey;column:ID;type:varchar(100)"`
	CEP         string `gorm:"column:cep;type:varchar(10);uniqueIndex;not null"`
	Logradouro  string `gorm:"column:logradouro;type:varchar(255)"`
	Complemento string `gorm:"column:complemento;type:varchar(255)"`
	Unidade     string `gorm:"column:unidade;type:varchar(50)"`
	Bairro      string `gorm:"column:bairro;type:varchar(100)"`
	Localidade  string `gorm:"column:localidade;type:varchar(100)"`
	UF          string `gorm:"column:uf;type:char(2)"`
	Estado      string `gorm:"column:estado;type:varchar(100)"`
	Regiao      string `gorm:"column:regiao;type:varchar(50)"`
	DDD         string `gorm:"column:ddd;type:varchar(5)"`
}

func (Endereco) TableName() string { return "enderecos" }

type Usuarios struct {
	ID             string    `gorm:"primaryKey;column:ID;type:varchar(100)"`
	Nome           string    `gorm:"column:nome;type:varchar(150);not null"`
	Email          string    `gorm:"column:email;type:varchar(200);uniqueIndex;not null"`
	Senha          string    `gorm:"column:senha;type:varchar(100);not null"`
	Genero         string    `gorm:"column:genero;type:char(10)"`
	CPF            string    `gorm:"column:CPF;type:char(11);uniqueIndex;not null"`
	DataNascimento time.Time `gorm:"column:data_nascimento"`
	EnderecoID     string    `gorm:"column:endereco_id"`
	Endereco       Endereco  `gorm:"foreignKey:EnderecoID;references:ID"`
}

func (Usuarios) TableName() string { return "usuarios" }

type Academias struct {
	ID        string   `gorm:"primaryKey;column:ID;type:varchar(100)"`
	Nome      string   `gorm:"column:nome;type:varchar(100);uniqueIndex;not null"`
	CNPJ      string   `gorm:"column:CNPJ;type:varchar(14);uniqueIndex;not null"`
	Latitude  *float64 `gorm:"column:latitude"`
	Longitude *float64 `gorm:"column:longitude"`
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
	ChaveConvite string    `gorm:"column:chave_convite;type:varchar(100);uniqueIndex;not null"`
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
	Faixa           string    `gorm:"column:faixa;type:varchar(15)"`
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
	Faixa       string    `gorm:"column:faixa;type:varchar(15)"`
	IDAcademia  string    `gorm:"column:id_academia"`
	IDInstrutor string    `gorm:"column:id_instrutor"`
	Academia    Academias `gorm:"foreignKey:IDAcademia;references:ID"`
	Instrutor   Usuarios  `gorm:"foreignKey:IDInstrutor;references:ID"`
}

func (Aulas) TableName() string { return "aulas" }

type Presencas struct {
	IDPresenca string    `gorm:"primaryKey;column:id_presenca;type:varchar(100)"`
	AlunoID    string    `gorm:"column:id_aluno;type:varchar(100)"`
	AulaID     string    `gorm:"column:id_aula;type:varchar(100)"`
	CheckinEm  time.Time `gorm:"column:checkin_em"`
	Aluno      Alunos    `gorm:"foreignKey:AlunoID;references:IDAluno"`
	Aula       Aulas     `gorm:"foreignKey:AulaID;references:IDAula"`
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

type Pagamentos struct {
	IDPagamento     string    `gorm:"primaryKey;column:id_pagamento;type:varchar(100)"`
	AlunoID         string    `gorm:"column:id_aluno;type:varchar(100)"`
	ValorCentavos   int64     `gorm:"column:valor_centavos"`
	Status          string    `gorm:"column:status;type:varchar(30)"`
	IDPaymentIntent string    `gorm:"column:id_payment_intent;type:varchar(100)"`
	CriadoEm        time.Time `gorm:"column:criado_em"`
	Aluno           Alunos    `gorm:"foreignKey:AlunoID;references:IDAluno"`
}

func (Pagamentos) TableName() string { return "pagamentos" }

type Produtos struct {
	IDProduto  string    `gorm:"primaryKey;column:id_produto;type:varchar(100)"`
	Nome       string    `gorm:"column:nome;type:varchar(150);not null"`
	Preco      float64   `gorm:"column:preco;not null"`
	Tamanho    string    `gorm:"column:tamanho;type:varchar(20)"`
	Quantidade int       `gorm:"column:quantidade;not null"`
	ImagemURL  string    `gorm:"column:imagem_url;type:varchar(500)"`
	IDAcademia string    `gorm:"column:id_academia"`
	Academia   Academias `gorm:"foreignKey:IDAcademia;references:ID"`
}

func (Produtos) TableName() string { return "produtos" }

type Interesses struct {
	IDInteresse string   `gorm:"primaryKey;column:id_interesse;type:varchar(100)"`
	AlunoID     string   `gorm:"column:id_aluno;type:varchar(100)"`
	ProdutoID   string   `gorm:"column:id_produto;type:varchar(100)"`
	Quantidade  int      `gorm:"column:quantidade"`
	Aluno       Alunos   `gorm:"foreignKey:AlunoID;references:IDAluno"`
	Produto     Produtos `gorm:"foreignKey:ProdutoID;references:IDProduto"`
}

func (Interesses) TableName() string { return "interesses" }

// Tabela associativa | muitos para muitos
type Academia_Aluno struct {
	ID       int
	Academia int
	Aluno    int
}
