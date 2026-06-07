package Presentation

import "net/http"

func Status200(response http.ResponseWriter, mensagem any) {
	response.WriteHeader(200)
	SerializarRespostaMensagem(mensagem, response)
}

func BadRequest(response http.ResponseWriter, err error) {
	response.WriteHeader(400)
	SerializarRespostaErro(err.Error(), response)
}
