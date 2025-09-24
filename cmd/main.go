package main

import (
	"fmt"
	"net/http"

	"github.com/Dionisius951/Golang-Blog-BE/internal/config"
	"github.com/Dionisius951/Golang-Blog-BE/internal/db"
	"github.com/Dionisius951/Golang-Blog-BE/internal/routes"
)

func main() {
	PORT := config.LoadConfig().APP_PORT
	db.ConnecDB()

	r := routes.SetupRouter()

	fmt.Printf("SERVER RUNNING ON PORT %v \n", PORT)
	http.ListenAndServe(PORT, r)
}
