package Presentation

import (
	entities "api-back-end/api/Entities"
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

func MapearAulas(aulas []entities.Aulas) []AulaDTO {
	resultado := make([]AulaDTO, 0, len(aulas))
	for _, aula := range aulas {
		resultado = append(resultado, AulaDTO{
			IDAula:      aula.IDAula,
			DataAula:    aula.DataAula,
			Conteudo:    aula.Conteudo,
			IDInstrutor: aula.IDInstrutor,
		})
	}
	return resultado
}

func MapearAlunos(alunos []entities.Alunos) []AlunoDTO {
	resultado := make([]AlunoDTO, 0, len(alunos))
	for _, aluno := range alunos {
		resultado = append(resultado, AlunoDTO{
			IDAluno: aluno.IDAluno,
			Nome:    aluno.Usuario.Nome,
			CPF:     aluno.Usuario.CPF,
			Faixa:   aluno.Faixa,
		})
	}
	return resultado
}

func MapearSolicitacoes(solicitacoes []entities.SolicitacoesConvite) []SolicitacaoEntradaDTO {
	resultado := make([]SolicitacaoEntradaDTO, 0, len(solicitacoes))
	for _, solicitacao := range solicitacoes {
		resultado = append(resultado, SolicitacaoEntradaDTO{
			IDSolicitacao: solicitacao.IDSolicitacao,
			Nome:          solicitacao.Usuario.Nome,
			CPF:           solicitacao.Usuario.CPF,
			CEP:           solicitacao.Usuario.Endereco.CEP,
		})
	}
	return resultado
}
