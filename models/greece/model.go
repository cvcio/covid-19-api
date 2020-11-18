package greece

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Greece represents the document structure of documents in the
// 'greece' collection parsed from sources JHU and WOM
type Greece struct {
	ID primitive.ObjectID `bson:"_id" json:"id"`
	// covid related date
	Date time.Time `bson:"date" json:"date"`
	// location data
	UID     int32  `bson:"uid" json:"uid"`
	Region  string `bson:"region" json:"region"`
	Country string `bson:"country" json:"country"`
	ISO2    string `bson:"iso2" json:"iso2"`
	ISO3    string `bson:"iso3" json:"iso3"`
	Loc     struct {
		Type        string    `bson:"type" json:"type"`
		Coordinates []float64 `bson:"coordinates" json:"coordinates"`
	} `bson:"loc" json:"loc"`
	Population int64 `bson:"population" json:"population"`
	// covid related metrics
	Cases        int32 `bson:"cases" json:"cases"`
	Deaths       int32 `bson:"deaths" json:"deaths"`
	Recovered    int32 `bson:"recovered" json:"recovered"`
	Active       int32 `bson:"active" json:"active"`
	Critical     int32 `bson:"critical" json:"critical"`
	Tests        int32 `bson:"tests" json:"tests"`
	NewCases     int32 `bson:"new_cases" json:"new_cases"`
	NewDeaths    int32 `bson:"new_deaths" json:"new_deaths"`
	NewRecovered int32 `bson:"new_recovered" json:"new_recovered"`
	// covid related claculated values
	CaseFatalityRatio float32 `bson:"case_fatality_ratio" json:"case_fatality_ratio"`
	IncidenceRate     float32 `bson:"incidence_rate" json:"incidence_rate"`
	// source
	Source string `bson:"source" json:"source"`
}
