package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/olawolu/outdoors/handlers"
	"github.com/olawolu/outdoors/service"
)

var addr string

var server = handlers.Server{}

func init() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Printf("Failed to load .env file: %v", err)
	}
	service.APIKey = os.Getenv("PLACES_KEY")
	// addr = ":" + os.Getenv("PORT")
}

func main() {
	server.Init()

	err := server.Run()
	if err != nil {
		log.Fatal(err)
	}
}
