package repository

import (
	"context"
	"time"
	"urlshortner/internal/constant"
	"urlshortner/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Test struct {
	Connection *mongo.Client
	Ctx        context.Context
	Cancel     context.CancelFunc
}

func (mgr *Test) Insert(data interface{}, collectionName string) (interface{}, error) {
	inst := mgr.Connection.Database(constant.Database).Collection(collectionName)
	result, err := inst.InsertOne(context.TODO(), data)
	return result.InsertedID, err
}

func (mgr *Test) GetUrlFromCode(code string, collectionName string) (resp models.UrlDb, err error) {
	inst := mgr.Connection.Database(constant.Database).Collection(collectionName)
	err = inst.FindOne(context.TODO(), bson.M{"url_code": code}).Decode(&resp)
	return resp, err
}

func (mgr *Test) GetUrlFromLongUrl(longUrl string, collectionName string) (resp models.UrlDb, err error) {
	inst := mgr.Connection.Database(constant.Database).Collection(collectionName)
	err = inst.FindOne(context.TODO(), bson.M{"long_url": longUrl}).Decode(&resp)
	return resp, err
}

func (mgr *Test) DeleteExpiredURLs(collectionName string) error {
	inst := mgr.Connection.Database(constant.Database).Collection(collectionName)

	currentTime := time.Now().Unix()

	filter := bson.M{"expired_at": bson.M{"$lte": currentTime}}

	_, err := inst.DeleteMany(context.TODO(), filter)
	return err
}
