package gr_vaccines

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	validKeys = []string{
		"date", "uid", "geo_unit", "state", "region", "loc",
		"population", "source",
		"total_distinct_persons", "total_vaccinations", "day_total", "day_diff",
	}
)

// Loc Location struct
type Loc struct {
	Type        string    `bson:"type" json:"type,omitempty"`
	Coordinates []float64 `bson:"coordinates" json:"coordinates,omitempty"`
}

// GRVaccine represents the document structure of documents in the
// 'GRVaccine' collection parsed from sources JHU and WOM
type GRVaccine struct {
	ID primitive.ObjectID `bson:"_id" json:"-"`
	// covid related date
	Date          time.Time `bson:"date" json:"date,omitempty"`
	LastUpdatedAt time.Time `bson:"last_updated_at" json:"last_updated_at,omitempty"`
	// location data
	UID        string `bson:"uid" json:"uid,omitempty"`
	AreaID     int    `bson:"areaid" json:"areaid,omitempty"`
	GeoUnit    string `bson:"geo_unit" json:"geo_unit,omitempty"`
	State      string `bson:"state" json:"state,omitempty"`
	Region     string `bson:"region" json:"region,omitempty"`
	Area       string `bson:"area" json:"area,omitempty"`
	Loc        *Loc   `bson:"loc" json:"loc,omitempty"`
	Population int64  `bson:"population" json:"population"`
	// vaccines related metrics
	DayDiff              int32 `bson:"daydiff" json:"daydiff,omitempty"`
	DayTotal             int32 `bson:"daytotal" json:"daytotal,omitempty"`
	TotalDistinctPersons int32 `bson:"totaldistinctpersons" json:"totaldistinctpersons,omitempty"`
	TotalVaccination     int32 `bson:"totalvaccinations" json:"totalvaccinations,omitempty"`
	// source
	Source string `bson:"source" json:"source,omitempty"`
}

// ListOptions represents the filter structure to query
// the database
type ListOptions struct {
	Limit int
	UID   string
	Keys  string
	From  time.Time
	To    time.Time
	Key   string
}

// NewListOpts create a new ListOptions struct
func NewListOpts() []func(*ListOptions) {
	return make([]func(*ListOptions), 0)
}

// Limit sets the limit
func Limit(i int) func(*ListOptions) {
	return func(l *ListOptions) {
		l.Limit = i
	}
}

// UID sets the uid country code
func UID(i string) func(*ListOptions) {
	return func(l *ListOptions) {
		l.UID = i
	}
}

// Keys sets the keys to return
func Keys(i string) func(*ListOptions) {
	return func(l *ListOptions) {
		l.Keys = i
	}
}

// From sets the start date to retrieve data from
func From(i time.Time) func(*ListOptions) {
	return func(l *ListOptions) {
		l.From = i
	}
}

// To sets the end date to retrieve data from
func To(i time.Time) func(*ListOptions) {
	return func(l *ListOptions) {
		l.To = i
	}
}

// DefaultOpts sets the defaults
func DefaultOpts() ListOptions {
	l := ListOptions{}
	l.Limit = -1
	return l
}

// IsValidKey checks if a string is in an array
func IsValidKey(str string, list []string) bool {
	for _, s := range list {
		if s == str {
			return true
		}
	}
	return false
}
