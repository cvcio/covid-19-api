package global

import (
	"context"
	"strings"
	"time"

	"github.com/cvcio/covid-19-api/pkg/db"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// List Endpoint
func List(dbConn *db.DB, optionsList ...func(*ListOptions)) ([]*Global, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// get current date
	year, month, day := time.Now().Date()
	date := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)

	// parse list options to build mongo query
	opts := DefaultOpts()
	for _, o := range optionsList {
		o(&opts)
	}

	filter := bson.M{}
	// set default date limit (today)
	filter["date"] = bson.M{"$gte": date}

	// set country filter if exists in query param
	if opts.ISO3 != "" {
		filter["iso3"] = strings.ToUpper(opts.ISO3)
	}

	// build date limit query
	var dateQuery bson.D
	// from param exists
	if !opts.From.IsZero() {
		dateQuery = append(dateQuery, bson.E{"$gte", opts.From})
	}
	// to param exists
	if !opts.To.IsZero() {
		dateQuery = append(dateQuery, bson.E{"$lte", opts.To})
	}
	// override default date query
	if len(dateQuery) > 0 {
		filter["date"] = dateQuery
	}
	// set projection fields
	projection := bson.D{{"_id", 0}}
	if !strings.Contains(opts.Keys, "all") {
		// validate fileds
		keys := strings.Split(opts.Keys, ",")
		for _, key := range keys {
			if IsValidKey(strings.TrimSpace(key), validKeys) {
				projection = append(projection, bson.E{key, 1})
			}
		}
	}

	// set find options
	findOptions := options.Find().SetSort(bson.D{{"date", 1}, {"iso3", 1}})

	// decode to list
	var list []*Global
	f := func(collection *mongo.Collection) error {
		c, err := collection.Find(ctx, filter, findOptions.SetProjection(projection))
		if err != nil {
			return err
		}

		defer c.Close(ctx)
		for c.Next(ctx) {
			var entry *Global
			err := c.Decode(&entry)
			if err != nil {
				return err
			}
			list = append(list, entry)
		}
		return nil
	}

	if err := dbConn.Execute("global", f); err != nil {
		return nil, errors.Wrap(err, "db.global.find()")
	}

	return list, nil
}
