package database

import (
	"context"
	"fmt"
	"time"
	"urlshortner/internal/config"
	"urlshortner/internal/constant"
	"urlshortner/internal/logger"
	"urlshortner/internal/models"
	"urlshortner/internal/repository"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Mgr Manager

type Manager interface {
	Insert(interface{}, string) (interface{}, error)
	GetUrlFromCode(string, string) (models.UrlDb, error)
	GetUrlFromLongUrl(string, string) (models.UrlDb, error)
	DeleteExpiredURLs(string) error // Add this new function to the interface
}

func ConnectDb() {

	databaseConfig := config.GetConfig()
	host := databaseConfig.Database.Host

	client, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("mongodb://%s", host)))
	if err != nil {
		logger.Log.Error("Error while making the new mongoDB client", err)
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		logger.Log.Error("Not able to connect to database", err)
		panic(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}
	Mgr = &repository.Test{Connection: client, Ctx: ctx, Cancel: cancel}
}

func DeleteExpiredURLs() {
	err := Mgr.DeleteExpiredURLs(constant.UrlCollection)
	if err != nil {
		logger.Log.Error("Error deleting expired URLs: ", err)
	}
}
