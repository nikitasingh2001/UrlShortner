package main

import (
	"urlshortner/internal/database"
	"urlshortner/internal/router"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("config/.env")
	if err != nil {
		panic(err)
	}
	database.ConnectDb()
}

func main() {
	router.ClientRoutes()
}
