package routes

import (
	"github.com/Dionisius951/Golang-Blog-BE/internal/controllers"
	"github.com/gorilla/mux"
)

func userRouter(r *mux.Router) {
	userRoute := r.PathPrefix("/users").Subrouter()

	userRoute.HandleFunc("/login",controllers.LoginHandler).Methods("GET")
}
