package Presentation

import "net/http"

func StatusCode200(response http.ResponseWriter, message any) {
	response.WriteHeader(200)
	SerializeMessageResponse(message, response)
}

func BadRequest(response http.ResponseWriter, err error) {
	response.WriteHeader(400)
	SerializeErrorMessageResponse(err.Error(), response)
}
