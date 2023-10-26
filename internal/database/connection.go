package database

import (
	"context"
	"fmt"
	"os"
	"time"
	"urlshortner/internal/constant"
	"urlshortner/internal/logger"
	"urlshortner/internal/models"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type test struct {
	connection *mongo.Client
	ctx        context.Context
	cancel     context.CancelFunc
}

var Mgr Manager

type Manager interface {
	Insert(interface{}, string) (interface{}, error)
	GetUrlFromCode(string, string) (models.UrlDb, error)
	GetUrlFromLongUrl(string, string) (models.UrlDb, error)
	DeleteExpiredURLs(string) error // Add this new function to the interface
}

func ConnectDb() {
	uri := os.Getenv("DB_HOST")
	client, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("%s%s", "mongodb://", uri)))
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
	Mgr = &test{connection: client, ctx: ctx, cancel: cancel}
}

func DeleteExpiredURLs() {
	err := Mgr.DeleteExpiredURLs(constant.UrlCollection)
	if err != nil {
		logger.Log.Error("Error deleting expired URLs: ", err)
	}
}
