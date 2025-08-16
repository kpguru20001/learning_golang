package database

import (
	"fmt"
	"log"
	"shadow-docs/configs"
	"shadow-docs/pkg/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var Database *mongo.Database

func Connect() error {
	ctx := utils.StandardContextTimeout()

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
	ctx := utils.StandardContextTimeout()

	if err := Client.Disconnect(ctx); err != nil {
		log.Fatalf("Error disconnecting from MongoDB: %v", err)
		return err
	}

	return nil
}
