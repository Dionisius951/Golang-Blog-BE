package routes

import (
	"github.com/Dionisius951/Golang-Blog-BE/internal/controllers"
	"github.com/Dionisius951/Golang-Blog-BE/internal/middleware"
	"github.com/gorilla/mux"
)

func postRouter(r *mux.Router) {
	publicRoute := r.PathPrefix("/content").Subrouter()
	publicRoute.HandleFunc("", controllers.GetAllPostHandler).Methods("GET")

	postRoute := r.PathPrefix("/content").Subrouter()
	postRoute.Use(middleware.LoggingMiddleware)

	postRoute.HandleFunc("", controllers.AddPostHandler).Methods("POST")
	postRoute.HandleFunc("/{id}", controllers.GetPostByIdHandler).Methods("GET")
	postRoute.HandleFunc("/{id}", controllers.UpdatePostHandler).Methods("PUT")
	postRoute.HandleFunc("/{id}", controllers.DeletePostHandler).Methods("DELETE")
}
