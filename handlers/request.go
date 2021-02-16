package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/olawolu/outdoors"
	"github.com/olawolu/outdoors/data"
	"github.com/olawolu/outdoors/service"
)

func (sv *Server) getTrips(w http.ResponseWriter, r *http.Request) {
	respond(w, r, data.Recommendations)
}

func (sv *Server) getDestinations(w http.ResponseWriter, r *http.Request) {
	q := &service.Query{
		Destinations: strings.Split(r.URL.Query().Get("journey"), "|"),
	}
	var err error

	q.Lat, err = strconv.ParseFloat(r.URL.Query().Get("lat"), 64)
	if err != nil {
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
	respond(w, r, places)
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
