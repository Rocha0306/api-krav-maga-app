package Presentation

import entities "api-back-end/api/Entities"

func MapearAulas(aulas []entities.Aulas) []AulaDTO {
	resultado := make([]AulaDTO, 0, len(aulas))
	for _, aula := range aulas {
		resultado = append(resultado, AulaDTO{
			IDAula:      aula.IDAula,
			DataAula:    aula.DataAula,
			Conteudo:    aula.Conteudo,
			Faixa:       aula.Faixa,
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

func MapearProdutos(produtos []entities.Produtos) []ProdutoDTO {
	resultado := make([]ProdutoDTO, 0, len(produtos))
	for _, p := range produtos {
		resultado = append(resultado, ProdutoDTO{
			IDProduto:  p.IDProduto,
			Nome:       p.Nome,
			Preco:      p.Preco,
			Tamanho:    p.Tamanho,
			Quantidade: p.Quantidade,
		})
	}
	return resultado
}

func MapearConvites(convites []entities.Convites) []ConviteResponseDTO {
	resultado := make([]ConviteResponseDTO, 0, len(convites))
	for _, convite := range convites {
		resultado = append(resultado, ConviteResponseDTO{
			IDConvite: convite.IDConvite,
			Link:      baseURLConvite + convite.ChaveConvite,
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
