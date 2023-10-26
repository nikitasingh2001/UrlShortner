package bootstrap

import (
	"github.com/robfig/cron/v3"
	"urlshortner/internal/database"
)

var Cron *cron.Cron

func StartCronJob() {
	Cron = cron.New()

	Cron.AddFunc("@every 1m", func() {
		database.DeleteExpiredURLs()
	})

	Cron.Start()
}
