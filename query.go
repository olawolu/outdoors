package outdoors

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

// Query payload sent to google's servers
type Query struct {
	Lat          float64
	Lng          float64
	Destinations []string
	Radius       int
	PriceRange   string
}

func (q *Query) find(types string) (*response, error) {
	var result = &response{}

	values, err := buildQuery(q, types)
	if err != nil {
		log.Println("Error Parsing Price Range: ", err)
		return nil, err
	}

	uri, err := url.Parse(nearbySearchURL + values.Encode())
	if err != nil {
		fmt.Println("url parsre fail:")
		return nil, err
	}

	res, err := makeRequest(uri)
	if err != nil {
		fmt.Println("Error making http request:")
		return nil, err
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	for {
		if err := decoder.Decode(&result); err != nil {
			log.Println(err)
		}
		break
	}
	return result, nil
}

func buildQuery(q *Query, types string) (url.Values, error) {
	values := make(url.Values)
	values.Set("location", fmt.Sprintf("%g, %g", q.Lat, q.Lng))
	values.Set("radius", fmt.Sprintf("%d", q.Radius))
	values.Set("type", types)
	values.Set("key", APIKey)
	if len(q.PriceRange) > 0 {
		// TODO: parse the price range
		r, err := ParsePriceRange(q.PriceRange)
		if err != nil {
			return nil, err
		}
		values.Set("minprice", fmt.Sprintf("%d", int(r.From)-1))
		values.Set("maxprice", fmt.Sprintf("%d", int(r.To)-1))
	}
	return values, nil
}

func makeRequest(uri *url.URL) (*http.Response, error) {
	// TODO: implement retries
	timeout := time.Duration(10 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	request, err := http.NewRequest("GET", uri.String(), nil)
	if err != nil {
		fmt.Println("Failed to create request object:")
		return nil, err
	}
	request.Header.Set("Content-type", "application/json")
	return client.Do(request)
}
