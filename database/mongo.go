package database

import (
	"context"
	"convy/conf"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getUri(cfg *conf.AppConfig) string {
	var connectionUri string

	connectionUri = fmt.Sprintf("mongodb://%s:%s@%s:%s",
		cfg.Database.Mongo.Username,
		cfg.Database.Mongo.Password,
		cfg.Database.Mongo.Host,
		cfg.Database.Mongo.Port,
	)

	return connectionUri
}

func ConnectToMongoDb(cfg *conf.AppConfig) (*mongo.Client, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(getUri(cfg)).SetServerAPIOptions(serverAPI)
	opts.MaxPoolSize = &cfg.Database.Mongo.ConnectionOptions.MaxPoolSize
	opts.ConnectTimeout = &cfg.Database.Mongo.ConnectionOptions.ConnectionTimeout
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return nil, err
	}

	return client, nil
}
