package mongo

import (
	"context"

	"github.com/esteam85/interviews-tracker/process/domain"
	"github.com/esteam85/interviews-tracker/process/infrastructure/log"

	"go.mongodb.org/mongo-driver/mongo"
)

type ProcessMongoRepository struct {
	collection *mongo.Collection
}

func NewProcessMongoRepository(client *mongo.Client) *ProcessMongoRepository {
	var dbName = "Cluster0"
	var collectionName = "processes"
	collection := client.Database(dbName).Collection(collectionName)
	return &ProcessMongoRepository{
		collection: collection,
	}
}

func (p *ProcessMongoRepository) Save(ctx context.Context, process *domain.Process) error {
	_, err := p.collection.InsertOne(ctx, process.ToPrimitives())
	if err != nil {
		log.Errorf("error trying to insert a process with id %s, %s", process.ProcessID(), err.Error())
		return err
	}
	return nil
}
