package Presentation

import (
	"api-back-end/api/UsersCase"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
)

const baseURLConvite = "https://kravconnect.vercel.app/invite?invite="

func ControllerCadastro(response_write http.ResponseWriter, request *http.Request) {

	usuario_dto := Desserializar[UsuarioDTO](request)

	err := validator.New(validator.WithRequiredStructEnabled()).Struct(usuario_dto)

	if err != nil {
		BadRequest(response_write, err)
		return
	}

	if !DominioEmailValido(usuario_dto.Email) {
		BadRequest(response_write, errors.New("use um email gmail, hotmail ou outlook"))
		return
	}

	if err := UsersCase.PrepararUsuario(usuario_dto.CPF, usuario_dto.CEP,
		usuario_dto.Email, usuario_dto.Senha); err != nil {
		BadRequest(response_write, err)
		return
	}

	Status200(response_write, "Um codigo de autententicao foi enviado a seu Email")

}

func ControllerCadastroConfirmar(response http.ResponseWriter, request *http.Request) {
	codigo_dto := Desserializar[CodigoUsuarioDTO](request)

	err_ := validator.New(validator.WithRequiredStructEnabled()).Struct(codigo_dto)

	if err_ != nil {
		BadRequest(response, err_)
		return
	}

	err := UsersCase.CriarUsuario(codigo_dto.CodigoAuth)

	if err != nil {
		BadRequest(response, err)
		return
	}

	Status200(response, "Usuario autenticado e registrado com sucesso")
}

func ControllerLogin(response http.ResponseWriter, request *http.Request) {
	usuario := Desserializar[LoginDTO](request)

	err_ := validator.New(validator.WithRequiredStructEnabled()).Struct(usuario)

	if err_ != nil {
		BadRequest(response, err_)
		return
	}

	if !DominioEmailValido(usuario.Email) {
		BadRequest(response, errors.New("use um email gmail, hotmail ou outlook"))
		return
	}

	token, err := UsersCase.Login(usuario.Email, usuario.Senha)

	if err != nil {
		BadRequest(response, err)
		return
	}

	Status200(response, token)
}

func ControllerPerfilUsuario(response http.ResponseWriter, request *http.Request) {
	id_usuario, err := ValidarJwt(*request)

	if err != nil {
		BadRequest(response, err)
		return
	}

	usuario, role, faixa, academias := UsersCase.PerfilUsuario(id_usuario)

	academiasDTO := make([]AcademiaVinculoDTO, 0, len(academias))
	for _, a := range academias {
		academiasDTO = append(academiasDTO, AcademiaVinculoDTO{
			ID:      a.ID,
			Nome:    a.Nome,
			CNPJ:    a.CNPJ,
			Vinculo: a.Vinculo,
			Faixa:   a.Faixa,
		})
	}

	Status200(response, PerfilUsuarioDTO{
		ID:             usuario.ID,
		Nome:           usuario.Nome,
		Email:          usuario.Email,
		Genero:         usuario.Genero,
		CPF:            usuario.CPF,
		DataNascimento: usuario.DataNascimento,
		EnderecoID:     usuario.EnderecoID,
		Role:           role,
		Faixa:          faixa,
		Academias:      academiasDTO,
	})
}

func ControllerEsqueciSenha(response http.ResponseWriter, request *http.Request) {
	dto := Desserializar[EsqueciSenhaDTO](request)
	if error_ := validator.New(validator.WithRequiredStructEnabled()).Struct(dto); error_ != nil {
		BadRequest(response, error_)
		return
	}

	if err := UsersCase.EsqueciSenha(dto.Email); err != nil {
		BadRequest(response, err)
		return
	}

	Status200(response, "Se o email existir, enviamos um codigo para redefinir a senha")
}

func ControllerRedefinirSenha(response http.ResponseWriter, request *http.Request) {
	dto := Desserializar[RedefinirSenhaDTO](request)
	if error_ := validator.New(validator.WithRequiredStructEnabled()).Struct(dto); error_ != nil {
		BadRequest(response, error_)
		return
	}

	if err := UsersCase.RedefinirSenha(dto.CodigoAuth, dto.NovaSenha); err != nil {
		BadRequest(response, err)
		return
	}

	Status200(response, "Senha redefinida com sucesso")
}

func ControllerCriarAcademia(response http.ResponseWriter, request *http.Request) {
	id_usuario, err := ValidarJwt(*request)

	if err != nil {
		BadRequest(response, err)
		return
	}

	academia_dto := Desserializar[AcademiaDTO](request)

	error_ := validator.New(validator.WithRequiredStructEnabled()).Struct(academia_dto)

	if error_ != nil {
		BadRequest(response, error_)
		return
	}

	if !CnpjValido(academia_dto.CNPJ) {
		BadRequest(response, errors.New("CNPJ invalido"))
		return
	}

	errw := UsersCase.CriarAcademia(academia_dto.CNPJ, academia_dto.Nome, id_usuario)

	if errw != nil {
		BadRequest(response, errw)
		return
	}

}

func ControllerGerarConvites(response http.ResponseWriter, request *http.Request) {
	id_usuario, err := ValidarJwt(*request)

	if err != nil {
		BadRequest(response, err)
		return
	}

	convite, err := UsersCase.GerarConvite(id_usuario)

	if err != nil {
		BadRequest(response, err)
		return
	}

	Status200(response, baseURLConvite+convite.ChaveConvite)

}

func ControllerMostrarConvites(response http.ResponseWriter, request *http.Request) {
	id_usuario, err := ValidarJwt(*request)

	if err != nil {
		BadRequest(response, err)
		return
	}

	convites, err_ := UsersCase.MostrarConvites(id_usuario)

	if err_ != nil {
		BadRequest(response, err_)
		return
	}

	Status200(response, MapearConvites(convites))
}

func ControllerDeletarConvite(response http.ResponseWriter, request *http.Request) {
	id_usuario, err := ValidarJwt(*request)

	if err != nil {
		BadRequest(response, err)
		return
	}

	id_convite := request.PathValue("id_convite")
	if id_convite == "" {
		BadRequest(response, errors.New("id_convite e obrigatorio"))
		return
	}

	if err := UsersCase.DeletarConvite(id_usuario, id_convite); err != nil {
		BadRequest(response, err)
		return
	}

	Status200(response, "Convite removido com sucesso")
}

func ControllerSolicitarEntrada(response http.ResponseWriter, request *http.Request) {
	id_usuario, err := ValidarJwt(*request)

	if err != nil {
		BadRequest(response, err)
		return
	}

	convite_dto := Desserializar[CodigoConviteDTO](request)
	error_ := validator.New(validator.WithRequiredStructEnabled()).Struct(convite_dto)

	if error_ != nil {
		BadRequest(response, error_)
		return
	}

	err = UsersCase.SolicitarEntrada(convite_dto.ConviteUUID, id_usuario)

	if err != nil {
		BadRequest(response, err)
		return
	}

	Status200(response, "Solicitacao feita pra entrar na academia, aguarde a aprovacao do professor")

}

func ControllerListarSolicitacoes(response http.ResponseWriter, request *http.Request) {
	id_usuario, err := ValidarJwt(*request)

	if err != nil {
		BadRequest(response, err)
		return
	}

	solicitacoes, err := UsersCase.ListarSolicitacoesEntrada(id_usuario)

	if err != nil {
		BadRequest(response, err)
		return
	}

	Status200(response, MapearSolicitacoes(solicitacoes))

}

func ControllerAprovarSolicitacao(response http.ResponseWriter, request *http.Request) {
	id_usuario, err := ValidarJwt(*request)

	if err != nil {
		BadRequest(response, err)
		return
	}

	aprovar_dto := Desserializar[AprovarSolicitacaoDTO](request)
	error_ := validator.New(validator.WithRequiredStructEnabled()).Struct(aprovar_dto)

	if error_ != nil {
		BadRequest(response, error_)
		return
	}

	if err := UsersCase.AprovarSolicitacaoEntrada(id_usuario, aprovar_dto.IDSolicitacao); err != nil {
		BadRequest(response, err)
		return
	}

	Status200(response, "Aluno aprovado e adicionado a academia")

}

func ControllerListarAlunos(response http.ResponseWriter, request *http.Request) {
	id_usuario, err := ValidarJwt(*request)

	if err != nil {
		BadRequest(response, err)
		return
	}

	alunos, err := UsersCase.ListarAlunos(id_usuario)

	if err != nil {
		BadRequest(response, err)
		return
	}

	Status200(response, MapearAlunos(alunos))

}

func ControllerRemoverAluno(response http.ResponseWriter, request *http.Request) {
	id_usuario, err := ValidarJwt(*request)

	if err != nil {
		BadRequest(response, err)
		return
	}

	id_aluno := request.PathValue("id_aluno")
	if id_aluno == "" {
		BadRequest(response, errors.New("id_aluno e obrigatorio"))
		return
	}

	if err := UsersCase.RemoverAluno(id_usuario, id_aluno); err != nil {
		BadRequest(response, err)
		return
	}

	Status200(response, "Aluno removido da academia")

}

func ControllerCriarAula(response http.ResponseWriter, request *http.Request) {
	id_usuario, err := ValidarJwt(*request)

	if err != nil {
		BadRequest(response, err)
		return
	}

	aula_dto := Desserializar[CriarAulaDTO](request)
	error_ := validator.New(validator.WithRequiredStructEnabled()).Struct(aula_dto)

	if error_ != nil {
		BadRequest(response, error_)
		return
	}

	if err := UsersCase.CriarAula(id_usuario, aula_dto.Conteudo, aula_dto.DataAula, aula_dto.Faixa); err != nil {
		BadRequest(response, err)
		return
	}

	Status200(response, "Aula criada com sucesso")

}

func ControllerRegistrarPresenca(response http.ResponseWriter, request *http.Request) {
	id_usuario, err := ValidarJwt(*request)

	if err != nil {
		BadRequest(response, err)
		return
	}

	presenca_dto := Desserializar[PresencaDTO](request)
	error_ := validator.New(validator.WithRequiredStructEnabled()).Struct(presenca_dto)

	if error_ != nil {
		BadRequest(response, error_)
		return
	}

	if err := UsersCase.RegistrarPresenca(id_usuario, presenca_dto.IDAula, presenca_dto.Latitude, presenca_dto.Longitude); err != nil {
		BadRequest(response, err)
		return
	}

	Status200(response, "Presenca registrada com sucesso")

}

func ControllerContarPresencasAluno(response http.ResponseWriter, request *http.Request) {
	id_usuario, err := ValidarJwt(*request)

	if err != nil {
		BadRequest(response, err)
		return
	}

	id_aluno := request.URL.Query().Get("id_aluno")

	if id_aluno == "" {
		BadRequest(response, errors.New("id_aluno e obrigatorio"))
		return
	}

	presencas, err := UsersCase.ContarPresencasAluno(id_usuario, id_aluno)
	if err != nil {
		BadRequest(response, err)
		return
	}

	Status200(response, MapearPresencasAluno(id_aluno, presencas))

}

func ControllerAtualizarFaixaAluno(response http.ResponseWriter, request *http.Request) {
	id_usuario, err := ValidarJwt(*request)

	if err != nil {
		BadRequest(response, err)
		return
	}

	id_aluno := request.PathValue("id_aluno")
	if id_aluno == "" {
		BadRequest(response, errors.New("id_aluno e obrigatorio"))
		return
	}

	dto := Desserializar[AtualizarFaixaDTO](request)
	if error_ := validator.New(validator.WithRequiredStructEnabled()).Struct(dto); error_ != nil {
		BadRequest(response, error_)
		return
	}

	if err := UsersCase.AtualizarFaixaAluno(id_usuario, id_aluno, dto.Faixa); err != nil {
		BadRequest(response, err)
		return
	}

	Status200(response, "Faixa do aluno atualizada com sucesso")

}

func ControllerCriarInstrutor(response http.ResponseWriter, request *http.Request) {
	id_usuario, err := ValidarJwt(*request)

	if err != nil {
		BadRequest(response, err)
		return
	}

	instrutor_dto := Desserializar[CriarInstrutorDTO](request)
	error_ := validator.New(validator.WithRequiredStructEnabled()).Struct(instrutor_dto)

	if error_ != nil {
		BadRequest(response, error_)
		return
	}

	if err := UsersCase.CriarInstrutor(id_usuario, instrutor_dto.IDAluno); err != nil {
		BadRequest(response, err)
		return
	}

	Status200(response, "Instrutor criado com sucesso")

}

func ControllerCriarProduto(response http.ResponseWriter, request *http.Request) {
	id_usuario, err := ValidarJwt(*request)
	if err != nil {
		BadRequest(response, err)
		return
	}

	dto := Desserializar[CriarProdutoDTO](request)
	if error_ := validator.New(validator.WithRequiredStructEnabled()).Struct(dto); error_ != nil {
		BadRequest(response, error_)
		return
	}

	if err := UsersCase.CriarProduto(id_usuario, dto.Nome, dto.Preco, dto.Tamanho, dto.Quantidade, dto.ImagemURL); err != nil {
		BadRequest(response, err)
		return
	}

	Status200(response, "Produto criado com sucesso")
}

func ControllerAtualizarProduto(response http.ResponseWriter, request *http.Request) {
	id_usuario, err := ValidarJwt(*request)
	if err != nil {
		BadRequest(response, err)
		return
	}

	dto := Desserializar[AtualizarProdutoDTO](request)
	if error_ := validator.New(validator.WithRequiredStructEnabled()).Struct(dto); error_ != nil {
		BadRequest(response, error_)
		return
	}

	if err := UsersCase.AtualizarProduto(id_usuario, dto.IDProduto, dto.Nome, dto.Preco, dto.Tamanho, dto.Quantidade, dto.ImagemURL); err != nil {
		BadRequest(response, err)
		return
	}

	Status200(response, "Produto atualizado com sucesso")
}

func ControllerDeletarProduto(response http.ResponseWriter, request *http.Request) {
	id_usuario, err := ValidarJwt(*request)
	if err != nil {
		BadRequest(response, err)
		return
	}

	id_produto := request.PathValue("id_produto")
	if id_produto == "" {
		BadRequest(response, errors.New("id_produto e obrigatorio"))
		return
	}

	if err := UsersCase.DeletarProduto(id_usuario, id_produto); err != nil {
		BadRequest(response, err)
		return
	}

	Status200(response, "Produto removido com sucesso")
}

func ControllerListarCatalogo(response http.ResponseWriter, request *http.Request) {
	id_usuario, err := ValidarJwt(*request)
	if err != nil {
		BadRequest(response, err)
		return
	}

	produtos, err := UsersCase.ListarCatalogo(id_usuario)
	if err != nil {
		BadRequest(response, err)
		return
	}

	Status200(response, MapearProdutos(produtos))
}

func ControllerSinalizarInteresse(response http.ResponseWriter, request *http.Request) {
	id_usuario, err := ValidarJwt(*request)
	if err != nil {
		BadRequest(response, err)
		return
	}

	dto := Desserializar[SinalizarInteresseDTO](request)
	if error_ := validator.New(validator.WithRequiredStructEnabled()).Struct(dto); error_ != nil {
		BadRequest(response, error_)
		return
	}

	if err := UsersCase.SinalizarInteresse(id_usuario, dto.IDProduto, dto.Quantidade); err != nil {
		BadRequest(response, err)
		return
	}

	Status200(response, "Interesse sinalizado com sucesso")
}

func ControllerRegistrarLocalizacaoAcademia(response http.ResponseWriter, request *http.Request) {
	id_usuario, err := ValidarJwt(*request)
	if err != nil {
		BadRequest(response, err)
		return
	}

	localizacao_dto := Desserializar[LocalizacaoAcademiaDTO](request)
	error_ := validator.New(validator.WithRequiredStructEnabled()).Struct(localizacao_dto)
	if error_ != nil {
		BadRequest(response, error_)
		return
	}

	if err := UsersCase.RegistrarLocalizacaoAcademia(id_usuario, localizacao_dto.Latitude, localizacao_dto.Longitude); err != nil {
		BadRequest(response, err)
		return
	}

	Status200(response, "Localizacao da academia registrada com sucesso")
}

func ControllerLocalizacaoAcademia(response http.ResponseWriter, request *http.Request) {
	id_usuario, err := ValidarJwt(*request)

	if err != nil {
		BadRequest(response, err)
		return
	}

	latitude, longitude, err := UsersCase.LocalizacaoAcademia(id_usuario)
	if err != nil {
		BadRequest(response, err)
		return
	}

	Status200(response, LocalizacaoAcademiaDTO{Latitude: latitude, Longitude: longitude})
}

func ControllerRealizarPagamento(response http.ResponseWriter, request *http.Request) {
	id_usuario, err := ValidarJwt(*request)
	if err != nil {
		BadRequest(response, err)
		return
	}

	pagamento_dto := Desserializar[PagamentoDTO](request)
	error_ := validator.New(validator.WithRequiredStructEnabled()).Struct(pagamento_dto)
	if error_ != nil {
		BadRequest(response, error_)
		return
	}

	id_pagamento, client_secret, err := UsersCase.RealizarPagamento(id_usuario, pagamento_dto.ValorCentavos)
	if err != nil {
		BadRequest(response, err)
		return
	}

	Status200(response, RespostaPagamentoDTO{IDPagamento: id_pagamento, ClientSecret: client_secret})
}

func ControllerListarAulasDoDia(response http.ResponseWriter, request *http.Request) {
	id_usuario, err := ValidarJwt(*request)

	if err != nil {
		BadRequest(response, err)
		return
	}

	aulas, err := UsersCase.ListarAulasDoDia(id_usuario)

	if err != nil {
		BadRequest(response, err)
		return
	}

	Status200(response, MapearAulas(aulas))

}

func ControllerListarAulasProfessor(response http.ResponseWriter, request *http.Request) {
	id_usuario, err := ValidarJwt(*request)

	if err != nil {
		BadRequest(response, err)
		return
	}

	data := request.URL.Query().Get("data")

	aulas, err := UsersCase.ListarAulasProfessor(id_usuario, data)

	if err != nil {
		BadRequest(response, err)
		return
	}

	Status200(response, MapearAulas(aulas))

}
