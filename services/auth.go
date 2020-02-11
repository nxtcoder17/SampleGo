package services

import (
	"SampleCrud/db"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

const collectionName string = "users"

func getCollection() *mongo.Collection {
	return db.GetDB().Collection(collectionName)
}

func Register(user db.User) *mongo.InsertOneResult {
	inserted, err := getCollection().InsertOne(db.Context, user)
	if err != nil {
		log.Fatal(err)
	}
	return inserted
}

