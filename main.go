package main

import (
	"api-back-end/api/InterfaceAdapters"
	"api-back-end/api/Presentation"
	"net/http"
)

func main() {

	InterfaceAdapters.Migrations()
	http.HandleFunc("POST /Users/Auth", Presentation.ControllerLogin)
	http.HandleFunc("GET /Users/Me", Presentation.ControllerPerfilUsuario)
	http.HandleFunc("POST /Users/Registration", Presentation.ControllerCadastro)
	http.HandleFunc("POST /Users/Registration/Confirm", Presentation.ControllerCadastroConfirmar)
	http.HandleFunc("POST /Users/Password/Forgot", Presentation.ControllerEsqueciSenha)
	http.HandleFunc("POST /Users/Password/Reset", Presentation.ControllerRedefinirSenha)
	http.HandleFunc("POST /Gyms/Creation", Presentation.ControllerCriarAcademia)
	http.HandleFunc("POST /Gyms/Invites/Creation", Presentation.ControllerGerarConvites)
	http.HandleFunc("GET /Gyms/Invites/List", Presentation.ControllerMostrarConvites)
	http.HandleFunc("DELETE /Gyms/Invites/{id_convite}", Presentation.ControllerDeletarConvite)
	http.HandleFunc("POST /Gyms/Requests/Join", Presentation.ControllerSolicitarEntrada)
	http.HandleFunc("GET /Gyms/Requests/Join/List", Presentation.ControllerListarSolicitacoes)
	http.HandleFunc("POST /Gyms/Requests/Join/Approve", Presentation.ControllerAprovarSolicitacao)
	http.HandleFunc("GET /Gyms/Students/List", Presentation.ControllerListarAlunos)
	http.HandleFunc("DELETE /Gyms/Students/{id_aluno}", Presentation.ControllerRemoverAluno)
	http.HandleFunc("PUT /Gyms/Students/{id_aluno}/Belt", Presentation.ControllerAtualizarFaixaAluno)
	http.HandleFunc("POST /Gyms/Instructors/Creation", Presentation.ControllerCriarInstrutor)
	http.HandleFunc("POST /Gyms/Classes/Creation", Presentation.ControllerCriarAula)
	http.HandleFunc("PUT /Gyms/Classes/Update", Presentation.ControllerAtualizarAula)
	http.HandleFunc("GET /Gyms/Classes/Day", Presentation.ControllerListarAulasDoDia)
	http.HandleFunc("GET /Gyms/Classes", Presentation.ControllerListarAulasProfessor)
	http.HandleFunc("POST /Student/Presence", Presentation.ControllerRegistrarPresenca)
	http.HandleFunc("GET /Student/Presence/Count", Presentation.ControllerContarPresencasAluno)
	http.HandleFunc("PUT /Gyms/Location", Presentation.ControllerRegistrarLocalizacaoAcademia)
	http.HandleFunc("GET /Gyms/Geolocation", Presentation.ControllerLocalizacaoAcademia)
	http.HandleFunc("POST /Gyms/Catalog/Creation", Presentation.ControllerCriarProduto)
	http.HandleFunc("PUT /Gyms/Catalog/Update", Presentation.ControllerAtualizarProduto)
	http.HandleFunc("DELETE /Gyms/Catalog/{id_produto}", Presentation.ControllerDeletarProduto)
	http.HandleFunc("GET /Gyms/Catalog", Presentation.ControllerListarCatalogo)
	http.HandleFunc("POST /Student/Interest", Presentation.ControllerSinalizarInteresse)
	http.HandleFunc("POST /Student/Payment", Presentation.ControllerRealizarPagamento)
	http.ListenAndServe(":8080", cors(http.DefaultServeMux))

}

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}
