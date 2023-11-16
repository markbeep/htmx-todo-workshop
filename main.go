package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// command line flags
var (
	port = flag.String("port", "3000", "the port to serve the website on")
	host = flag.String("host", "0.0.0.0", "the host to serve the website on")
)

func main() {
	flag.Parse() // parses command line flags defined above

	r := chi.NewRouter()
	r.Use(middleware.Logger) // adds helpful logging

	// starts listening
	fullHost := fmt.Sprintf("%s:%s", *host, *port)
	log.Printf("serving on %s", fullHost)
	if err := http.ListenAndServe(fullHost, r); err != nil {
		log.Fatal(err)
	}
}
