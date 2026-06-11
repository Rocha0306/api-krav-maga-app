package Presentation

import (
	"time"
)

type LoginDTO struct {
	Email string `json:"email" validate:"required,email"`
	Senha string `json:"senha" validate:"required"`
}

type UsuarioDTO struct {
	Email string `json:"email" validate:"required,max=100,email"`
	Senha string `json:"senha" validate:"required,max=20"`
	CPF   string `json:"cpf"   validate:"required,len=11"`
	CEP   string `json:"cep"   validate:"required,len=8"`
}

type CodigoUsuarioDTO struct {
	CodigoAuth string `json:"codigo_auth" validate:"required,max=6"`
}

type EsqueciSenhaDTO struct {
	Email string `json:"email" validate:"required,email"`
}

type RedefinirSenhaDTO struct {
	CodigoAuth string `json:"codigo_auth" validate:"required,max=6"`
	NovaSenha  string `json:"nova_senha" validate:"required,max=20"`
}

type PerfilUsuarioDTO struct {
	ID             string    `json:"id"`
	Nome           string    `json:"name"`
	Email          string    `json:"email"`
	Genero         string    `json:"genero"`
	CPF            string    `json:"cpf"`
	DataNascimento time.Time `json:"data_nascimento"`
	EnderecoID     string    `json:"endereco_id"`
	Role           string    `json:"role"`
	Faixa          string    `json:"faixa"`
}

type CodigoConviteDTO struct {
	ConviteUUID string `json:"convite_uuid" validate:"required,uuid"`
}

type ConviteResponseDTO struct {
	IDConvite string `json:"id_convite"`
	Link      string `json:"link"`
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

type CriarAulaDTO struct {
	Conteudo string `json:"conteudo" validate:"required,max=255"`
	DataAula string `json:"data_aula" validate:"required"`
	Faixa    string `json:"faixa" validate:"required,max=15"`
}

type AtualizarFaixaDTO struct {
	Faixa string `json:"faixa" validate:"required,max=15"`
}

type PresencaDTO struct {
	IDAula    string  `json:"id_aula"    validate:"required,max=100"`
	Latitude  float64 `json:"latitude"   validate:"required"`
	Longitude float64 `json:"longitude"  validate:"required"`
}

type LocalizacaoAcademiaDTO struct {
	Latitude  float64 `json:"latitude"  validate:"required"`
	Longitude float64 `json:"longitude" validate:"required"`
}

type PresencaDetalheDTO struct {
	NomeAula string    `json:"nome_aula"`
	Data     time.Time `json:"data"`
}

type ContagemPresencaDTO struct {
	IDAluno   string               `json:"id_aluno"`
	Contagem  int64                `json:"contagem"`
	Presencas []PresencaDetalheDTO `json:"presencas"`
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
	Faixa       string    `json:"faixa"`
	IDInstrutor string    `json:"id_instrutor"`
}

type CriarProdutoDTO struct {
	Nome       string  `json:"nome"       validate:"required,max=100"`
	Preco      float64 `json:"preco"      validate:"required,max=5000"`
	Tamanho    string  `json:"tamanho"    validate:"required,max=20"`
	Quantidade int     `json:"quantidade" validate:"required,min=1,max=20"`
}

type AtualizarProdutoDTO struct {
	IDProduto  string  `json:"id_produto"  validate:"required,max=100"`
	Nome       string  `json:"nome"        validate:"required,max=100"`
	Preco      float64 `json:"preco"       validate:"required"`
	Tamanho    string  `json:"tamanho"     validate:"required,max=20"`
	Quantidade int     `json:"quantidade"  validate:"required,min=1"`
}

type ProdutoDTO struct {
	IDProduto  string  `json:"id_produto"`
	Nome       string  `json:"nome"`
	Preco      float64 `json:"preco"`
	Tamanho    string  `json:"tamanho"`
	Quantidade int     `json:"quantidade"`
}

type SinalizarInteresseDTO struct {
	IDProduto  string `json:"id_produto"  validate:"required,max=100"`
	Quantidade int    `json:"quantidade"  validate:"required,min=1"`
}

type PagamentoDTO struct {
	ValorCentavos int64 `json:"valor_centavos" validate:"required,min=100"`
}

type RespostaPagamentoDTO struct {
	IDPagamento  string `json:"id_pagamento"`
	ClientSecret string `json:"client_secret"`
}
