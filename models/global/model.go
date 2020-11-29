package global

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	validKeys = []string{
		"date", "uid", "country", "iso2", "iso3", "loc",
		"population", "cases", "deaths", "recovered",
		"active", "critical", "tests", "new_cases",
		"new_deaths", "new_recovered", "case_fatality_ratio",
		"incidence_rate", "source", "cumulative_rtpcr_tests_raw",
		"estimated_new_rtpcr_tests", "cumulative_rapid_tests_raw",
		"esitmated_new_rapid_tests", "estimated_new_total_tests",
	}
)

// Loc Location struct
type Loc struct {
	Type        string    `bson:"type" json:"type,omitempty"`
	Coordinates []float64 `bson:"coordinates" json:"coordinates,omitempty"`
}

// Global represents the document structure of documents in the
// 'global' collection parsed from sources JHU and WOM
type Global struct {
	ID primitive.ObjectID `bson:"_id" json:"-"`
	// covid related date
	Date          time.Time `bson:"date" json:"date,omitempty"`
	LastUpdatedAt time.Time `bson:"last_updated_at" json:"last_updated_at,omitempty"`
	// location data
	UID        int32  `bson:"uid" json:"uid,omitempty"`
	Country    string `bson:"country" json:"country,omitempty"`
	ISO2       string `bson:"iso2" json:"iso2,omitempty"`
	ISO3       string `bson:"iso3" json:"iso3,omitempty"`
	Loc        *Loc   `bson:"loc" json:"loc,omitempty"`
	Population int64  `bson:"population" json:"population,omitempty"`
	// covid related metrics
	Cases        int32 `bson:"cases" json:"cases,omitempty"`
	Deaths       int32 `bson:"deaths" json:"deaths,omitempty"`
	Recovered    int32 `bson:"recovered" json:"recovered,omitempty"`
	Active       int32 `bson:"active" json:"active,omitempty"`
	Critical     int32 `bson:"critical" json:"critical,omitempty"`
	Tests        int32 `bson:"tests" json:"tests,omitempty"`
	NewCases     int32 `bson:"new_cases" json:"new_cases,omitempty"`
	NewDeaths    int32 `bson:"new_deaths" json:"new_deaths,omitempty"`
	NewRecovered int32 `bson:"new_recovered" json:"new_recovered,omitempty"`
	// covid related claculated values
	CaseFatalityRatio float64 `bson:"case_fatality_ratio" json:"case_fatality_ratio,omitempty"`
	IncidenceRate     float64 `bson:"incidence_rate" json:"incidence_rate,omitempty"`
	// source
	Source string `bson:"source" json:"source,omitempty"`
}

// ListOptions represents the filter structure to query
// the database
type ListOptions struct {
	Limit int
	ISO3  string
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

// ISO3 sets the iso3 country code
func ISO3(i string) func(*ListOptions) {
	return func(l *ListOptions) {
		l.ISO3 = i
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
	l.ISO3 = ""
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
