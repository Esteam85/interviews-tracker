package mongo

import (
	"context"
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
	var mongoUri = "<Your Atlas Connection String>"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		mongoUri,
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
