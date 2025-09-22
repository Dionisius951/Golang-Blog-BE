package routes

import "github.com/gorilla/mux"

func SetupRouter() *mux.Router{
	r := mux.NewRouter()

	// all routes here
	userRouter(r)
	return r
}