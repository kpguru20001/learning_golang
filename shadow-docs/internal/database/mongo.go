package database

import (
	"context"
	"fmt"
	"log"
	"shadow-docs/configs"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var Database *mongo.Database

func StandardContext() context.Context {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return ctx
}

func Connect() error {
	ctx := StandardContext()

	uri := fmt.Sprintf("mongodb://%s:%s/%s?authSource=%s&ssl=%s",
		configs.Configuration.Database.Host,
		configs.Configuration.Database.Port,
		configs.Configuration.Database.DBName,
		configs.Configuration.Database.AuthDB,
		configs.Configuration.Database.SSLMode,
	)

	opts := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, opts)

	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
		return err
	}

	Client = client
	Database = client.Database(configs.Configuration.Database.DBName)

	return nil
}

func Disconnect() error {
	ctx := StandardContext()

	if err := Client.Disconnect(ctx); err != nil {
		log.Fatalf("Error disconnecting from MongoDB: %v", err)
		return err
	}

	return nil
}
