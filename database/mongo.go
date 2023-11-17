package database

import (
	"context"
	"convy/conf"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	Client *mongo.Client
	logger *logrus.Entry
}

func getUri(cfg *conf.AppConfig) string {
	var connectionUri string

	connectionUri = fmt.Sprintf("mongodb://%s:%s@%s:%s",
		cfg.Database.Mongo.Username,
		cfg.Database.Mongo.Password,
		cfg.Database.Mongo.Host,
		cfg.Database.Mongo.Port,
	)

	if cfg.Database.Mongo.ConnectionOptions.ConnectionTimeout == 0 {
	}

	return connectionUri
}

func ConnectToMongoDb(cfg *conf.AppConfig) (*Mongo, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(getUri(cfg)).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return nil, err
	}

	return &Mongo{Client: client}, nil
}

func (m Mongo) Disconnect(ctx context.Context) error {
	if err := m.Client.Disconnect(ctx); err != nil {
		m.logger.Error(err)
	}

	return nil
}

func (m Mongo) Ping(ctx context.Context) bool {
	if err := m.Client.Ping(ctx, nil); err != nil {
		m.logger.Error(err)

		return false
	}

	return true
}
