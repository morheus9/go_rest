package db

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
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
	return "", fmt.Errorf("failed to convert objectID to hex, probably objectID: %s", oid)
}

func (d *db) FindOne(ctx context.Context, id string) (u user.User, err error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return u, fmt.Errorf("failed to convert hex to objectID, hex: %s", id)
	}

	filter := bson.M{"id": oid}

	// options.FindOneOptions{}

	result := d.collection.FindOne(ctx, filter)
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return u, fmt.Errorf("ErrEntityNotFound. User not found by ID: %s", id)
		}
		// TODO: ErrEntityNotFound
		return u, fmt.Errorf("failed to find user by ID: %s due to error: %v", id, err)
	}
	if err = result.Decode(&u); err != nil {
		return u, fmt.Errorf("failed to decode user (id:%s) from DB due to error: %v", id, err)
	}
	return u, nil
}

func (d *db) Update(ctx context.Context, user user.User) error {
	objectID, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return fmt.Errorf("failed to convert hex (userID) to objectID, hex %s", user.ID)
	}
	filter := bson.M{"_id": objectID}

	userBytes, err := bson.Marshal(user)
	if err != nil {
		return fmt.Errorf("failed to marshal user (userID) to BSON, userID %v", err)
	}
	var updateUserObj bson.M
	err = bson.Unmarshal(userBytes, &updateUserObj)
	if err != nil {
		return fmt.Errorf("failed to unmarshal user bytes to BSON, error %v", err)
	}
	delete(updateUserObj, "_id")

	updade := bson.M{
		"$set": updateUserObj,
	}
	result, err := d.collection.UpdateOne(ctx, filter, updade)
	if err != nil {
		return fmt.Errorf("failed to execute update user query, error: %v", err)
	}
	if result.MatchedCount == 0 {
		// TODO: ErrEntityNotFound
		return fmt.Errorf("not found")
	}
	d.logger.Tracef("Matched %d documents amd modified %d", result.MatchedCount, result.ModifiedCount)
	return nil
}

func (d *db) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("failed to convert hex (userID) to objectID, (ID:%s)", id)
	}
	filter := bson.M{"_id": objectID}

	result, err := d.collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete ID %s: %v", filter, err)
	}
	if result.DeletedCount == 0 {
		return fmt.Errorf("not found")
	}
	d.logger.Tracef("Deleted %d documents", result.DeletedCount)
	return nil
}
func NewStorage(database *mongo.Database, collection string, logger *logging.Logger) user.Storage {

	return &db{
		collection: database.Collection(collection),
		logger:     logger,
	}
}
