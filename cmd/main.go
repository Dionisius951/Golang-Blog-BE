package main

import (
	"fmt"
	"net/http"

	"github.com/Dionisius951/Golang-Blog-BE/internal/config"
	"github.com/gorilla/mux"
)

func main() {
	PORT := config.LoadConfig().APP_PORT

	r := mux.NewRouter()

	fmt.Printf("SERVER RUNNING ON PORT %v", PORT)
	http.ListenAndServe(PORT, r)
}
