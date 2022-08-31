package storage

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	PersonCollection = "person"
)

type StorageMongoDB struct {
	client *mongo.Client
	db     *mongo.Database
}

func MongoConnect() StorageMongoDB {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	uri := os.Getenv("MONGODB_URI")

	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	db := client.Database("csv")

	return StorageMongoDB{client, db}
}

func (storageMongo StorageMongoDB) GetCollection(collectionName string) *mongo.Collection {
	return storageMongo.db.Collection(collectionName)
}

func (storageMongo StorageMongoDB) Disconnect() {
	err := storageMongo.client.Disconnect(context.TODO())

	if err != nil {
		panic(err)
	}
}
