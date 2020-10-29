package outdoors

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"sync"
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
		log.Println(" Error building query: ", err)
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
	// fmt.Println(result)
	return result, nil
}

func buildQuery(q *Query, types string) (url.Values, error) {
	values := make(url.Values)
	values.Set("location", fmt.Sprintf("%g, %g", q.Lat, q.Lng))
	values.Set("radius", fmt.Sprintf("%d", q.Radius))
	values.Set("type", types)
	values.Set("key", APIKey)
	if len(q.PriceRange) > 0 {
		r, err := ParsePriceRange(q.PriceRange)
		if err != nil {
			log.Println("Error parsing price range")
			return nil, err
		}
		values.Set("minprice", fmt.Sprintf("%d", int(r.From)-1))
		values.Set("maxprice", fmt.Sprintf("%d", int(r.To)-1))
	}
	return values, nil
}

func makeRequest(uri *url.URL) (*http.Response, error) {
	// TODO: implement retries
	// set a timeout and initialize http client
	// timeout := time.Duration(10 * time.Second)
	client := http.Client{
		// Timeout: timeout,
	}
	request, err := http.NewRequest("GET", uri.String(), nil)
	if err != nil {
		fmt.Println("Failed to create request object:")
		return nil, err
	}
	request.Header.Set("Content-type", "application/json")
	return client.Do(request)
}

// Run makes concurrent queries to Google's servers
func (q *Query) Run() []interface{} {
	var w sync.WaitGroup
	places := make([]interface{}, len(q.Destinations))
	rand.Seed(time.Now().UnixNano())
	var l sync.Mutex
	// placechan := make(chan []interface{})
	for i, dst := range q.Destinations {
		w.Add(1)
		// go searchPlace(q, dst, i, &w, placechan)
		go func(types string, i int) {
			defer w.Done()
			// get the result from google's api endpoint query
			res, err := q.find(types)
			if err != nil {
				log.Println("Failed to find places:", err)
				return
			}
			// check for found places
			if len(res.Results) == 0 {
				log.Println("No result found for ", types)
				return
			}
			// load the photos of found places
			for _, result := range res.Results {
				for _, photo := range result.Photos {
					photo.URL = photosURL + "maxwidth=1000&photoreference=" + photo.PhotoRef + "&key=" + APIKey
				}
			}
			// create a random integer to randomize picks
			randI := rand.Intn(len(res.Results))
			l.Lock()
			// log.Println(res.Results)
			fmt.Println(randI)
			// for randI = range res.Results {
			// 	places = append(places, res.Results[randI])
			// }
			places[i] = res.Results[randI]
			l.Unlock()
		}(dst, i)
	}
	w.Wait()
	fmt.Println("run")
	return places
}

// func searchPlace(q *Query, types string, i int, w *sync.WaitGroup, c chan []interface{}) {
// 	// create a mutual exclusion lock to allow the goroutines access the shared varible (the map) concurrently
// 	places := make([]interface{}, len(q.Destinations))
// 	rand.Seed(time.Now().UnixNano())
// 	var l sync.Mutex
// 	defer w.Done()
// 	// get the result from google's api endpoint query
// 	res, err := q.find(types)
// 	if err != nil {
// 		log.Println("Failed to find places:", err)
// 		return
// 	}
// 	// check for found places
// 	if len(res.Results) == 0 {
// 		log.Println("No result found for ", types)
// 		return
// 	}
// 	// load the photos of found places
// 	for _, result := range res.Results {
// 		for _, photo := range result.Photos {
// 			photo.URL = photosURL + "maxwidth=1000&photoreference=" + photo.PhotoRef + "&key=" + APIKey
// 		}
// 	}
// 	// create a random integer to randomize picks
// 	randI := rand.Intn(len(res.Results))
// 	l.Lock()
// 	// log.Println(res.Results)
// 	places[i] = res.Results[randI]
// 	c <- places
// 	close(c)
// 	l.Unlock()
// }
