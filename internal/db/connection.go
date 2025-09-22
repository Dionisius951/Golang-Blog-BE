package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Dionisius951/Golang-Blog-BE/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func ConnecDB() {
	DB_URL := config.LoadConfig().DB_URL
	var err error

	Pool, err = pgxpool.New(context.Background(), DB_URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	err = Pool.Ping(context.Background())
	if err != nil {
		fmt.Println("Error While Trying Connect DB")
	}

	log.Print("Database Connect Successfully")
}

func CloseDB(){
	if Pool != nil {
		Pool.Close()
		log.Print("Database Connection Closed")
	}
}