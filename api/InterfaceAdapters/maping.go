package InterfaceAdapters

import (
	entities "api-back-end/api/Entities"
	"fmt"
	"time"
)

func MapearUsuario(email string, senha string, dados_usuario CpfApiResponse, dados_cep CepApiResponse) *entities.Usuarios {
	endereco := &entities.Endereco{}
	endereco.ID = GerarId()
	endereco.CEP = dados_cep.CEP
	endereco.Logradouro = dados_cep.Logradouro
	endereco.Complemento = dados_cep.Complemento
	endereco.Unidade = dados_cep.Unidade
	endereco.Bairro = dados_cep.Bairro
	endereco.Localidade = dados_cep.Localidade
	endereco.UF = dados_cep.UF
	endereco.Estado = dados_cep.Estado
	endereco.Regiao = dados_cep.Regiao
	endereco.DDD = dados_cep.DDD

	usuario := &entities.Usuarios{}
	usuario.ID = GerarId()
	usuario.Nome = dados_usuario.Data.Name
	usuario.CPF = dados_usuario.Data.CPF
	usuario.Genero = dados_usuario.Data.Gender
	data, _ := time.Parse("02/01/2006", dados_usuario.Data.BirthDate)
	usuario.DataNascimento = data
	usuario.Email = email
	usuario.Senha = HashSenha(senha)
	usuario.EnderecoID = endereco.ID
	usuario.Endereco = *endereco

	return usuario
}

func MapearAcademia(cnpj string, nome string) *entities.Academias {
	academia := &entities.Academias{
		ID:   GerarId(),
		CNPJ: cnpj,
		Nome: nome,
	}

	return academia

}

func MapearProfessor(id_usuario string, id_academia string) entities.Professores {
	professor := entities.Professores{
		IDProfessor:         GerarId(),
		IDAcademiaProfessor: id_academia,
		IDUsuarioProfessor:  id_usuario,
	}

	return professor
}

func MapearConvite(id_academia string) *entities.Convites {
	return &entities.Convites{
		IDConvite:    GerarId(),
		IDAcademia:   id_academia,
		ChaveConvite: fmt.Sprintf("URLFrontComParametro?invite=%s", GerarUUID()),
	}
}

func MapearSolicitacaoConvite(id_academia string, id_usuario string) *entities.SolicitacoesConvite {
	return &entities.SolicitacoesConvite{
		IDSolicitacao: GerarId(),
		IDUsuario:     id_usuario,
		IDAcademia:    id_academia,
	}
}

// Uma aula (sessao): data, conteudo dado, academia e instrutor.
func MapearAula(conteudo string, id_academia string, id_instrutor string, data_aula time.Time) *entities.Aulas {
	return &entities.Aulas{
		IDAula:      GerarId(),
		DataAula:    data_aula,
		Conteudo:    conteudo,
		IDAcademia:  id_academia,
		IDInstrutor: id_instrutor,
	}
}

// Check-in do aluno numa aula. checkin_em = momento do registro.
func MapearPresenca(id_aluno string, id_aula string) *entities.Presencas {
	return &entities.Presencas{
		IDPresenca: GerarId(),
		IDAluno:    id_aluno,
		IDAula:     id_aula,
		CheckinEm:  time.Now(),
	}
}

// Promove um usuario (aluno) a instrutor da academia.
func MapearInstrutor(id_usuario string, id_academia string) *entities.Instrutores {
	return &entities.Instrutores{
		IDInstrutor:         GerarId(),
		IDUsuarioInstrutor:  id_usuario,
		IDAcademiaInstrutor: id_academia,
	}
}

func MapearProduto(nome string, preco float64, tamanho string, quantidade int, id_academia string) *entities.Produtos {
	return &entities.Produtos{
		IDProduto:  GerarId(),
		Nome:       nome,
		Preco:      preco,
		Tamanho:    tamanho,
		Quantidade: quantidade,
		IDAcademia: id_academia,
	}
}

func MapearInteresse(id_aluno string, id_produto string, quantidade int) *entities.Interesses {
	return &entities.Interesses{
		IDInteresse: GerarId(),
		IDAluno:     id_aluno,
		IDProduto:   id_produto,
		Quantidade:  quantidade,
	}
}

func MapearPagamento(id_aluno string, id_payment_intent string, valor_centavos int64) *entities.Pagamentos {
	return &entities.Pagamentos{
		IDPagamento:     GerarId(),
		IDAluno:         id_aluno,
		ValorCentavos:   valor_centavos,
		Status:          "pendente",
		IDPaymentIntent: id_payment_intent,
		CriadoEm:        time.Now(),
	}
}

// Monta o aluno que entra na academia quando o professor aprova a solicitacao.
func MapearAlunoAprovado(id_usuario string, id_academia string) entities.Alunos {
	return entities.Alunos{
		IDAluno:         GerarId(),
		Faixa:           "Branca",
		IDUsuarioAluno:  id_usuario,
		IDAcademiaAluno: id_academia,
	}
}
