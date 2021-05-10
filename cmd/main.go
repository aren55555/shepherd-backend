package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/aren55555/shepherd-backend/api"
	"github.com/aren55555/shepherd-backend/data/mem"
)

var (
	port         = flag.Int("port", 8080, "the port for the HTTP server to listen on")
	fileLocation = flag.String("location", "seed.json", "the location of the seed data JSON file")
)

func main() {
	flag.Parse()

	memDataStore := mem.New()
	if err := memDataStore.Seed(*fileLocation); err != nil {
		panic(err)
	}

	apiHandler := api.New(memDataStore)
	http.Handle("/api/", apiHandler)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
