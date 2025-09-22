package main

import (
	"fmt"
	"net/http"

	"github.com/Dionisius951/Golang-Blog-BE/internal/config"
	"github.com/Dionisius951/Golang-Blog-BE/internal/db"
	"github.com/gorilla/mux"
)

func main() {
	PORT := config.LoadConfig().APP_PORT
	db.ConnecDB()

	r := mux.NewRouter()

	fmt.Printf("SERVER RUNNING ON PORT %v", PORT)
	http.ListenAndServe(PORT, r)
}
