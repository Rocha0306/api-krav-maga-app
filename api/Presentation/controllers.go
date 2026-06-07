package Presentation

import (
	"api-back-end/api/UsersCase"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ControllerCadastro(response_write http.ResponseWriter, request *http.Request) {

	usuario_dto := Desserializar[UsuarioDTO](request)

	err := validator.New(validator.WithRequiredStructEnabled()).Struct(usuario_dto)

	if err != nil {
		BadRequest(response_write, err)
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

	token, err := UsersCase.Login(usuario.Email, usuario.Senha)

	if err != nil {
		BadRequest(response, err)
		return
	}

	Status200(response, token)
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

	err = UsersCase.GerarConvite(id_usuario)

	if err != nil {
		BadRequest(response, err)
		return
	}

	Status200(response, "Convite criado")

}

func ControllerMostrarConvites(response http.ResponseWriter, request *http.Request) {
	id_usuario, err := ValidarJwt(*request)

	if err != nil {
		BadRequest(response, err)
		return
	}

	entidade, err_ := UsersCase.MostrarConvites(id_usuario)

	if err_ != nil {
		BadRequest(response, err_)
		return
	}

	Status200(response, entidade)
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
		BadRequest(response, error_)
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

	remover_dto := Desserializar[RemoverAlunoDTO](request)
	error_ := validator.New(validator.WithRequiredStructEnabled()).Struct(remover_dto)

	if error_ != nil {
		BadRequest(response, error_)
		return
	}

	if err := UsersCase.RemoverAluno(id_usuario, remover_dto.IDAluno); err != nil {
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

	if err := UsersCase.CriarAula(id_usuario, aula_dto.Conteudo, aula_dto.DataAula); err != nil {
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

	if err := UsersCase.RegistrarPresenca(id_usuario, presenca_dto.IDAula); err != nil {
		BadRequest(response, err)
		return
	}

	Status200(response, "Presenca registrada com sucesso")

}

func ControllerContarPresencasAluno(response http.ResponseWriter, request *http.Request) {
	_, err := ValidarJwt(*request)

	if err != nil {
		BadRequest(response, err)
		return
	}

	id_aluno := request.URL.Query().Get("id_aluno")

	if id_aluno == "" {
		BadRequest(response, errors.New("id_aluno e obrigatorio"))
		return
	}

	contagem := UsersCase.ContarPresencasAluno(id_aluno)
	Status200(response, ContagemPresencaDTO{IDAluno: id_aluno, Contagem: contagem})

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
