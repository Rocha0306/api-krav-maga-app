package UsersCase

import (
	entities "api-back-end/api/Entities"
	"api-back-end/api/InterfaceAdapters"
	"api-back-end/api/Repository"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"
)

/*
1 - Pega o Response da API de CPF e CEP
2 - Popula o struct que e a entidade no banco
3 - Insere no banco o Cep e coloca o Usuario em cache
*/
func PrepararUsuario(cpf string, cep string, email string, senha string) error {

	if Repository.Exists[entities.Usuarios]("email", email) {
		return errors.New("Usuario ja cadastrado")
	}

	var wg sync.WaitGroup
	var err_cpf, err_cep error
	var resposta_api_cpf InterfaceAdapters.CpfApiResponse
	var resposta_api_cep InterfaceAdapters.CepApiResponse

	wg.Add(2)
	go func() {
		defer wg.Done()
		err_cpf, resposta_api_cpf = InterfaceAdapters.CpfApi(cpf)
	}()
	go func() {
		defer wg.Done()
		err_cep, resposta_api_cep = InterfaceAdapters.CepApi(cep)
	}()
	wg.Wait()

	if err_cpf != nil {
		return err_cpf
	}
	if err_cep != nil {
		return err_cep
	}

	entidade_final := InterfaceAdapters.MapearUsuario(email, senha, resposta_api_cpf, resposta_api_cep)

	if Repository.Exists[entities.Endereco]("cep", entidade_final.Endereco.CEP) {
		endereco_existente := Repository.SelectWhere[entities.Endereco]("cep", entidade_final.Endereco.CEP)
		entidade_final.EnderecoID = endereco_existente.ID
		entidade_final.Endereco = *endereco_existente
	}

	numero_auth := strconv.Itoa(InterfaceAdapters.GerarNumeroAuth())
	InterfaceAdapters.SalvarUsuarioCache(numero_auth, *entidade_final)
	err := InterfaceAdapters.EnviarEmail(fmt.Sprintf("Codigo de Autenticacao: %s", numero_auth), []string{entidade_final.Email})

	if err != nil {
		InterfaceAdapters.EscreverLogsMongoDb(err.Error(), "Falha ao Enviar Email")
		InterfaceAdapters.LogFatal(err.Error())
	}

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
	usuario := Repository.SelectWhere[entities.Usuarios]("email", email)

	if usuario.Email != email || !InterfaceAdapters.ComparaSenhaHash(usuario.Senha, senha) {
		return "", errors.New("Aluno nao existe no sistema")
	}

	return InterfaceAdapters.GerarTokenJwt(usuario.ID), nil
}

// AcademiaVinculo descreve uma academia a que o usuario pertence e em qual papel.
// Um mesmo usuario pode ser professor de uma academia e aluno de outra, por isso
// o perfil retorna uma lista.
type AcademiaVinculo struct {
	ID      string
	Nome    string
	CNPJ    string
	Vinculo string // "professor" | "instrutor" | "aluno"
	Faixa   string // preenchido apenas quando o vinculo e "aluno"
}

func PerfilUsuario(id_usuario string) (entities.Usuarios, string, string, []AcademiaVinculo) {
	usuario := Repository.SelectWhere[entities.Usuarios]("ID", id_usuario)

	academias := []AcademiaVinculo{}

	professores := Repository.SelectWhereList[entities.Professores]("id_usuario_professor", id_usuario)
	for _, p := range professores {
		ac := Repository.SelectWhere[entities.Academias]("ID", p.IDAcademiaProfessor)
		academias = append(academias, AcademiaVinculo{ID: ac.ID, Nome: ac.Nome, CNPJ: ac.CNPJ, Vinculo: "professor"})
	}

	instrutores := Repository.SelectWhereList[entities.Instrutores]("id_usuario_instrutor", id_usuario)
	for _, i := range instrutores {
		ac := Repository.SelectWhere[entities.Academias]("ID", i.IDAcademiaInstrutor)
		academias = append(academias, AcademiaVinculo{ID: ac.ID, Nome: ac.Nome, CNPJ: ac.CNPJ, Vinculo: "instrutor"})
	}

	alunos := Repository.SelectWhereList[entities.Alunos]("id_usuario_aluno", id_usuario)
	for _, a := range alunos {
		ac := Repository.SelectWhere[entities.Academias]("ID", a.IDAcademiaAluno)
		academias = append(academias, AcademiaVinculo{ID: ac.ID, Nome: ac.Nome, CNPJ: ac.CNPJ, Vinculo: "aluno", Faixa: a.Faixa})
	}

	role := "student"
	if len(professores) > 0 {
		role = "teacher"
	}

	faixa := ""
	if len(alunos) > 0 {
		faixa = alunos[0].Faixa
	}

	return *usuario, role, faixa, academias
}

func EsqueciSenha(email string) error {
	usuario := Repository.SelectWhere[entities.Usuarios]("email", email)
	if usuario.ID == "" {
		return nil
	}

	if InterfaceAdapters.ThrottleAtingido("reset-throttle:"+email, 2*time.Minute) {
		return nil
	}

	codigo := strconv.Itoa(InterfaceAdapters.GerarNumeroAuth())
	InterfaceAdapters.SalvarValorCache("reset:"+codigo, usuario.ID)
	InterfaceAdapters.EnviarEmail(fmt.Sprintf("Codigo para redefinir sua senha: %s", codigo), []string{usuario.Email})
	return nil
}

func RedefinirSenha(codigo string, nova_senha string) error {
	id_usuario, err := InterfaceAdapters.PegarValorCache("reset:" + codigo)
	if err != nil {
		return errors.New("codigo invalido ou expirado")
	}

	Repository.UpdateSenhaUsuario(id_usuario, InterfaceAdapters.HashSenha(nova_senha))
	InterfaceAdapters.RemoverValorCache("reset:" + codigo)
	return nil
}

func CriarAcademia(cnpj string, nome_academia string, id_usuario string) error {
	entidade_academia := InterfaceAdapters.MapearAcademia(cnpj, nome_academia)
	Repository.Inserir(entidade_academia)
	entidade_professor := InterfaceAdapters.MapearProfessor(id_usuario, entidade_academia.ID)
	Repository.Inserir(&entidade_professor)
	return nil
}

func GerarConvite(id_usuario string) (entities.Convites, error) {
	if !Repository.Exists[entities.Professores]("id_usuario_professor", id_usuario) {
		return entities.Convites{}, errors.New("Voce nao pode criar convites ja que nao eh professor")
	}

	professor := Repository.SelectWhere[entities.Professores]("id_usuario_professor", id_usuario)

	if Repository.Count[entities.Convites]("id_academia", professor.IDAcademiaProfessor) >= 10 {
		return entities.Convites{}, errors.New("limite de 10 convites por academia atingido, remova algum para criar outro")
	}

	entidade_convite := InterfaceAdapters.MapearConvite(professor.IDAcademiaProfessor)
	Repository.Inserir(entidade_convite)
	return *entidade_convite, nil
}

func MostrarConvites(id_usuario string) ([]entities.Convites, error) {
	if !Repository.Exists[entities.Professores]("id_usuario_professor", id_usuario) {
		return nil, errors.New("Voce nao pode visualizar convites ja que nao eh um professor")
	}

	professor := Repository.SelectWhere[entities.Professores]("id_usuario_professor", id_usuario)
	return Repository.SelectWhereList[entities.Convites]("id_academia", professor.IDAcademiaProfessor), nil
}

func DeletarConvite(id_usuario string, id_convite string) error {
	professor := Repository.SelectWhere[entities.Professores]("id_usuario_professor", id_usuario)
	if professor.IDProfessor == "" {
		return errors.New("voce nao eh professor de nenhuma academia")
	}

	convite := Repository.SelectWhere[entities.Convites]("id_convite", id_convite)
	if convite.IDConvite == "" {
		return errors.New("convite nao encontrado")
	}

	if convite.IDAcademia != professor.IDAcademiaProfessor {
		return errors.New("esse convite nao pertence a sua academia")
	}

	Repository.Delete[entities.Convites]("id_convite", id_convite)
	return nil
}

func SolicitarEntrada(codigo_convite string, id_usuario string) error {
	if !Repository.Exists[entities.Convites]("chave_convite", codigo_convite) {
		return errors.New("convite nao encontrado")
	}

	convite := Repository.SelectWhere[entities.Convites]("chave_convite", codigo_convite)

	if Repository.ExistsTwo[entities.Professores]("id_usuario_professor", id_usuario, "id_academia_professor", convite.IDAcademia) {
		return errors.New("voce ja eh professor dessa academia")
	}

	if Repository.ExistsTwo[entities.Alunos]("id_usuario_aluno", id_usuario, "id_academia_aluno", convite.IDAcademia) {
		return errors.New("voce ja eh aluno dessa academia")
	}

	if Repository.ExistsTwo[entities.SolicitacoesConvite]("id_usuario", id_usuario, "id_academia", convite.IDAcademia) {
		return errors.New("voce ja solicitou entrada nessa academia, aguarde a aprovacao do professor")
	}

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

func CriarAula(id_usuario string, conteudo string, data_aula string, faixa string) error {
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

	entidade_aula := InterfaceAdapters.MapearAula(conteudo, id_academia, id_usuario, data, faixa)
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

	if aula.Faixa != aluno.Faixa {
		return errors.New("essa aula nao eh para a sua graduacao")
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

func LocalizacaoAcademia(id_usuario string) (float64, float64, error) {
	id_academia := ""

	aluno := Repository.SelectWhere[entities.Alunos]("id_usuario_aluno", id_usuario)
	if aluno.IDAluno != "" {
		id_academia = aluno.IDAcademiaAluno
	}

	if id_academia == "" {
		professor := Repository.SelectWhere[entities.Professores]("id_usuario_professor", id_usuario)
		if professor.IDProfessor != "" {
			id_academia = professor.IDAcademiaProfessor
		}
	}

	if id_academia == "" {
		return 0, 0, errors.New("voce nao pertence a nenhuma academia")
	}

	academia := Repository.SelectWhere[entities.Academias]("ID", id_academia)
	if academia.Latitude == nil || academia.Longitude == nil {
		return 0, 0, errors.New("localizacao da academia ainda nao foi cadastrada")
	}

	return *academia.Latitude, *academia.Longitude, nil
}

func ContarPresencasAluno(id_usuario string, id_aluno string) ([]entities.Presencas, error) {
	professor := Repository.SelectWhere[entities.Professores]("id_usuario_professor", id_usuario)
	if professor.IDProfessor == "" {
		return nil, errors.New("voce nao pode ver presencas ja que nao eh um professor")
	}

	aluno := Repository.SelectWhere[entities.Alunos]("id_aluno", id_aluno)
	if aluno.IDAluno == "" {
		return nil, errors.New("aluno nao encontrado")
	}

	if aluno.IDAcademiaAluno != professor.IDAcademiaProfessor {
		return nil, errors.New("esse aluno nao pertence a sua academia")
	}

	return Repository.SelectPresencasComAula(id_aluno), nil
}

func AtualizarFaixaAluno(id_usuario string, id_aluno string, faixa string) error {
	professor := Repository.SelectWhere[entities.Professores]("id_usuario_professor", id_usuario)
	if professor.IDProfessor == "" {
		return errors.New("voce nao pode atualizar faixa ja que nao eh um professor")
	}

	aluno := Repository.SelectWhere[entities.Alunos]("id_aluno", id_aluno)
	if aluno.IDAluno == "" {
		return errors.New("aluno nao encontrado")
	}

	if aluno.IDAcademiaAluno != professor.IDAcademiaProfessor {
		return errors.New("esse aluno nao pertence a sua academia")
	}

	Repository.UpdateAlunoFaixa(id_aluno, faixa)
	return nil
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

func CriarProduto(id_usuario string, nome string, preco float64, tamanho string, quantidade int, imagem_url string) error {
	professor := Repository.SelectWhere[entities.Professores]("id_usuario_professor", id_usuario)
	if professor.IDProfessor == "" {
		return errors.New("voce nao eh professor de nenhuma academia")
	}

	produto := InterfaceAdapters.MapearProduto(nome, preco, tamanho, quantidade, imagem_url, professor.IDAcademiaProfessor)
	Repository.Inserir(produto)
	return nil
}

func AtualizarProduto(id_usuario string, id_produto string, nome string, preco float64, tamanho string, quantidade int, imagem_url string) error {
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

	Repository.UpdateProduct(id_produto, nome, preco, tamanho, quantidade, imagem_url)
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

	if id_academia == "" {
		professor := Repository.SelectWhere[entities.Professores]("id_usuario_professor", id_usuario)
		if professor.IDProfessor != "" {
			id_academia = professor.IDAcademiaProfessor
		}
	}

	agora := InterfaceAdapters.DateTimeNow()
	inicio := time.Date(agora.Year(), agora.Month(), agora.Day(), 0, 0, 0, 0, agora.Location())
	fim := inicio.Add(24 * time.Hour)

	return Repository.SelectDayPresences(id_academia, inicio, fim), nil
}

func ListarAulasProfessor(id_usuario string, data_aula string) ([]entities.Aulas, error) {
	professor := Repository.SelectWhere[entities.Professores]("id_usuario_professor", id_usuario)
	if professor.IDProfessor == "" {
		return nil, errors.New("voce nao eh professor de nenhuma academia")
	}

	if data_aula == "" {
		return Repository.SelectWhereList[entities.Aulas]("id_academia", professor.IDAcademiaProfessor), nil
	}

	data, err := time.Parse("2006-01-02", data_aula)
	if err != nil {
		return nil, errors.New("data invalida, use o formato AAAA-MM-DD")
	}

	inicio := data
	fim := data.Add(24 * time.Hour)
	return Repository.SelectDayPresences(professor.IDAcademiaProfessor, inicio, fim), nil
}
