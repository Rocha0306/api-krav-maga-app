package Presentation

import (
	"time"
)

type LoginDTO struct {
	Email string `json:"email" validate:"required"`
	Senha string `json:"senha" validate:"required"`
}

type UsuarioDTO struct {
	Email string `json:"email" validate:"required,max=100"`
	Senha string `json:"senha" validate:"required,max=20"`
	CPF   string `json:"cpf"   validate:"required,len=11"`
	CEP   string `json:"cep"   validate:"required,len=8"`
}

type CodigoUsuarioDTO struct {
	CodigoAuth string `json:"codigo_auth" validate:"required,max=6"`
}

type CodigoConviteDTO struct {
	ConviteUUID string `json:"convite_uuid" validate:"required,max=100"`
}

type AcademiaDTO struct {
	CNPJ string `json:"cnpj" validate:"required,len=14"`
	Nome string `json:"nome" validate:"required,max=100"`
}

type RespostaControllerDTO struct {
	Mensagem string `json:"mensagem"`
}

type RespostaErroDTO struct {
	Mensagem string `json:"mensagem"`
}

type SolicitacaoEntradaDTO struct {
	IDSolicitacao string `json:"id_solicitacao"`
	Nome          string `json:"nome"`
	CPF           string `json:"cpf"`
	CEP           string `json:"cep"`
}

type AprovarSolicitacaoDTO struct {
	IDSolicitacao string `json:"id_solicitacao" validate:"required,max=100"`
}

type RemoverAlunoDTO struct {
	IDAluno string `json:"id_aluno" validate:"required,max=100"`
}

type CriarAulaDTO struct {
	Conteudo string `json:"conteudo" validate:"required,max=255"`
	DataAula string `json:"data_aula" validate:"required"`
}

type PresencaDTO struct {
	IDAula string `json:"id_aula" validate:"required,max=100"`
}

type ContagemPresencaDTO struct {
	IDAluno  string `json:"id_aluno"`
	Contagem int64  `json:"contagem"`
}

type AlunoDTO struct {
	IDAluno string `json:"id_aluno"`
	Nome    string `json:"nome"`
	CPF     string `json:"cpf"`
	Faixa   string `json:"faixa"`
}

type CriarInstrutorDTO struct {
	IDAluno string `json:"id_aluno" validate:"required,max=100"`
}

type AulaDTO struct {
	IDAula      string    `json:"id_aula"`
	DataAula    time.Time `json:"data_aula"`
	Conteudo    string    `json:"conteudo"`
	IDInstrutor string    `json:"id_instrutor"`
}
