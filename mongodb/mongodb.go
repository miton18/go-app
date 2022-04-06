package mongodb

import (
	"context"
	"os"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetHost() string {
	return os.Getenv(MONGODB_ADDON_HOST)
}

func GetURI() string {
	return os.Getenv(MONGODB_ADDON_URI)
}

func New() (*mongo.Client, error) {

	client, err := mongo.NewClient(
		options.Client().ApplyURI(GetURI()),
	)
	if err != nil {
		return nil, errors.Wrap(err, "cannot create MongoDB client")
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "cannot connect to MongoDB instance")
	}

	return client, nil
}
