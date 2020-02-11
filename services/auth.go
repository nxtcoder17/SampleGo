package services

import (
	"SampleCrud/db"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

const collectionName string = "users"

func getCollection() *mongo.Collection {
	collection := db.GetDB().Collection(collectionName)
	return collection
}

func Register(user db.User) *mongo.InsertOneResult {
	inserted, err := getCollection().InsertOne(db.Context, user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Item has been inserted successfully")
	return inserted
}

