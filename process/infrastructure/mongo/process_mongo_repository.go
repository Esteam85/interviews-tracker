package mongo

import (
	"context"
	"errors"

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

func (p *ProcessMongoRepository) Save(ctx context.Context, process *domain.Process) (err error) {
	_, err = p.collection.InsertOne(ctx, fromPrimitives(process.ToPrimitives()))
	if err != nil {
		var mongoErr mongo.WriteException
		if errors.As(err, &mongoErr) {
			log.Errorf("error trying to insert a process with id %s, %s", process.ProcessID(), mongoErr.Error())
		}
		if mongo.IsDuplicateKeyError(err) {
			return domain.ErrProcessAlreadyExist
		}
	}
	return err
}
