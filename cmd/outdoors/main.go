package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"

	"github.com/olawolu/outdoors"
)

func init() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatalf("Failed to load environment variables: %v", err)
	}
}

func main() {
	outdoors.APIKey = os.Getenv("PLACES_KEY")
	http.HandleFunc("/journeys", cors(func(w http.ResponseWriter, r *http.Request) {
		respond(w, r, outdoors.Recommendations)
	}))
	http.HandleFunc("/recommendations", cors(func(w http.ResponseWriter, r *http.Request) {
		q := &outdoors.Query{
			Destinations: strings.Split(r.URL.Query().Get("journey"), "|"),
		}
		// fmt.Println(r.URL.Query())
		// fmt.Println(r.URL.Query().Get("lat"))

		// fmt.Println(strings.Split(r.URL.Query().Get("journey"), "|"))

		var err error
		q.Lat, err = strconv.ParseFloat(r.URL.Query().Get("lat"), 64)
		if err != nil {
			// fmt.Println(q)
			handleError(err)
		}
		q.Lng, err = strconv.ParseFloat(r.URL.Query().Get("lng"), 64)
		if err != nil {
			handleError(err)
		}
		q.Radius, err = strconv.Atoi(r.URL.Query().Get("radius"))
		if err != nil {
			handleError(err)
		}
		q.PriceRange = r.URL.Query().Get("cost")
		places := q.Run()
		// fmt.Println(places)
		// places, ok := <-placechan
		// if ok {
		// 	fmt.Println("Channel is open!")
		// 	respond(w, r, places)
		// } else {
		// 	fmt.Println("Channel is closed!")
		// }
		respond(w, r, places)

	}))
	log.Println("listening")
	err := http.ListenAndServe(":8080", http.DefaultServeMux)
	if err != nil {
		log.Fatal("An error occured")
	}
}

func respond(w http.ResponseWriter, r *http.Request, data []interface{}) error {
	log.Println("respond")
	publicData := make([]interface{}, len(data))
	fmt.Println(len(data))
	for i, d := range data {
		publicData[i] = outdoors.Display(d)
	}
	// log.Println(publicData)

	return json.NewEncoder(w).Encode(publicData)
}

func cors(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		f(w, r)
	}
}

func handleError(err error) {
	var w http.ResponseWriter
	http.Error(w, err.Error(), http.StatusBadRequest)
	log.Println(err)
	return
}
