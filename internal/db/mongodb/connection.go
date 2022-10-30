package mongodb

import (
	"context"
	"fmt"

	"github.com/Kolyan4ik99/log-audit/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewClient(ctx context.Context, c *config.MongoConfig) (*mongo.Client, error) {
	opts := options.Client().ApplyURI(c.Uri)

	opts.SetAuth(options.Credential{
		Username: c.User,
		Password: c.Pass,
	})

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}
	fmt.Println("OK")

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	return client, nil
}
