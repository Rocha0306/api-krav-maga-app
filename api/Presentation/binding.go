package Presentation

import (
	"encoding/json"
	"io"
	"net/http"
)

func DesserializeUserDTO(request *http.Request) UserDTO {
	var studentdto UserDTO
	json.NewDecoder(request.Body).Decode(&studentdto)
	return studentdto

}

func DesserializeLoginDTO(request *http.Request) LoginDTO {
	var studentdto LoginDTO
	json.NewDecoder(request.Body).Decode(&studentdto)
	return studentdto

}

func DesserializeGymDTO(body io.ReadCloser) GymDTO {
	var gym_dto GymDTO
	json.NewDecoder(body).Decode(&gym_dto)
	return gym_dto
}

func SerializeErrorMessageResponse(error_message string, response http.ResponseWriter) {
	json.NewEncoder(response).Encode(error_message)
}

func SerializeMessageResponse(message string, response http.ResponseWriter) {
	json.NewEncoder(response).Encode(message)
}
