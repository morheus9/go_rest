package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/mongo"

	"github.com/morheus9/go_rest/internal/user"
	"github.com/morheus9/go_rest/pkg/logging"
)

type db struct {
	collection *mongo.Collection
	logger     *logging.Logger
}

func (d *db) Create(ctx context.Context, user user.User) (string, error) {
	d.logger.Debug("create user")
	result, err := d.collection.InsertOne(ctx, user)
	if err != nil {
		return "", fmt.Errorf("failed to create user %v", err)
	}
	d.logger.Debug("convert insertedID to ObjectID")
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		oid.Hex()
		return oid.Hex(), nil
	}
	d.logger.Trace(user)
	return "", fmt.Errorf("failed to convert objectID to hex")
}

func (d *db) FindOne(context.Context, string) (user.User, error) {
	panic("implement me")
}

func (d *db) Update(context.Context, user.User) error {
	panic("implement me")
}

func (d *db) Delete(context.Context, string) error {
	panic("implement me")
}
func NewStorage(database *mongo.Database, collection string, logger *logging.Logger) user.Storage {

	return &db{
		collection: database.Collection(collection),
		logger:     logger,
	}
}
