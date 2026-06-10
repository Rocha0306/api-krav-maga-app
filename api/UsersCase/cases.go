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

	if !InterfaceAdapters.DominioEmailValido(email) {
		return errors.New("use um email gmail, hotmail ou outlook")
	}

	if Repository.Exists[entities.Usuarios]("email", email) {
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
	if !InterfaceAdapters.DominioEmailValido(email) {
		return "", errors.New("use um email gmail, hotmail ou outlook")
	}

	usuario := Repository.SelectWhere[entities.Usuarios]("email", email)
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
	if !Repository.Exists[entities.Professores]("id_usuario_professor", id_usuario) {
		return errors.New("Voce nao pode criar convites ja que nao eh professor")
	}

	professor := Repository.SelectWhere[entities.Professores]("id_usuario_professor", id_usuario)
	entidade_convite := InterfaceAdapters.MapearConvite(professor.IDAcademiaProfessor)
	Repository.Inserir(entidade_convite)
	return nil
}

func MostrarConvites(id_usuario string) (entities.Convites, error) {
	if !Repository.Exists[entities.Professores]("id_usuario_professor", id_usuario) {
		return entities.Convites{}, errors.New("Voce nao pode visualizar convites ja que nao eh um professor")
	}

	professor := Repository.SelectWhere[entities.Professores]("id_usuario_professor", id_usuario)
	convite := Repository.SelectWhere[entities.Convites]("id_academia", professor.IDAcademiaProfessor)
	return *convite, nil
}

func SolicitarEntrada(codigo_convite string, id_usuario string) error {
	if !Repository.Exists[entities.Convites]("chave_convite", codigo_convite) {
		return errors.New("convite nao encontrado")
	}

	convite := Repository.SelectWhere[entities.Convites]("chave_convite", codigo_convite)
	solicitacao := InterfaceAdapters.MapearSolicitacaoConvite(convite.IDAcademia, id_usuario)
	Repository.Inserir(solicitacao)
	return nil
}

func ListarSolicitacoesEntrada(id_usuario string) ([]entities.SolicitacoesConvite, error) {
	if !Repository.Exists[entities.Professores]("id_usuario_professor", id_usuario) {
		return nil, errors.New("Voce nao pode visualizar convites ja que nao eh um professor")
	}

	professor := Repository.SelectWhere[entities.Professores]("id_usuario_professor", id_usuario)
	return Repository.SelectInvitesByGym(professor.IDAcademiaProfessor), nil
}

func AprovarSolicitacaoEntrada(id_usuario string, id_solicitacao string) error {
	professor := Repository.SelectWhere[entities.Professores]("id_usuario_professor", id_usuario)
	if professor.IDProfessor == "" {
		return errors.New("voce nao pode aprovar ja que nao eh professor")
	}

	solicitacao := Repository.SelectWhere[entities.SolicitacoesConvite]("id_solicitacao", id_solicitacao)
	if solicitacao.IDSolicitacao == "" {
		return errors.New("solicitacao nao encontrada")
	}

	if solicitacao.IDAcademia != professor.IDAcademiaProfessor {
		return errors.New("essa solicitacao nao pertence a sua academia")
	}

	aluno := InterfaceAdapters.MapearAlunoAprovado(solicitacao.IDUsuario, solicitacao.IDAcademia)
	Repository.Inserir(&aluno)
	Repository.Delete[entities.SolicitacoesConvite]("id_solicitacao", id_solicitacao)
	return nil
}

func ListarAlunos(id_usuario string) ([]entities.Alunos, error) {
	professor := Repository.SelectWhere[entities.Professores]("id_usuario_professor", id_usuario)
	if professor.IDProfessor == "" {
		return nil, errors.New("voce nao pode visualizar alunos ja que nao eh um professor")
	}

	return Repository.SelectStudentsByGym(professor.IDAcademiaProfessor), nil
}

func RemoverAluno(id_usuario string, id_aluno string) error {
	professor := Repository.SelectWhere[entities.Professores]("id_usuario_professor", id_usuario)
	if professor.IDProfessor == "" {
		return errors.New("voce nao pode remover alunos ja que nao eh um professor")
	}

	aluno := Repository.SelectWhere[entities.Alunos]("id_aluno", id_aluno)
	if aluno.IDAluno == "" {
		return errors.New("aluno nao encontrado")
	}

	if aluno.IDAcademiaAluno != professor.IDAcademiaProfessor {
		return errors.New("esse aluno nao pertence a sua academia")
	}

	Repository.Delete[entities.Alunos]("id_aluno", id_aluno)
	return nil
}

func CriarAula(id_usuario string, conteudo string, data_aula string) error {
	id_academia := ""

	professor := Repository.SelectWhere[entities.Professores]("id_usuario_professor", id_usuario)
	if professor.IDProfessor != "" {
		id_academia = professor.IDAcademiaProfessor
	}

	if id_academia == "" {
		instrutor := Repository.SelectWhere[entities.Instrutores]("id_usuario_instrutor", id_usuario)
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

func RegistrarPresenca(id_usuario string, id_aula string, latitude float64, longitude float64) error {
	aluno := Repository.SelectWhere[entities.Alunos]("id_usuario_aluno", id_usuario)
	if aluno.IDAluno == "" {
		return errors.New("voce nao eh aluno de nenhuma academia")
	}

	aula := Repository.SelectWhere[entities.Aulas]("id_aula", id_aula)
	if aula.IDAula == "" {
		return errors.New("aula nao encontrada")
	}

	if aula.IDAcademia != aluno.IDAcademiaAluno {
		return errors.New("essa aula nao eh da sua academia")
	}

	if Repository.ExistsTwo[entities.Presencas]("id_aluno", aluno.IDAluno, "id_aula", id_aula) {
		return errors.New("presenca ja registrada para essa aula")
	}

	academia := Repository.SelectWhere[entities.Academias]("ID", aluno.IDAcademiaAluno)
	if academia.Latitude == nil || academia.Longitude == nil {
		return errors.New("localizacao da academia ainda nao foi cadastrada pelo professor")
	}

	distancia := InterfaceAdapters.CalcularDistanciaMetros(latitude, longitude, *academia.Latitude, *academia.Longitude)
	if distancia >= 100 {
		return errors.New("voce esta longe demais da academia para registrar presenca")
	}

	presenca := InterfaceAdapters.MapearPresenca(aluno.IDAluno, id_aula)
	Repository.Inserir(presenca)
	return nil
}

func RegistrarLocalizacaoAcademia(id_usuario string, latitude float64, longitude float64) error {
	professor := Repository.SelectWhere[entities.Professores]("id_usuario_professor", id_usuario)
	if professor.IDProfessor == "" {
		return errors.New("voce nao eh professor de nenhuma academia")
	}

	Repository.UpdateLocationGym(professor.IDAcademiaProfessor, latitude, longitude)
	return nil
}

func ContarPresencasAluno(id_aluno string) int64 {
	return Repository.Count[entities.Presencas]("id_aluno", id_aluno)
}

func CriarInstrutor(id_usuario string, id_aluno string) error {
	professor := Repository.SelectWhere[entities.Professores]("id_usuario_professor", id_usuario)
	if professor.IDProfessor == "" {
		return errors.New("voce nao pode criar instrutores ja que nao eh um professor")
	}

	aluno := Repository.SelectWhere[entities.Alunos]("id_aluno", id_aluno)
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

func CriarProduto(id_usuario string, nome string, preco float64, tamanho string, quantidade int) error {
	professor := Repository.SelectWhere[entities.Professores]("id_usuario_professor", id_usuario)
	if professor.IDProfessor == "" {
		return errors.New("voce nao eh professor de nenhuma academia")
	}

	produto := InterfaceAdapters.MapearProduto(nome, preco, tamanho, quantidade, professor.IDAcademiaProfessor)
	Repository.Inserir(produto)
	return nil
}

func AtualizarProduto(id_usuario string, id_produto string, nome string, preco float64, tamanho string, quantidade int) error {
	professor := Repository.SelectWhere[entities.Professores]("id_usuario_professor", id_usuario)
	if professor.IDProfessor == "" {
		return errors.New("voce nao eh professor de nenhuma academia")
	}

	produto := Repository.SelectWhere[entities.Produtos]("id_produto", id_produto)
	if produto.IDProduto == "" {
		return errors.New("produto nao encontrado")
	}

	if produto.IDAcademia != professor.IDAcademiaProfessor {
		return errors.New("esse produto nao pertence a sua academia")
	}

	Repository.UpdateProduct(id_produto, nome, preco, tamanho, quantidade)
	return nil
}

func DeletarProduto(id_usuario string, id_produto string) error {
	professor := Repository.SelectWhere[entities.Professores]("id_usuario_professor", id_usuario)
	if professor.IDProfessor == "" {
		return errors.New("voce nao eh professor de nenhuma academia")
	}

	produto := Repository.SelectWhere[entities.Produtos]("id_produto", id_produto)
	if produto.IDProduto == "" {
		return errors.New("produto nao encontrado")
	}

	if produto.IDAcademia != professor.IDAcademiaProfessor {
		return errors.New("esse produto nao pertence a sua academia")
	}

	Repository.Delete[entities.Produtos]("id_produto", id_produto)
	return nil
}

func ListarCatalogo(id_usuario string) ([]entities.Produtos, error) {
	id_academia := ""

	professor := Repository.SelectWhere[entities.Professores]("id_usuario_professor", id_usuario)
	if professor.IDProfessor != "" {
		id_academia = professor.IDAcademiaProfessor
	}

	if id_academia == "" {
		aluno := Repository.SelectWhere[entities.Alunos]("id_usuario_aluno", id_usuario)
		if aluno.IDAluno != "" {
			id_academia = aluno.IDAcademiaAluno
		}
	}

	if id_academia == "" {
		return nil, errors.New("voce nao pertence a nenhuma academia")
	}

	return Repository.SelectWhereList[entities.Produtos]("id_academia", id_academia), nil
}

func SinalizarInteresse(id_usuario string, id_produto string, quantidade int) error {
	aluno := Repository.SelectWhere[entities.Alunos]("id_usuario_aluno", id_usuario)
	if aluno.IDAluno == "" {
		return errors.New("voce nao eh aluno de nenhuma academia")
	}

	produto := Repository.SelectWhere[entities.Produtos]("id_produto", id_produto)
	if produto.IDProduto == "" {
		return errors.New("produto nao encontrado")
	}

	if produto.IDAcademia != aluno.IDAcademiaAluno {
		return errors.New("esse produto nao pertence a sua academia")
	}

	interesse := InterfaceAdapters.MapearInteresse(aluno.IDAluno, id_produto, quantidade)
	Repository.Inserir(interesse)
	return nil
}

func RealizarPagamento(id_usuario string, valor_centavos int64) (string, string, error) {
	aluno := Repository.SelectWhere[entities.Alunos]("id_usuario_aluno", id_usuario)
	if aluno.IDAluno == "" {
		return "", "", errors.New("voce nao eh aluno de nenhuma academia")
	}

	usuario := Repository.SelectWhere[entities.Usuarios]("ID", id_usuario)

	descricao := fmt.Sprintf("Mensalidade academia - aluno %s", usuario.Nome)
	resposta_stripe, err := InterfaceAdapters.StripeCriarPagamento(valor_centavos, descricao)
	if err != nil {
		return "", "", err
	}

	pagamento := InterfaceAdapters.MapearPagamento(aluno.IDAluno, resposta_stripe.ID, valor_centavos)
	Repository.Inserir(pagamento)

	return pagamento.IDPagamento, resposta_stripe.ClientSecret, nil
}

func ListarAulasDoDia(id_usuario string) ([]entities.Aulas, error) {
	id_academia := ""

	aluno := Repository.SelectWhere[entities.Alunos]("id_usuario_aluno", id_usuario)
	if aluno.IDAluno != "" {
		id_academia = aluno.IDAcademiaAluno
	}

	data := InterfaceAdapters.DateTimeNow()

	inicio := data
	fim := data.Add(3 * time.Hour)

	return Repository.SelectDayPresences(id_academia, inicio, fim), nil
}
