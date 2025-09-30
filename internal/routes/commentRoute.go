package routes

import (
	"github.com/Dionisius951/Golang-Blog-BE/internal/controllers"
	"github.com/Dionisius951/Golang-Blog-BE/internal/middleware"
	"github.com/gorilla/mux"
)

func commentRouter(r *mux.Router) {
	PublicRoute := r.PathPrefix("/comment").Subrouter()
	PublicRoute.HandleFunc("/{id}", controllers.GetCommentByIdHandler).Methods("GET")

	commentRoute := r.PathPrefix("/comment").Subrouter()
	commentRoute.Use(middleware.LoggingMiddleware)
	commentRoute.HandleFunc("", controllers.AddCommentHandler).Methods("POST")
	commentRoute.HandleFunc("/{id}", controllers.UpdateCommentHandler).Methods("PUT")
	commentRoute.HandleFunc("/{id}", controllers.DeleteCommentHandler).Methods("DELETE")
}
