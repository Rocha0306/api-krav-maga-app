package Presentation

import (
	"encoding/json"
	"net/http"
)

func Deserialize[T any](request *http.Request) T {
	var dto T
	json.NewDecoder(request.Body).Decode(&dto)
	return dto
}

func SerializeErrorMessageResponse(error_message string, response http.ResponseWriter) {
	json.NewEncoder(response).Encode(error_message)
}

func SerializeMessageResponse(message any, response http.ResponseWriter) {
	json.NewEncoder(response).Encode(message)
}
