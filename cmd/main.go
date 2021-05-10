package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/aren55555/shepherd-backend/api"
	"github.com/aren55555/shepherd-backend/auth"
	"github.com/aren55555/shepherd-backend/data/mem"
)

var (
	port         = flag.Int("port", 8080, "the port for the HTTP server to listen on")
	fileLocation = flag.String("location", "seed.json", "the location of the seed data JSON file")
	authEnabled  = flag.Bool("auth", false, "whether the cookie based auth is enabled or not")
)

func main() {
	flag.Parse()

	memDataStore := mem.New()
	if err := memDataStore.Seed(*fileLocation); err != nil {
		panic(err)
	}

	apiHandler := api.New(memDataStore)

	var h http.Handler = apiHandler
	if *authEnabled {
		fmt.Println("AUTH IS ENABLED, ENSURE AN AUTH COOKIE IS SENT!")
		h = auth.Wrap(apiHandler)
	}

	http.Handle("/api/v1/", h)

	fmt.Printf("Server started on %d", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
