package Presentation

import (
	"encoding/json"
	"net/http"
)

func Desserializar[T any](request *http.Request) T {
	var dto T
	json.NewDecoder(request.Body).Decode(&dto)
	return dto
}

func SerializarRespostaErro(mensagem_erro string, response http.ResponseWriter) {
	json.NewEncoder(response).Encode(mensagem_erro)
}

func SerializarRespostaMensagem(mensagem any, response http.ResponseWriter) {
	json.NewEncoder(response).Encode(mensagem)
}
