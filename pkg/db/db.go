package db

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ErrInvalidDBProvided = errors.New("invalid DB provided")

// DB is a collection of support for different DB technologies. Currently
// only MongoDB has been implemented. We want to be able to access the raw
// database support for the given DB so an interface does not work. Each
// database is too different.
type DB struct {
	// MongoDB Support.
	Database *mongo.Database
	Context  *context.Context
}

// New return a new mongo database
func New(uri, database string, timeout time.Duration) (*DB, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri), options.Client().SetSocketTimeout(60*time.Second))
	if err != nil {
		return nil, fmt.Errorf("Couldn't connect to mongo: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, fmt.Errorf("Mongo client couldn't connect with background context: %v", err)
	}

	return &DB{
		Database: client.Database(database),
		Context:  &ctx,
	}, nil
}

// Close closes a DB value being used with MongoDB.
func (db *DB) Close() error {
	return db.Database.Client().Disconnect(*db.Context)
}

// Copy returns a new DB value for use with MongoDB based on master session.
func (db *DB) Copy() *mongo.Database {
	return db.Database
}

// Execute is used to execute MongoDB commands.
func (db *DB) Execute(collName string, f func(*mongo.Collection) error) error {
	if db == nil { //|| db.session == nil {
		return errors.Wrap(ErrInvalidDBProvided, "db == nil")
	}

	return f(db.Database.Collection(collName))
}
