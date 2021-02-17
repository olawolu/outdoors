package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	var addr = flag.String("addr", os.Getenv("PORT"), "website address")
	flag.Parse()
	mux := http.NewServeMux()
	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("client"))))
	log.Println("Serving website at:", *addr)
	err := http.ListenAndServe(":"+*addr, mux)
	if err != nil {
		log.Fatal(err)
	}
}
