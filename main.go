package main

import (
	"github.com/joho/godotenv"
	"urlshortner/internal/bootstrap"
	"urlshortner/internal/database"
	"urlshortner/internal/router"
)

func init() {
	err := godotenv.Load("conf/.toml")
	if err != nil {
		panic(err)
	}
	database.ConnectDb()
	bootstrap.StartCronJob()
}

func main() {
	router.ClientRoutes()
}
