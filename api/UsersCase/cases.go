package UsersCase

import (
	entities "api-back-end/api/Entities"
	"api-back-end/api/InterfaceAdapters"
	"api-back-end/api/Repository"
	"errors"
	"fmt"
	"strconv"
	"time"
)

/*
1 - Pega o Response da API de CPF e CEP
2 - Popula o struct que e a entidade no banco
3 - Insere no banco o Cep e coloca o Usuario em cache
*/
func PrepararUsuario(cpf string, cep string, email string, senha string) error {

	if Repository.Existe[entities.Usuarios]("email", email) {
		return errors.New("Usuario ja cadastrado")
	}

	err, resposta_api_cpf := InterfaceAdapters.CpfApi(cpf)
	if err != nil {
		return err
	}

	err, resposta_api_cep := InterfaceAdapters.CepApi(cep)
	if err != nil {
		return err
	}

	entidade_final := InterfaceAdapters.MapearUsuario(email, senha, resposta_api_cpf, resposta_api_cep)
	Repository.Inserir(&entidade_final.Endereco)
	numero_auth := strconv.Itoa(InterfaceAdapters.GerarNumeroAuth())
	InterfaceAdapters.SalvarUsuarioCache(numero_auth, *entidade_final)
	InterfaceAdapters.EnviarEmail(fmt.Sprintf("Codigo de Autenticacao: %s", numero_auth), []string{entidade_final.Email})
	return nil
}

/*
1. Pega o Usuario do cache
2. Insere no banco
*/
func CriarUsuario(codigo_chave string) error {
	usuario_banco, err := InterfaceAdapters.PegarUsuarioCache(codigo_chave)
	if err != nil {
		return errors.New("codigo de autenticacao invalido ou expirado")
	}

	Repository.Inserir(&usuario_banco)
	InterfaceAdapters.RemoverValorCache(codigo_chave)
	return nil
}

func Login(email string, senha string) (string, error) {
	usuario := Repository.BuscarOnde[entities.Usuarios]("email", email)
	senha_hash := InterfaceAdapters.HashSenha(senha)

	if usuario.Email != email && usuario.Senha != senha_hash {
		return "", errors.New("Aluno nao existe no sistema")
	}

	return InterfaceAdapters.GerarTokenJwt(usuario.ID), nil
}

func CriarAcademia(cnpj string, nome_academia string, id_usuario string) error {
	entidade_academia := InterfaceAdapters.MapearAcademia(cnpj, nome_academia)
	Repository.Inserir(entidade_academia)
	entidade_professor := InterfaceAdapters.MapearProfessor(id_usuario, entidade_academia.ID)
	Repository.Inserir(&entidade_professor)
	return nil
}

func GerarConvite(id_usuario string) error {
	if !Repository.Existe[entities.Professores]("id_usuario_professor", id_usuario) {
		return errors.New("Voce nao pode criar convites ja que nao eh professor")
	}

	professor := Repository.BuscarOnde[entities.Professores]("id_usuario_professor", id_usuario)
	entidade_convite := InterfaceAdapters.MapearConvite(professor.IDAcademiaProfessor)
	Repository.Inserir(entidade_convite)
	return nil
}

func MostrarConvites(id_usuario string) (entities.Convites, error) {
	if !Repository.Existe[entities.Professores]("id_usuario_professor", id_usuario) {
		return entities.Convites{}, errors.New("Voce nao pode visualizar convites ja que nao eh um professor")
	}

	professor := Repository.BuscarOnde[entities.Professores]("id_usuario_professor", id_usuario)
	convite := Repository.BuscarOnde[entities.Convites]("id_academia", professor.IDAcademiaProfessor)
	return *convite, nil
}

func SolicitarEntrada(codigo_convite string, id_usuario string) error {
	if !Repository.Existe[entities.Convites]("chave_convite", codigo_convite) {
		return errors.New("convite nao encontrado")
	}

	convite := Repository.BuscarOnde[entities.Convites]("chave_convite", codigo_convite)
	solicitacao := InterfaceAdapters.MapearSolicitacaoConvite(convite.IDAcademia, id_usuario)
	Repository.Inserir(solicitacao)
	return nil
}

func ListarSolicitacoesEntrada(id_usuario string) ([]entities.SolicitacoesConvite, error) {
	if !Repository.Existe[entities.Professores]("id_usuario_professor", id_usuario) {
		return nil, errors.New("Voce nao pode visualizar convites ja que nao eh um professor")
	}

	professor := Repository.BuscarOnde[entities.Professores]("id_usuario_professor", id_usuario)
	return Repository.BuscarSolicitacoesPorAcademia(professor.IDAcademiaProfessor), nil
}

func AprovarSolicitacaoEntrada(id_usuario string, id_solicitacao string) error {
	professor := Repository.BuscarOnde[entities.Professores]("id_usuario_professor", id_usuario)
	if professor.IDProfessor == "" {
		return errors.New("voce nao pode aprovar ja que nao eh professor")
	}

	solicitacao := Repository.BuscarOnde[entities.SolicitacoesConvite]("id_solicitacao", id_solicitacao)
	if solicitacao.IDSolicitacao == "" {
		return errors.New("solicitacao nao encontrada")
	}

	if solicitacao.IDAcademia != professor.IDAcademiaProfessor {
		return errors.New("essa solicitacao nao pertence a sua academia")
	}

	aluno := InterfaceAdapters.MapearAlunoAprovado(solicitacao.IDUsuario, solicitacao.IDAcademia)
	Repository.Inserir(&aluno)
	Repository.Deletar[entities.SolicitacoesConvite]("id_solicitacao", id_solicitacao)
	return nil
}

func ListarAlunos(id_usuario string) ([]entities.Alunos, error) {
	professor := Repository.BuscarOnde[entities.Professores]("id_usuario_professor", id_usuario)
	if professor.IDProfessor == "" {
		return nil, errors.New("voce nao pode visualizar alunos ja que nao eh um professor")
	}

	return Repository.BuscarAlunosPorAcademia(professor.IDAcademiaProfessor), nil
}

func RemoverAluno(id_usuario string, id_aluno string) error {
	professor := Repository.BuscarOnde[entities.Professores]("id_usuario_professor", id_usuario)
	if professor.IDProfessor == "" {
		return errors.New("voce nao pode remover alunos ja que nao eh um professor")
	}

	aluno := Repository.BuscarOnde[entities.Alunos]("id_aluno", id_aluno)
	if aluno.IDAluno == "" {
		return errors.New("aluno nao encontrado")
	}

	if aluno.IDAcademiaAluno != professor.IDAcademiaProfessor {
		return errors.New("esse aluno nao pertence a sua academia")
	}

	Repository.Deletar[entities.Alunos]("id_aluno", id_aluno)
	return nil
}

func CriarAula(id_usuario string, conteudo string, data_aula string) error {
	id_academia := ""

	professor := Repository.BuscarOnde[entities.Professores]("id_usuario_professor", id_usuario)
	if professor.IDProfessor != "" {
		id_academia = professor.IDAcademiaProfessor
	}

	if id_academia == "" {
		instrutor := Repository.BuscarOnde[entities.Instrutores]("id_usuario_instrutor", id_usuario)
		if instrutor.IDInstrutor != "" {
			id_academia = instrutor.IDAcademiaInstrutor
		}
	}

	if id_academia == "" {
		return errors.New("voce nao pode criar aulas ja que nao eh professor nem instrutor")
	}

	data, err := time.Parse("2006-01-02 15:04", data_aula)
	if err != nil {
		return errors.New("data da aula invalida, use o formato AAAA-MM-DD HH:MM")
	}

	entidade_aula := InterfaceAdapters.MapearAula(conteudo, id_academia, id_usuario, data)
	Repository.Inserir(entidade_aula)
	return nil
}

func RegistrarPresenca(id_usuario string, id_aula string) error {
	aluno := Repository.BuscarOnde[entities.Alunos]("id_usuario_aluno", id_usuario)
	if aluno.IDAluno == "" {
		return errors.New("voce nao eh aluno de nenhuma academia")
	}

	aula := Repository.BuscarOnde[entities.Aulas]("id_aula", id_aula)
	if aula.IDAula == "" {
		return errors.New("aula nao encontrada")
	}

	if aula.IDAcademia != aluno.IDAcademiaAluno {
		return errors.New("essa aula nao eh da sua academia")
	}

	presenca := InterfaceAdapters.MapearPresenca(aluno.IDAluno, id_aula)
	Repository.Inserir(presenca)
	return nil
}

func ContarPresencasAluno(id_aluno string) int64 {
	return Repository.Contar[entities.Presencas]("id_aluno", id_aluno)
}

func CriarInstrutor(id_usuario string, id_aluno string) error {
	professor := Repository.BuscarOnde[entities.Professores]("id_usuario_professor", id_usuario)
	if professor.IDProfessor == "" {
		return errors.New("voce nao pode criar instrutores ja que nao eh um professor")
	}

	aluno := Repository.BuscarOnde[entities.Alunos]("id_aluno", id_aluno)
	if aluno.IDAluno == "" {
		return errors.New("aluno nao encontrado")
	}

	if aluno.IDAcademiaAluno != professor.IDAcademiaProfessor {
		return errors.New("esse aluno nao pertence a sua academia")
	}

	instrutor := InterfaceAdapters.MapearInstrutor(aluno.IDUsuarioAluno, aluno.IDAcademiaAluno)
	Repository.Inserir(instrutor)
	return nil
}

func ListarAulasDoDia(id_usuario string, data string) ([]entities.Aulas, error) {
	id_academia := ""

	aluno := Repository.BuscarOnde[entities.Alunos]("id_usuario_aluno", id_usuario)
	if aluno.IDAluno != "" {
		id_academia = aluno.IDAcademiaAluno
	}

	if id_academia == "" {
		professor := Repository.BuscarOnde[entities.Professores]("id_usuario_professor", id_usuario)
		if professor.IDProfessor != "" {
			id_academia = professor.IDAcademiaProfessor
		}
	}

	if id_academia == "" {
		return nil, errors.New("voce nao esta vinculado a nenhuma academia")
	}

	dia := time.Now()
	if data != "" {
		parsed, err := time.Parse("2006-01-02", data)
		if err != nil {
			return nil, errors.New("data invalida, use o formato AAAA-MM-DD")
		}
		dia = parsed
	}

	inicio := time.Date(dia.Year(), dia.Month(), dia.Day(), 0, 0, 0, 0, dia.Location())
	fim := inicio.Add(24 * time.Hour)

	return Repository.BuscarAulasDoDia(id_academia, inicio, fim), nil
}
