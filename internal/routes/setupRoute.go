package routes

import "github.com/gorilla/mux"

func SetupRouter() *mux.Router {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()

	// all routes here
	postRouter(api)
	userRouter(api)
	return r
}
