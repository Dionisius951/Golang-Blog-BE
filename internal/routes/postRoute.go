package routes

import (
	"github.com/Dionisius951/Golang-Blog-BE/internal/controllers"
	"github.com/gorilla/mux"
)

func postRouter(r *mux.Router) {
	postRoute := r.PathPrefix("/content").Subrouter()

	postRoute.HandleFunc("", controllers.AddPostHandler).Methods("POST")
	postRoute.HandleFunc("", controllers.GetAllPostHandler).Methods("GET")
	postRoute.HandleFunc("/{id}", controllers.GetPostByIdHandler).Methods("GET")
	postRoute.HandleFunc("/{id}", controllers.UpdatePostHandler).Methods("PUT")
	postRoute.HandleFunc("/{id}", controllers.DeletePostHandler).Methods("DELETE")
}
