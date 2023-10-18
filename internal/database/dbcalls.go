package database

import (
	"context"
	"urlshortner/internal/constant"
	"urlshortner/internal/types"

	"go.mongodb.org/mongo-driver/bson"
)

func (mgr *manager) Insert(data interface{}, collectionName string) (interface{}, error) {
	inst := mgr.connection.Database(constant.Database).Collection(collectionName)
	result, err := inst.InsertOne(context.TODO(), data)
	return result.InsertedID, err
}

func (mgr *manager) GetUrlFromCode(code string, collectionName string) (resp types.UrlDb, err error) {
	inst := mgr.connection.Database(constant.Database).Collection(collectionName)
	err = inst.FindOne(context.TODO(), bson.M{"url_code": code}).Decode(&resp)
	return resp, err
}

func (mgr *manager) GetUrlFromLongUrl(longUrl string, collectionName string) (resp types.UrlDb, err error) {
	inst := mgr.connection.Database(constant.Database).Collection(collectionName)
	err = inst.FindOne(context.TODO(), bson.M{"long_url": longUrl}).Decode(&resp)
	return resp, err
}

func (mgr *manager) DeleteUrlByCode(code string, collectionName string) error {
	inst := mgr.connection.Database(constant.Database).Collection(collectionName)
	_, err := inst.DeleteOne(context.TODO(), bson.M{"url_code": code})
	return err
}
