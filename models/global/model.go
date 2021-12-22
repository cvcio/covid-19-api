package global

import (
	"time"
)

var (
	validKeys = []string{
		"date", "uid", "country", "iso2", "iso3", "loc",
		"population", "cases", "deaths", "recovered",
		"active", "critical", "tests", "new_cases",
		"new_deaths", "new_recovered", "case_fatality_ratio",
		"incidence_rate", "source",
		"tests", "new_tests",
		"tests_rtpcr", "new_tests_rtpcr",
		"tests_rapid", "new_tests_rapid",

		"icu_discharges",
		"hospital_admissions",
		"hospital_discharges",
		"new_hospital_admissions",
		"new_hospital_discharges",
		"intubated_unvac",
		"intubated_vac",

		"icu_occupancy",
		"beds_occupancy",
		"icu_availability",
	}
)

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
