package main

import (
	"api-back-end/api/Presentation"
	"net/http"
)

func main() {

	http.HandleFunc("/Users/Auth", Presentation.ControllerLogin)
	http.HandleFunc("/Users/Registration", Presentation.ControllerRegistration)
	http.HandleFunc("/Users/Registration/Confirm", Presentation.ControllerRegistrationConfirm)
	http.HandleFunc("/Gyms/Creation", Presentation.ControllerGymCreation)
	http.HandleFunc("/Gyms/Invites/Creation", Presentation.ControllerGenerateInvites)
	http.HandleFunc("/Gyms/Invites/List", Presentation.ControllerShowInvites)
	http.HandleFunc("/Gyms/Requests/Join", Presentation.ControllerGymsRequestsJoin)
	http.HandleFunc("/Gyms/Requests/Join/List", Presentation.ControllerListRequestsJoin)
	http.ListenAndServe(":8080", nil)

}
