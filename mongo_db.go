package main

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongoDB() *mongo.Database {
	clientOptions := options.Client().ApplyURI(os.Getenv("DB_URI")).SetAuth(options.Credential{
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
	})

	client, err := mongo.Connect(context.TODO(), clientOptions)
	LogIfError(err)

	err = client.Ping(context.TODO(), nil)
	LogIfError(err)

	return client.Database(os.Getenv("DB_NAME"))
}
