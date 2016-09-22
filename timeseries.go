package mrfusion

import "golang.org/x/net/context"

// Query retrieves a Response from a TimeSeries.
type Query struct {
	Command string // Command is the query itself
	DB      string // DB is optional and if empty will not be used.
	RP      string // RP is a retention policy and optional; if empty will not be used.
}

// Response is the result of a query against a TimeSeries
type Response interface {
	MarshalJSON() ([]byte, error)
}

// MonitoredService is a service sending monitoring data to a `TimeSeries`
type MonitoredService struct {
	TagKey   string `json:"tagKey"`
	TagValue string `json:"tagValue"`
	Type     string `json:"type"`
}

// Represents a queryable time series database.
type TimeSeries interface {
	// Query retrieves time series data from the database.
	Query(context.Context, Query) (Response, error)
	// MonitoredServices retrieves all services sending monitoring data to this `TimeSeries`
	MonitoredServices(context.Context) ([]MonitoredService, error)
}