package routes

import (
	"github.com/Dionisius951/Golang-Blog-BE/internal/controllers"
	"github.com/gorilla/mux"
)

func commentRouter(r *mux.Router) {
	commentRoute := r.PathPrefix("/comment").Subrouter()

	commentRoute.HandleFunc("", controllers.AddCommentHandler).Methods("POST")
}
