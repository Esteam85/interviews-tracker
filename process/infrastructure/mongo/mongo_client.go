package mongo

import (
	"context"
	"os"
	"time"

	"github.com/esteam85/interviews-tracker/process/infrastructure/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client struct {
	client *mongo.Client
	ctx    context.Context
}

func NewClient() (*Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		os.Getenv("MONGO_URI"),
	))
	if err != nil {
		log.Error("error trying to connect to mongo,", err.Error())
		return nil, err
	}

	return &Client{
		client: client,
		ctx:    ctx,
	}, nil
}

func (m *Client) Disconnect() {
	if err := m.client.Disconnect(m.ctx); err != nil {
		log.Error("error trying to disconnect mongo client,", err.Error())
		panic(err)
	}
}

func (m *Client) Client() *mongo.Client {
	return m.client
}
