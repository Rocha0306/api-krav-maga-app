package main

import (
	"encoding/json"
	"net/http"
)

func DesserializeStudentDTO(request *http.Request) StudentDTO {
	var studentdto StudentDTO
	json.NewDecoder(request.Body).Decode(&studentdto)
	return studentdto
}
