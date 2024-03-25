package mongo

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"

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

func (p *ProcessMongoRepository) GetAll(ctx context.Context) (domain.Processes, error) {
	cursor, err := p.collection.Find(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	var processes domain.Processes
	for cursor.Next(ctx) {
		var processDTO ProcessDTO
		err = cursor.Decode(&processDTO)
		if err != nil {
			log.Error(err)
			continue
		}
		process, err := processDTO.ToProcess()
		if err != nil {
			log.Error(err)
			continue
		}
		processes = append(processes, *process)
	}

	if err = cursor.Err(); err != nil {
		log.Error(err)
	}

	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err = cursor.Close(ctx)
		if err != nil {
			log.Error(err)
		}
	}(cursor, ctx)

	return processes, nil
}
