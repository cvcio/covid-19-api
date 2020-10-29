package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB is a collection of support for different DB technologies. Currently
// only MongoDB has been implemented. We want to be able to access the raw
// database support for the given DB so an interface does not work. Each
// database is too different.
type DB struct {
	// MongoDB Support.
	database *mongo.Database
}

// New return a new mongo database
func New(uri, database string, timeout time.Duration) (*DB, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("Couldn't connect to mongo: %v", err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		return nil, fmt.Errorf("Mongo client couldn't connect with background context: %v", err)
	}

	return &DB{
		database: client.Database(database),
	}, nil
}

// Close closes a DB value being used with MongoDB.
func (db *DB) Close() error {
	return db.database.Client().Disconnect(context.Background())
}

// Copy returns a new DB value for use with MongoDB based on master session.
func (db *DB) Copy() *mongo.Database {
	return db.database
}
