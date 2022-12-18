package storage

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dalebradley/expenser-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//Repository
type Mongo struct{}

// CreateExpenseResource creates an Expense
func CreateExpenseResource(expenseResource models.ExpenseResourceRest) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://host.docker.internal:27017"))

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())

	collection := client.Database("expenser").Collection("expenses")
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = collection.InsertOne(ctx, expenseResource)

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	return err
}

// GetExpenseResource gets an Expense
func GetExpenseResource(id string) (*models.ExpenseResourceDB, error) {
	var resource models.ExpenseResourceDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://host.docker.internal:27017"))

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())

	collection := client.Database("expenser").Collection("expenses")
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	dbResource := collection.FindOne(ctx, bson.M{"id": id})
	err = dbResource.Err()
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			err = fmt.Errorf("no expense resources found for id: %s", id)
			return nil, nil
		}
		return nil, err
	}

	err = dbResource.Decode(&resource)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	return &resource, nil
}
