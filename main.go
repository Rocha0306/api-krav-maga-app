package main

import (
	"api-back-end/api/InterfaceAdapters"
	"api-back-end/api/Presentation"
	"net/http"
)

func main() {

	InterfaceAdapters.Migrations()
	http.HandleFunc("/Users/Auth", Presentation.ControllerLogin)
	http.HandleFunc("/Users/Registration", Presentation.ControllerCadastro)
	http.HandleFunc("/Users/Registration/Confirm", Presentation.ControllerCadastroConfirmar)
	http.HandleFunc("/Gyms/Creation", Presentation.ControllerCriarAcademia)
	http.HandleFunc("/Gyms/Invites/Creation", Presentation.ControllerGerarConvites)
	http.HandleFunc("/Gyms/Invites/List", Presentation.ControllerMostrarConvites)
	http.HandleFunc("/Gyms/Requests/Join", Presentation.ControllerSolicitarEntrada)
	http.HandleFunc("/Gyms/Requests/Join/List", Presentation.ControllerListarSolicitacoes)
	http.HandleFunc("/Gyms/Requests/Join/Approve", Presentation.ControllerAprovarSolicitacao)
	http.HandleFunc("/Gyms/Students/List", Presentation.ControllerListarAlunos)
	http.HandleFunc("/Gyms/Students/Remove", Presentation.ControllerRemoverAluno)
	http.HandleFunc("/Gyms/Instructors/Creation", Presentation.ControllerCriarInstrutor)
	http.HandleFunc("/Gyms/Classes/Creation", Presentation.ControllerCriarAula)
	http.HandleFunc("/Gyms/Classes/Day", Presentation.ControllerListarAulasDoDia)
	http.HandleFunc("/Student/Presence", Presentation.ControllerRegistrarPresenca)
	http.HandleFunc("/Student/Presence/Count", Presentation.ControllerContarPresencasAluno)
	http.HandleFunc("/Gyms/Location", Presentation.ControllerRegistrarLocalizacaoAcademia)
	http.HandleFunc("/Gyms/Catalog/Creation", Presentation.ControllerCriarProduto)
	http.HandleFunc("/Gyms/Catalog/Update", Presentation.ControllerAtualizarProduto)
	http.HandleFunc("/Gyms/Catalog/Delete", Presentation.ControllerDeletarProduto)
	http.HandleFunc("/Gyms/Catalog", Presentation.ControllerListarCatalogo)
	http.HandleFunc("/Student/Interest", Presentation.ControllerSinalizarInteresse)
	http.HandleFunc("/Student/Payment", Presentation.ControllerRealizarPagamento)
	http.ListenAndServe(":8080", nil)

}
