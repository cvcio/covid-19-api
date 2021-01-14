package gr_vaccines

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
func List(dbConn *db.DB, optionsList ...func(*ListOptions)) ([]*map[string]interface{}, error) {
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
	if opts.UID != "" {
		if !strings.Contains(opts.UID, "all") {
			filter["uid"] = strings.ToUpper(opts.UID)
		}
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
	if !strings.Contains(opts.Keys, "all") && opts.Keys != "" {
		// validate fileds
		keys := strings.Split(opts.Keys, ",")
		for _, key := range keys {
			if IsValidKey(strings.TrimSpace(key), validKeys) {
				projection = append(projection, bson.E{strings.TrimSpace(key), 1})
			}
		}
	}

	// set find options
	findOptions := options.Find().SetSort(bson.D{{"date", 1}, {"uid", 1}})

	// decode to list
	var list []*map[string]interface{}
	f := func(collection *mongo.Collection) error {
		c, err := collection.Find(ctx, filter, findOptions.SetProjection(projection))
		if err != nil {
			return err
		}

		defer c.Close(ctx)
		for c.Next(ctx) {
			var entry *map[string]interface{}
			err := c.Decode(&entry)
			if err != nil {
				return err
			}
			list = append(list, entry)
		}
		return nil
	}

	if err := dbConn.Execute("gr_vaccines", f); err != nil {
		return nil, errors.Wrap(err, "db.gr_vaccines.find()")
	}

	return list, nil
}

// Agg Aggregate Data
func Agg(dbConn *db.DB, optionsList ...func(*ListOptions)) ([]*map[string]interface{}, error) {
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
	if opts.UID != "" {
		if !strings.Contains(opts.UID, "all") {
			filter["uid"] = strings.ToUpper(opts.UID)
		}
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
	// set group fields
	group := bson.D{
		{"_id", "$uid"},
		{"uid", bson.D{{"$first", "$uid"}}},
		{"geo_unit", bson.D{{"$first", "$geo_unit"}}},
		{"state", bson.D{{"$first", "$state"}}},
		{"loc", bson.D{{"$first", "$loc"}}},
		{"region", bson.D{{"$first", "$region"}}},
		{"sources", bson.D{{"$addToSet", "$source"}}},
		{"population", bson.D{{"$first", "$population"}}},
		{"from", bson.D{{"$first", "$date"}}},
		{"to", bson.D{{"$last", "$date"}}},
		{"last_updated_at", bson.D{{"$last", "$last_updated_at"}}},
	}
	if !strings.Contains(opts.Keys, "all") && opts.Keys != "" {
		// validate fileds
		keys := strings.Split(opts.Keys, ",")
		for _, key := range keys {
			if IsValidKey(strings.TrimSpace(key), validKeys) {
				group = append(group, bson.E{strings.TrimSpace(key), bson.D{{"$push", "$" + strings.TrimSpace(key)}}})
			}
		}
	} else {
		group = append(group, bson.E{"day_diff", bson.D{{"$push", "$day_diff"}}})
		group = append(group, bson.E{"day_total", bson.D{{"$push", "$day_total"}}})
		group = append(group, bson.E{"total_distinct_persons", bson.D{{"$push", "$total_distinct_persons"}}})
		group = append(group, bson.E{"total_vaccinations", bson.D{{"$push", "$total_vaccinations"}}})
		group = append(group, bson.E{"new_total_distinct_persons", bson.D{{"$push", "$new_total_distinct_persons"}}})
		group = append(group, bson.E{"new_total_vaccinations", bson.D{{"$push", "$new_total_vaccinations"}}})
	}

	// set agg options
	o := options.Aggregate()

	// set aggregation pipeline
	pipeline := mongo.Pipeline{
		{{"$match", filter}},
		{{"$group", group}},
		{{"$sort", bson.D{{"iso3", 1}}}},
		{{"$project", bson.D{{"_id", 0}}}},
	}
	// decode to list
	var list []*map[string]interface{}
	f := func(collection *mongo.Collection) error {
		c, err := collection.Aggregate(ctx, pipeline, o)
		if err != nil {
			return err
		}

		defer c.Close(ctx)
		for c.Next(ctx) {
			var entry *map[string]interface{}
			err := c.Decode(&entry)
			if err != nil {
				return err
			}
			list = append(list, entry)
		}
		return nil
	}

	if err := dbConn.Execute("gr_vaccines", f); err != nil {
		return nil, errors.Wrap(err, "db.gr_vaccines.agg()")
	}

	return list, nil
}

// Sum Data
func Sum(dbConn *db.DB, optionsList ...func(*ListOptions)) ([]*map[string]interface{}, error) {
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
	if opts.UID != "" {
		if !strings.Contains(opts.UID, "all") {
			filter["uid"] = strings.ToUpper(opts.UID)
		}
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
	// set group fields
	group := bson.D{
		{"_id", "$uid"},
		{"uid", bson.D{{"$first", "$uid"}}},
		{"geo_unit", bson.D{{"$first", "$geo_unit"}}},
		{"state", bson.D{{"$first", "$state"}}},
		{"loc", bson.D{{"$first", "$loc"}}},
		{"region", bson.D{{"$first", "$region"}}},
		{"sources", bson.D{{"$addToSet", "$source"}}},
		{"population", bson.D{{"$first", "$population"}}},
		{"last_updated_at", bson.D{{"$last", "$last_updated_at"}}},
	}

	group = append(group, bson.E{"total_distinct_persons", bson.D{{"$last", "$total_distinct_persons"}}})
	group = append(group, bson.E{"total_vaccinations", bson.D{{"$last", "$total_vaccinations"}}})

	group = append(group, bson.E{"day_diff", bson.D{{"$sum", "$day_diff"}}})
	group = append(group, bson.E{"day_total", bson.D{{"$sum", "$day_total"}}})

	group = append(group, bson.E{"new_total_distinct_persons", bson.D{{"$sum", "$new_total_distinct_persons"}}})
	group = append(group, bson.E{"new_total_vaccinations", bson.D{{"$sum", "$new_total_vaccinations"}}})

	// set agg options
	o := options.Aggregate()

	// set aggregation pipeline
	pipeline := mongo.Pipeline{
		{{"$match", filter}},
		{{"$group", group}},
		{{"$sort", bson.D{{"iso3", 1}}}},
		{{"$project", bson.D{{"_id", 0}}}},
	}
	// decode to list
	var list []*map[string]interface{}
	f := func(collection *mongo.Collection) error {
		c, err := collection.Aggregate(ctx, pipeline, o)
		if err != nil {
			return err
		}

		defer c.Close(ctx)
		for c.Next(ctx) {
			var entry *map[string]interface{}
			err := c.Decode(&entry)
			if err != nil {
				return err
			}
			list = append(list, entry)
		}
		return nil
	}

	if err := dbConn.Execute("gr_vaccines", f); err != nil {
		return nil, errors.Wrap(err, "db.gr_vaccines.sum()")
	}

	return list, nil
}
