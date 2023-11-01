package bootstrap

import (
	"context"
	"github.com/robfig/cron/v3"
	"urlshortner/internal/config"
	"urlshortner/internal/database"
	"urlshortner/internal/router"
)

var loadConfig = config.LoadConfig

func Initialize(ctx context.Context, basePath string, env string) {

	baseInit(ctx, basePath, env)

	database.ConnectDb()
	StartCronJob()
	router.ClientRoutes()
}

func baseInit(ctx context.Context, basePath string, env string) {
	loadConfig(basePath, env)
}

var Cron *cron.Cron

func StartCronJob() {
	Cron = cron.New()

	Cron.AddFunc("@every 1m", func() {
		database.DeleteExpiredURLs()
	})

	Cron.Start()
}
