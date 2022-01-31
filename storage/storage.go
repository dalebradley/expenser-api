package storage

import (
	"context"
	"time"

	"github.com/dalebradley/expenser-api/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//Repository
type Mongo struct{}

// GetAlbums returns the list of Albums
func CreateExpenseResource(expenseResource models.ExpenseResourceDB) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://host.docker.internal:27017"))

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())
	collection := client.Database("expenser").Collection("expenses")
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	_, err = collection.InsertOne(ctx, expenseResource)

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	return err
}
