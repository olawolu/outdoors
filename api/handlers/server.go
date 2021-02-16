package handlers

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Server defines the rest API server
type Server struct {
	Host   string
	Port   string
	Router *mux.Router
}

// Init creates the server object
func (sv *Server) Init() {
	sv.Host = "0.0.0.0"
	sv.Port = os.Getenv("PORT")
	sv.Router = mux.NewRouter()
}

// Run the server and serve requests
func (sv *Server) Run() error {
	sv.Router.Methods("GET").Path("/journeys").HandlerFunc(sv.getTrips)
	sv.Router.Methods("GET").Path("/recommendations").HandlerFunc(sv.getDestinations)

	headers := handlers.AllowedHeaders([]string{"Accept", "Accept-Language", "Content-Type", "Content-Language", "Origin"})
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "OPTIONS"})

	log.Printf("listening on: %s:%s", sv.Host, sv.Port)
	srv := http.Server{
		Addr:         sv.Host + ":" + sv.Port,
		Handler:      handlers.CORS(headers, origins, methods)(sv.Router),
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  5 * time.Second,
	}
	return srv.ListenAndServe()
}
