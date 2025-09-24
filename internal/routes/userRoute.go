package routes

import (
	"github.com/Dionisius951/Golang-Blog-BE/internal/controllers"
	"github.com/gorilla/mux"
)

func userRouter(r *mux.Router) {
	userRoute := r.PathPrefix("/users").Subrouter()

	userRoute.HandleFunc("/register",controllers.RegisterHandler).Methods("POST")
}
