package mongo

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/esteam85/interviews-tracker/process/infrastructure/log"
	"go.mongodb.org/mongo-driver/mongo"
)

type Client struct {
	client *mongo.Client
	ctx    context.Context
}

func NewClient() (*Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI(
		os.Getenv("MONGO_URI"),
	).SetMaxPoolSize(1600)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Error("error trying to connect to mongo,", err.Error())
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("success!, now you are Connected to MongoDB Services!")

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
