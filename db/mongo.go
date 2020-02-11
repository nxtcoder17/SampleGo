package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

const url string = "http://localhost:27017"
var client *mongo.Client
var Context context.Context

const DB string = "sampleDb"

func Init() {
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	Context = ctx
}

func GetDB() *mongo.Database {
	return client.Database(DB)
}

