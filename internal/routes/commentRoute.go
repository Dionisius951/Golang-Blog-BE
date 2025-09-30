package routes

import (
	"github.com/Dionisius951/Golang-Blog-BE/internal/controllers"
	"github.com/gorilla/mux"
)

func commentRouter(r *mux.Router) {
	commentRoute := r.PathPrefix("/comment").Subrouter()

	commentRoute.HandleFunc("", controllers.AddCommentHandler).Methods("POST")
	commentRoute.HandleFunc("/{id}", controllers.GetCommentByIdHandler).Methods("GET")
	commentRoute.HandleFunc("/{id}", controllers.UpdateCommentHandler).Methods("PUT")
	commentRoute.HandleFunc("/{id}", controllers.DeleteCommentHandler).Methods("DELETE")
}
