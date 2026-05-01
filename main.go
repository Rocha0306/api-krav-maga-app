package main

import (
	"api-back-end/api/Presentation"
	"net/http"
)

func main() {

	http.HandleFunc("/Users/Auth", Presentation.ControllerLogin)
	http.HandleFunc("/Users/Registration", Presentation.ControllerRegistration)
	http.HandleFunc("/Gyms/Creation", Presentation.ControllerGymCreation)
	http.ListenAndServe(":8080", nil)

}
