package bootstrap

import (
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
	"urlshortner/internal/database"
	"urlshortner/internal/router"
)

func Initalize() {
	err := godotenv.Load("conf/.toml")
	if err != nil {
		panic(err)
	}
	database.ConnectDb()
	StartCronJob()
	router.ClientRoutes()
}

var Cron *cron.Cron

func StartCronJob() {
	Cron = cron.New()

	Cron.AddFunc("@every 1m", func() {
		database.DeleteExpiredURLs()
	})

	Cron.Start()
}
