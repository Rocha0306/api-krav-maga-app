package Presentation

import (
	"encoding/json"
	"net/http"
)

func DesserializeStudentDTO(request *http.Request) UserDTO {
	var studentdto UserDTO
	json.NewDecoder(request.Body).Decode(&studentdto)
	return studentdto

}

func SerializeErrorMessageResponse(error_message string, response http.ResponseWriter) {
	json.NewEncoder(response).Encode(error_message)
}

func SerializeMessageResponse(message string, response http.ResponseWriter) {
	json.NewEncoder(response).Encode(message)
}
