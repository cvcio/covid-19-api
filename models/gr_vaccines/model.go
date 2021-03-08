package gr_vaccines

import (
	"time"
)

var (
	validKeys = []string{
		"date", "uid", "geo_unit", "state", "region", "loc",
		"population", "source",
		"total_distinct_persons", "total_vaccinations", "day_total", "day_diff",
		"daily_dose_1", "daily_dose_2", "total_dose_1", "total_dose_2",
	}
)

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
