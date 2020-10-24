package outdoors

import (
	"fmt"
	"net/http"
	"net/url"
)

// Query payload sent to google's servers
type Query struct {
	Lat          float64
	Lng          float64
	Destinations []string
	Radius       int
	CostRange    string
}

func (q *Query) construct(types string) (*http.Request, error) {
	values := make(url.Values)
	values.Set("location", fmt.Sprintf("%g, %g", q.Lat, q.Lng))
	values.Set("radius", fmt.Sprintf("%d", q.Radius))
	values.Set("type", types)
	values.Set("key", APIKey)
	if len(q.CostRange) > 0 {
		// TODO: parse the cost range
	}
}
