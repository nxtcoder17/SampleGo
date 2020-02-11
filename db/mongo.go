package db

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var Client *mongo.Client
var Context context.Context

const uri string = "mongodb://localhost:27017"
const DB string = "sampleDb"

func Init() {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	_, err = client.ListDatabaseNames(ctx, gin.H{})
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(databases)

	Context = ctx
	Client = client
}

func GetDB() *mongo.Database {
	return Client.Database(DB)
}

